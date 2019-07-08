## Docker Images
|                          	| 386 	| amd64 	| arm 	| arm64 	| ppc64le 	| s390x 	|
|--------------------------	|:---:	|:-----:	|:---:	|:-----:	|:-------:	|:-----:	|
| cloud-controller-manager 	|  ✅  	|   ✅   	|  ✅  	|   ✅   	|    ✅    	|   ✅   	|
| conformance              	|  ✅  	|   ✅   	|  ✅  	|   ✅   	|    ✅    	|   ✅   	|
| hyperkube                	|  ✅  	|   ✅   	|  ✅  	|   ✅   	|    ✅    	|   ✅   	|
| kube-apiserver           	|  ✅  	|   ✅   	|  ✅  	|   ✅   	|    ✅    	|   ✅   	|
| kube-controller-manager  	|  ✅  	|   ✅   	|  ✅  	|   ✅   	|    ✅    	|   ✅   	|
| kube-proxy               	|  ✅  	|   ✅   	|  ✅  	|   ✅   	|    ✅    	|   ✅   	|
| kube-scheduler           	|  ✅  	|   ✅   	|  ✅  	|   ✅   	|    ✅    	|   ✅   	|

## Storage

### Binaries
|  | darwin/386 | darwin/amd64 | linux/386 | linux/amd64 | linux/arm | linux/arm64 | linux/ppc64le | linux/s390x | windows/386 | windows/amd64 |
|------------------------|----------|------------|---------|-----------|---------|-----------|-------------|-----------|-----------|-------------|
| apiextensions-apiserver |  |  |  | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 |  |  |
| cloud-controller-manager |  |  |  | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 |  |  |
| hyperkube |  |  |  | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 |  |  |
| kube-apiserver |  |  |  | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 |  |  |
| kube-controller-manager |  |  |  | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 |  |  |
| kube-proxy |  |  |  | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 |  | exe,<br />exe.md5,<br />exe.sha1,<br />exe.sha512 |
| kube-scheduler |  |  |  | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 | [binary_file],<br />docker_tag,<br />docker_tag.md5,<br />docker_tag.sha1,<br />docker_tag.sha512,<br />md5,<br />sha1,<br />sha512,<br />tar,<br />tar.md5,<br />tar.sha1,<br />tar.sha512 |  |  |
| kubeadm |  |  |  | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 |  | exe,<br />exe.md5,<br />exe.sha1,<br />exe.sha512 |
| kubectl | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 | exe,<br />exe.md5,<br />exe.sha1,<br />exe.sha512 | exe,<br />exe.md5,<br />exe.sha1,<br />exe.sha512 |
| kubelet |  |  |  | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 |  | exe,<br />exe.md5,<br />exe.sha1,<br />exe.sha512 |
| mounter |  |  |  | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 | [binary_file],<br />md5,<br />sha1,<br />sha512 |  |  |

### Extra files (for GCE)

There are also files: .md5, .sha1, .sha512 for every position below

| File 	|
|-----------------------------	|
| configure.sh 	|
| master.yaml 	|
| node.yaml 	|
| shutdown.sh 	|
|  |
| windows/common.psm1 	|
| windows/configure.ps1 	|
| windows/install-ssh.psm1 	|
| windows/k8s-node-setup.psm1 	|
| windows/user-profile.psm1 	|

### Tar archives

There are also files: .md5, .sha1, .sha512 for every position below

| File&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; | darwin/386 | darwin/amd64 | linux/386 | linux/amd64 | linux/arm | linux/arm64 | linux/ppc64le | linux/s390x | windows/386 | windows/amd64 | portable |
|----|:----------:|:------------:|:---------:|:-----------:|:---------:|:-----------:|:-------------:|:-----------:|:------------:|:-------------:|:--------:|
| kubernetes-client-**[system-arch]**.tar.gz | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |  |
| kubernetes-node-**[system-arch]**.tar.gz |  |  |  | ✅ | ✅ | ✅ | ✅ | ✅ |  | ✅ |  |
| kubernetes-server-**[system-arch]**.tar.gz |  |  |  | ✅ | ✅ | ✅ | ✅ | ✅ |  |  |  |
| kubernetes-test-**[system-arch]**.tar.gz |  | ✅ |  | ✅ | ✅ | ✅ | ✅ | ✅ |  | ✅ | ✅ |
|  |  |  |  |  |  |  |  |  |  |  |  |
| kubernetes-manifests.tar.gz |  |  |  |  |  |  |  |  |  |  |  |
| kubernetes-test.tar.gz |  |  |  |  |  |  |  |  |  |  |  |
| kubernetes-src.tar.gz |  |  |  |  |  |  |  |  |  |  |  |
| kubernetes.tar.gz |  |  |  |  |  |  |  |  |  |  |  |

##### [Content of `kubernetes-client-[system-arch].tar.gz` on example of `kubernetes-client-linux-386.tar.gz` (directories removed from list)](https://github.com/bartsmykla/k8s-artifacts/tree/master/content_of_tars#content-of-kubernetes-client-system-archtargz-on-example-of-kubernetes-client-linux-386targz-directories-removed-from-list)
##### [Content of `kubernetes-node-[system-arch].tar.gz` on example of `kubernetes-node-linux-amd64.tar.gz` (directories removed from list)](https://github.com/bartsmykla/k8s-artifacts/tree/master/content_of_tars#content-of-kubernetes-node-system-archtargz-on-example-of-kubernetes-node-linux-amd64targz-directories-removed-from-list)
##### [Content of `kubernetes-server-[system-arch].tar.gz` on example of `kubernetes-server-linux-amd64.tar.gz` (directories removed from list)](https://github.com/bartsmykla/k8s-artifacts/tree/master/content_of_tars#content-of-kubernetes-server-system-archtargz-on-example-of-kubernetes-server-linux-amd64targz-directories-removed-from-list)
##### [Content of `kubernetes-test-[system-arch].tar.gz` on example of `ubernetes-test-linux-amd64.tar.gz` (directories removed from list)](https://github.com/bartsmykla/k8s-artifacts/tree/master/content_of_tars#content-of-kubernetes-test-system-archtargz-on-example-of-kubernetes-test-linux-amd64targz-directories-removed-from-list)
##### [Content of `kubernetes-manifests.tar.gz` (directories removed from list)](https://github.com/bartsmykla/k8s-artifacts/tree/master/content_of_tars#content-of-kubernetes-manifeststargz-directories-removed-from-list)
##### [Content of `kubernetes-test.tar.gz` (directories removed from list)](content_of_tars/kubernetes-test_list-raw.txt)
##### [Content of `kubernetes-src.tar.gz`  (directories removed from list)](content_of_tars/kubernetes-src_list-raw.txt)
##### [Content of `kubernetes.tar.gz`  (directories removed from list)](content_of_tars/kubernetes_list-raw.txt)

## Packages

### RPMs
|  	| aarch64 	| armhfp 	| ppc64le 	| s390x 	| x86_64 	|
|------------------------------------------	|:-------:	|:------:	|:-------:	|:-----:	|:------:	|
| cri-tools-**[k8s-version]**.**[arch]**.rpm 	| ✅ 	| ✅ 	| ✅ 	| ✅ 	| ✅ 	|
| kubectl-**[k8s-version]**.**[arch]**.rpm 	| ✅ 	| ✅ 	| ✅ 	| ✅ 	| ✅ 	|
| kubernetes-cni-**[cni-version]**.**[arch]**.rpm 	| ✅ 	| ✅ 	| ✅ 	| ✅ 	| ✅ 	|
| kubeadm-**[k8s-version]**.**[arch]**.rpm 	| ✅ 	| ✅ 	| ✅ 	| ✅ 	| ✅ 	|
| kubelet-**[k8s-version]**.**[arch]**.rpm 	| ✅ 	| ✅ 	| ✅ 	| ✅ 	| ✅ 	|
|  	|  	|  	|  	|  	|  	|
| repodata/**[checksum]**-primary.sqlite.bz2 	| ✅ 	| ✅ 	| ✅ 	| ✅ 	| ✅ 	|
| repodata/**[checksum]**-filelists.sqlite.bz2 	| ✅ 	| ✅ 	| ✅ 	| ✅ 	| ✅ 	|
| repodata/**[checksum]**-primary.xml.gz 	| ✅ 	| ✅ 	| ✅ 	| ✅ 	| ✅ 	|
| repodata/**[checksum]**-other.sqlite.bz2 	| ✅ 	| ✅ 	| ✅ 	| ✅ 	| ✅ 	|
| repodata/**[checksum]**-other.xml.gz 	| ✅ 	| ✅ 	| ✅ 	| ✅ 	| ✅ 	|
| repodata/**[checksum]**-filelists.xml.gz 	| ✅ 	| ✅ 	| ✅ 	| ✅ 	| ✅ 	|
| repodata/repomd.xml 	| ✅ 	| ✅ 	| ✅ 	| ✅ 	| ✅ 	|