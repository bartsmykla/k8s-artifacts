package main

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"sort"
	"strings"

	"google.golang.org/api/iterator"
	"cloud.google.com/go/storage"
)

const (
	StageObject = iota
	ReleaseBinary
	ReleaseExtra
	ReleaseOther
	ReleaseUnknown
)

type Object struct {
	Type int
	Name string
	Version string
	Rest []string
	raw string
}

type Objects struct {
	c chan *storage.ObjectAttrs
}

func main() {
	projectId := "kubernetes-release-learning"
	credentialsPath := "gcp_credentials.json"
	bucketNames := []string{
		"kubernetes-release-learning-gcb",
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialsPath))
	if err != nil {
		log.Fatal(err)
	}

	buckets := client.Buckets(ctx, projectId)

	for {
		_, bucketErr := buckets.Next()
		if bucketErr == iterator.Done {
			break
		}
		if bucketErr != nil {
			log.Fatal(bucketErr)
		}

		//fmt.Println("bucket:", bucketAttrs.Name)
	}
	fmt.Println() // print empty line

	for _, bucketName := range bucketNames {
		handleProcessedObjects(handleBuckets(client, ctx, bucketName))
	}
}

func handleBuckets(
	client *storage.Client,
	ctx context.Context,
	name string,
) []*Object {
	objects := client.Bucket(name).Objects(ctx, nil)

	size := uint64(0)

	result := []*Object{}

	for {
		object, objectErr := objects.Next()
		if objectErr == iterator.Done {
			break
		}
		if objectErr != nil {
			log.Fatal(objectErr)
		}

		result = append(result, processObject(object))

		size += uint64(object.Size)
	}

	//fmt.Println("Size:", humanize.Bytes(size))

	return result
}

func processObject(object *storage.ObjectAttrs) *Object {
	raw := object.Name
	parts := strings.Split(object.Name, "/")

	if len(parts) < 2 {
		log.Fatal("fatal in: ", object.Name)
	}

	var objectName string
	var objectType string

	objectName, parts = parts[len(parts)-1], parts[:len(parts)-1]
	objectType, parts = parts[0], parts[1:]

	finalObject := Object{
		Type: 0,
		Name: objectName,
		Rest: []string{},
		raw: raw,
	}
	switch objectType {
	case "release":
		if len(parts) < 2 {
			finalObject.Rest = parts
			finalObject.Type = ReleaseUnknown
			fmt.Println(raw)
			return &finalObject
		}
		var objectVersion string
		var objectSubtype string
		objectVersion, parts = parts[0], parts[1:]
		objectSubtype, parts = parts[0], parts[1:]

		finalObject.Version = objectVersion
		finalObject.Rest = parts

		switch objectSubtype {
		case "bin":
			finalObject.Type = ReleaseBinary
			break
		case "extra":
			finalObject.Type = ReleaseExtra
			break
		default:
			finalObject.Type = ReleaseOther
		}
	case "stage":
	case "archive":
		break
	default:
		fmt.Println(raw)
	}

	//if finalObject.Type == ReleaseUnknown || finalObject.Type == ReleaseOther {
	//	fmt.Println(finalObject.raw)
	//}

	return &finalObject
}


func parseName(fullName string) (name string, extension string) {
	parts := strings.Split(fullName, ".")
	name, parts = parts[0], parts[1:]

	if len(parts) < 1 {
		return name, ""
	}

	return name, strings.Join(parts, ".")
}

func handleProcessedObjects(objects []*Object) {
	releaseResult := map[string]map[[2]string][]string{}
	for _, object := range objects {
		switch object.Type {
		case ReleaseBinary:
			name, extension := parseName(object.Name)

			if extension == "" {
				extension = "[binary_file]"
			}

			if releaseResult[name] == nil {
				releaseResult[name] = map[[2]string][]string{}
			}

			if len(object.Rest) < 2 {
				log.Fatal("wrong number of strings to generate system and architecture: ", object. Rest)
			}

			arch := [2]string{object.Rest[0], object.Rest[1]}

			releaseResult[name][arch] = append(releaseResult[name][arch], extension)
			break
		}
	}

	printTable(releaseResult)
}

func printTable(result map[string]map[[2]string][]string) {
	artifacts := []string{}
	archs := map[[2]string]struct{}{}

	for artifact, value := range result {
		artifacts = append(artifacts, artifact)
		for arch, _ := range value {
			archs[arch] = struct{}{}
		}
	}

	resultArchs := [][2]string{}
	resultArchsStrings := []string{}
	for arch, _ := range archs {
		resultArchsStrings = append(resultArchsStrings, fmt.Sprintf("%s/%s", arch[0], arch[1]))
		resultArchs = append(resultArchs, arch)
	}

	sort.Slice(resultArchs, func(i, j int) bool {
		a := fmt.Sprintf("%s/%s", resultArchs[i][0], resultArchs[i][1])
		b := fmt.Sprintf("%s/%s", resultArchs[j][0], resultArchs[j][1])

		return a < b
	})
	sort.Strings(resultArchsStrings)
	sort.Strings(artifacts)

	//fmt.Printf("|  | %s |\n", strings.Join(resultArchsStrings, " | "))

	spacer := fmt.Sprintf("|%s|", genMultipliedString("-", getLongestLength(artifacts)))
	for _, arch := range resultArchsStrings {
		spacer += fmt.Sprintf("%s|", genMultipliedString("-", len(arch)))
	}
	//fmt.Println(spacer)

	for _, artifact := range artifacts {
		line := fmt.Sprintf("| %s |", artifact)
		for _, arch := range resultArchs {
			elem := result[artifact][arch]
			sort.Strings(elem)
			if len(elem) < 1 {
				line += "  |"
				continue
			}

			line += fmt.Sprintf(" %s |", strings.Join(elem, ",<br />"))
		}

		//fmt.Println(line)
	}

	//fmt.Println(resultArchs)
}

func genMultipliedString(str string, multiplier int) string {
	result := ""
	for i := 0; i < multiplier; i++ {
		result += str
	}
	return result
}

func getLongestLength(strs []string) int {
	longest := 0

	for _, s := range strs {
		if len(s) > longest {
			longest = len(s)
		}
	}

	return longest
}