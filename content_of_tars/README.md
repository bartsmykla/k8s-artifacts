##### Content of `kubernetes-client-[system-arch].tar.gz` on example of `kubernetes-client-linux-386.tar.gz` (directories removed from list):
```
➜ $ tar -ztvf kubernetes-client-linux-386.tar.gz
-rwxr-xr-x  0 root   root 37748128 Jul  5 16:26 kubernetes/client/bin/kubectl
```

##### Content of `kubernetes-node-[system-arch].tar.gz` on example of `kubernetes-node-linux-amd64.tar.gz` (directories removed from list):
```
➜ $ tar -ztvf kubernetes-node-linux-amd64.tar.gz
-rwxr-xr-x  0 root   root 39313856 Jul  5 16:26 kubernetes/node/bin/kubeadm
-rwxr-xr-x  0 root   root 36692480 Jul  5 16:26 kubernetes/node/bin/kube-proxy
-rwxr-xr-x  0 root   root 43035584 Jul  5 16:26 kubernetes/node/bin/kubectl
-rwxr-xr-x  0 root   root 113803872 Jul  5 16:26 kubernetes/node/bin/kubelet
-rw-r--r--  0 root   root   1171762 Jul  5 16:26 kubernetes/LICENSES
-rw-r--r--  0 root   root  29150046 Jul  5 16:26 kubernetes/kubernetes-src.tar.gz
```

##### Content of `kubernetes-server-[system-arch].tar.gz` on example of `kubernetes-server-linux-amd64.tar.gz` (directories removed from list):
```
➜ $ tar -ztvf kubernetes-server-linux-amd64.tar.gz
-rwxr-xr-x  0 root   root 110018784 Jul  5 16:30 kubernetes/server/bin/kube-controller-manager
-rw-r--r--  0 root   root 138259968 Jul  5 16:27 kubernetes/server/bin/cloud-controller-manager.tar
-rwxr-xr-x  0 root   root  39313856 Jul  5 16:30 kubernetes/server/bin/kubeadm
-rwxr-xr-x  0 root   root 195544384 Jul  5 16:30 kubernetes/server/bin/hyperkube
-rw-r--r--  0 root   root        16 Jul  5 16:27 kubernetes/server/bin/cloud-controller-manager.docker_tag
-rw-r--r--  0 root   root        16 Jul  5 16:27 kubernetes/server/bin/kube-apiserver.docker_tag
-rw-r--r--  0 root   root        16 Jul  5 16:27 kubernetes/server/bin/kube-scheduler.docker_tag
-rw-r--r--  0 root   root 153907712 Jul  5 16:27 kubernetes/server/bin/kube-controller-manager.tar
-rwxr-xr-x  0 root   root  38884704 Jul  5 16:30 kubernetes/server/bin/kube-scheduler
-rw-r--r--  0 root   root 203098112 Jul  5 16:27 kubernetes/server/bin/kube-apiserver.tar
-rw-r--r--  0 root   root        16 Jul  5 16:27 kubernetes/server/bin/kube-proxy.docker_tag
-rwxr-xr-x  0 root   root 159209248 Jul  5 16:30 kubernetes/server/bin/kube-apiserver
-rwxr-xr-x  0 root   root  36692480 Jul  5 16:30 kubernetes/server/bin/kube-proxy
-rwxr-xr-x  0 root   root  43035584 Jul  5 16:30 kubernetes/server/bin/kubectl
-rw-r--r--  0 root   root  83986944 Jul  5 16:27 kubernetes/server/bin/kube-proxy.tar
-rwxr-xr-x  0 root   root   1648224 Jul  5 16:30 kubernetes/server/bin/mounter
-rwxr-xr-x  0 root   root  94371200 Jul  5 16:30 kubernetes/server/bin/cloud-controller-manager
-rw-r--r--  0 root   root        16 Jul  5 16:27 kubernetes/server/bin/kube-controller-manager.docker_tag
-rw-r--r--  0 root   root  82773504 Jul  5 16:27 kubernetes/server/bin/kube-scheduler.tar
-rwxr-xr-x  0 root   root  43653824 Jul  5 16:30 kubernetes/server/bin/apiextensions-apiserver
-rwxr-xr-x  0 root   root 113803872 Jul  5 16:30 kubernetes/server/bin/kubelet
-rw-r--r--  0 root   root   1171762 Jul  5 16:30 kubernetes/LICENSES
-rw-r--r--  0 root   root  29150046 Jul  5 16:30 kubernetes/kubernetes-src.tar.gz
```

##### Content of `kubernetes-test-[system-arch].tar.gz` on example of `kubernetes-test-linux-amd64.tar.gz` (directories removed from list):
```
➜ $ tar -ztvf kubernetes-test-linux-amd64.tar.gz
-rwxr-xr-x  0 root   root 199571040 Jul  5 16:37 kubernetes/test/bin/genman
-rwxr-xr-x  0 root   root  42366560 Jul  5 16:37 kubernetes/test/bin/gendocs
-rwxr-xr-x  0 root   root 178091560 Jul  5 16:37 kubernetes/test/bin/e2e_node.test
-rwxr-xr-x  0 root   root 192029152 Jul  5 16:37 kubernetes/test/bin/genkubedocs
-rwxr-xr-x  0 root   root   5483328 Jul  5 16:37 kubernetes/test/bin/linkcheck
-rwxr-xr-x  0 root   root   4093952 Jul  5 16:37 kubernetes/test/bin/genswaggertypedocs
-rwxr-xr-x  0 root   root 124288648 Jul  5 16:37 kubernetes/test/bin/e2e.test
-rwxr-xr-x  0 root   root   8814784 Jul  5 16:37 kubernetes/test/bin/ginkgo
-rwxr-xr-x  0 root   root  42358368 Jul  5 16:37 kubernetes/test/bin/genyaml
-rwxr-xr-x  0 root   root 110000664 Jul  5 16:37 kubernetes/test/bin/kubemark
```

##### Content of `kubernetes-manifests.tar.gz` (directories removed from list):
```
➜ $ tar -ztvf kubernetes-manifests.tar.gz
-rw-r--r--  0 root   root      425 Jul  5 16:05 kubernetes/gci-trusty/ip-masq-agent/podsecuritypolicies/ip-masq-agent-psp-binding.yaml
-rw-r--r--  0 root   root     1857 Jul  5 16:05 kubernetes/gci-trusty/ip-masq-agent/ip-masq-agent.yaml
-rw-r--r--  0 root   root     1289 Jul  5 16:26 kubernetes/gci-trusty/abac-authz-policy.jsonl
-rw-r--r--  0 root   root     1116 Jul  5 16:05 kubernetes/gci-trusty/cluster-loadbalancing/glbc/default-svc-controller.yaml
-rw-r--r--  0 root   root      577 Jul  5 16:05 kubernetes/gci-trusty/cluster-loadbalancing/glbc/default-svc.yaml
-rw-r--r--  0 root   root     5746 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/calico-node-daemonset.yaml
-rw-r--r--  0 root   root     2078 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/typha-deployment.yaml
-rw-r--r--  0 root   root      185 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/calico-serviceaccount.yaml
-rw-r--r--  0 root   root      378 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/hostendpoints-crd.yaml
-rw-r--r--  0 root   root      394 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/globalnetworksets-crd.yaml
-rw-r--r--  0 root   root      189 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/typha-horizontal-autoscaler-serviceaccount.yaml
-rw-r--r--  0 root   root      354 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/ippool-crd.yaml
-rw-r--r--  0 root   root      408 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/globalnetworkpolicy-crd.yaml
-rw-r--r--  0 root   root      259 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/typha-horizontal-autoscaler-clusterrole.yaml
-rw-r--r--  0 root   root     1145 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/calico-node-vertical-autoscaler-deployment.yaml
-rw-r--r--  0 root   root      398 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/globalfelixconfig-crd.yaml
-rw-r--r--  0 root   root      451 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/calico-node-vertical-autoscaler-configmap.yaml
-rw-r--r--  0 root   root      394 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/bgpconfigurations-crd.yaml
-rw-r--r--  0 root   root      452 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/typha-horizontal-autoscaler-configmap.yaml
-rw-r--r--  0 root   root      402 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/felixconfigurations-crd.yaml
-rw-r--r--  0 root   root      189 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/typha-vertical-autoscaler-serviceaccount.yaml
-rw-r--r--  0 root   root      456 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/typha-vertical-autoscaler-configmap.yaml
-rw-r--r--  0 root   root      411 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/podsecuritypolicies/calico-node-psp-binding.yaml
-rw-r--r--  0 root   root      377 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/typha-horizontal-autoscaler-role.yaml
-rw-r--r--  0 root   root     1153 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/typha-vertical-autoscaler-deployment.yaml
-rw-r--r--  0 root   root      345 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/calico-clusterrolebinding.yaml
-rw-r--r--  0 root   root      374 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/typha-horizontal-autoscaler-rolebinding.yaml
-rw-r--r--  0 root   root      356 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/typha-service.yaml
-rw-r--r--  0 root   root      363 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/typha-vertical-autoscaler-clusterrolebinding.yaml
-rw-r--r--  0 root   root      402 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/clusterinformations-crd.yaml
-rw-r--r--  0 root   root      363 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/typha-horizontal-autoscaler-clusterrolebinding.yaml
-rw-r--r--  0 root   root      364 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/calico-cpva-clusterrole.yaml
-rw-r--r--  0 root   root     1650 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/calico-clusterrole.yaml
-rw-r--r--  0 root   root      390 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/globalbgpconfig-crd.yaml
-rw-r--r--  0 root   root     1137 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/typha-horizontal-autoscaler-deployment.yaml
-rw-r--r--  0 root   root      387 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/networkpolicies-crd.yaml
-rw-r--r--  0 root   root      190 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/calico-cpva-serviceaccount.yaml
-rw-r--r--  0 root   root      366 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/calico-cpva-clusterrolebinding.yaml
-rw-r--r--  0 root   root      349 Jul  5 16:05 kubernetes/gci-trusty/calico-policy-controller/typha-vertical-autoscaler-clusterrole.yaml
-rw-r--r--  0 root   root      427 Jul  5 16:05 kubernetes/gci-trusty/metadata-proxy/gce/podsecuritypolicies/metadata-proxy-psp-binding.yaml
-rw-r--r--  0 root   root     2768 Jul  5 16:05 kubernetes/gci-trusty/metadata-proxy/gce/metadata-proxy.yaml
-rw-r--r--  0 root   root     3368 Jul  5 16:05 kubernetes/gci-trusty/dns-horizontal-autoscaler/dns-horizontal-autoscaler.yaml
-rw-r--r--  0 root   root     2508 Jul  5 16:05 kubernetes/gci-trusty/node-termination-handler/daemonset.yaml
-rw-r--r--  0 root   root      398 Jul  5 16:05 kubernetes/gci-trusty/metrics-server/auth-delegator.yaml
-rw-r--r--  0 root   root      801 Jul  5 16:05 kubernetes/gci-trusty/metrics-server/resource-reader.yaml
-rw-r--r--  0 root   root      336 Jul  5 16:05 kubernetes/gci-trusty/metrics-server/metrics-server-service.yaml
-rw-r--r--  0 root   root      388 Jul  5 16:05 kubernetes/gci-trusty/metrics-server/metrics-apiservice.yaml
-rw-r--r--  0 root   root      419 Jul  5 16:05 kubernetes/gci-trusty/metrics-server/auth-reader.yaml
-rw-r--r--  0 root   root     3255 Jul  5 16:05 kubernetes/gci-trusty/metrics-server/metrics-server-deployment.yaml
-rw-r--r--  0 root   root      456 Jul  5 16:26 kubernetes/gci-trusty/etcd-empty-dir-cleanup.yaml
-rw-r--r--  0 root   root     3394 Jul  5 16:26 kubernetes/gci-trusty/etcd.manifest
-rw-r--r--  0 root   root     2275 Jul  5 16:05 kubernetes/gci-trusty/rbac/cluster-loadbalancing/glbc/roles.yaml
-rw-r--r--  0 root   root      647 Jul  5 16:05 kubernetes/gci-trusty/rbac/cluster-loadbalancing/glbc/user-rolebindings.yaml
-rw-r--r--  0 root   root      782 Jul  5 16:05 kubernetes/gci-trusty/rbac/legacy-kubelet-user-disable/kubelet-binding.yaml
-rw-r--r--  0 root   root      565 Jul  5 16:05 kubernetes/gci-trusty/rbac/legacy-kubelet-user/kubelet-binding.yaml
-rw-r--r--  0 root   root     2059 Jul  5 16:05 kubernetes/gci-trusty/rbac/cluster-autoscaler/cluster-autoscaler-rbac.yaml
-rw-r--r--  0 root   root      338 Jul  5 16:05 kubernetes/gci-trusty/rbac/kubelet-api-auth/kubelet-api-admin-role.yaml
-rw-r--r--  0 root   root      427 Jul  5 16:05 kubernetes/gci-trusty/rbac/kubelet-api-auth/kube-apiserver-kubelet-api-admin-binding.yaml
-rw-r--r--  0 root   root     1399 Jul  5 16:05 kubernetes/gci-trusty/rbac/kubelet-cert-rotation/kubelet-certificate-management.yaml
-rw-r--r--  0 root   root      397 Jul  5 16:05 kubernetes/gci-trusty/node-problem-detector/standalone/npd-binding.yaml
-rw-r--r--  0 root   root      423 Jul  5 16:05 kubernetes/gci-trusty/node-problem-detector/podsecuritypolicies/npd-psp-binding.yaml
-rw-r--r--  0 root   root     2418 Jul  5 16:05 kubernetes/gci-trusty/node-problem-detector/npd.yaml
-rw-r--r--  0 root   root     4721 Jul  5 16:05 kubernetes/gci-trusty/cluster-monitoring/googleinfluxdb/heapster-controller-combined.yaml
-rw-r--r--  0 root   root      492 Jul  5 16:05 kubernetes/gci-trusty/cluster-monitoring/influxdb/grafana-service.yaml
-rw-r--r--  0 root   root     2738 Jul  5 16:05 kubernetes/gci-trusty/cluster-monitoring/influxdb/influxdb-grafana-controller.yaml
-rw-r--r--  0 root   root      302 Jul  5 16:05 kubernetes/gci-trusty/cluster-monitoring/influxdb/heapster-service.yaml
-rw-r--r--  0 root   root      393 Jul  5 16:05 kubernetes/gci-trusty/cluster-monitoring/influxdb/influxdb-service.yaml
-rw-r--r--  0 root   root     4712 Jul  5 16:05 kubernetes/gci-trusty/cluster-monitoring/influxdb/heapster-controller.yaml
-rw-r--r--  0 root   root     1123 Jul  5 16:05 kubernetes/gci-trusty/cluster-monitoring/heapster-rbac.yaml
-rw-r--r--  0 root   root      302 Jul  5 16:05 kubernetes/gci-trusty/cluster-monitoring/stackdriver/heapster-service.yaml
-rw-r--r--  0 root   root     4221 Jul  5 16:05 kubernetes/gci-trusty/cluster-monitoring/stackdriver/heapster-controller.yaml
-rw-r--r--  0 root   root      302 Jul  5 16:05 kubernetes/gci-trusty/cluster-monitoring/standalone/heapster-service.yaml
-rw-r--r--  0 root   root     3009 Jul  5 16:05 kubernetes/gci-trusty/cluster-monitoring/standalone/heapster-controller.yaml
-rw-r--r--  0 root   root      302 Jul  5 16:05 kubernetes/gci-trusty/cluster-monitoring/google/heapster-service.yaml
-rw-r--r--  0 root   root     4632 Jul  5 16:05 kubernetes/gci-trusty/cluster-monitoring/google/heapster-controller.yaml
-rw-r--r--  0 root   root     4619 Jul  5 16:05 kubernetes/gci-trusty/dns/coredns/coredns.yaml.in
-rw-r--r--  0 root   root     6451 Jul  5 16:05 kubernetes/gci-trusty/dns/kube-dns/kube-dns.yaml.in
-rw-r--r--  0 root   root     3922 Jul  5 16:05 kubernetes/gci-trusty/dns/nodelocaldns/nodelocaldns.yaml
-rw-r--r--  0 root   root      166 Jul  5 16:05 kubernetes/gci-trusty/limit-range/limit-range.yaml
-rw-r--r--  0 root   root     4037 Jul  5 16:26 kubernetes/gci-trusty/kube-apiserver.manifest
-rw-r--r--  0 root   root     1457 Jul  5 16:05 kubernetes/gci-trusty/loadbalancing/cloud-provider-role.yaml
-rw-r--r--  0 root   root      685 Jul  5 16:05 kubernetes/gci-trusty/loadbalancing/cloud-provider-binding.yaml
-rw-r--r--  0 root   root     2183 Jul  5 16:05 kubernetes/gci-trusty/prometheus/alertmanager-deployment.yaml
-rw-r--r--  0 root   root      319 Jul  5 16:05 kubernetes/gci-trusty/prometheus/alertmanager-pvc.yaml
-rw-r--r--  0 root   root     2239 Jul  5 16:05 kubernetes/gci-trusty/prometheus/kube-state-metrics-rbac.yaml
-rw-r--r--  0 root   root      371 Jul  5 16:05 kubernetes/gci-trusty/prometheus/alertmanager-service.yaml
-rw-r--r--  0 root   root     3063 Jul  5 16:05 kubernetes/gci-trusty/prometheus/prometheus-statefulset.yaml
-rw-r--r--  0 root   root      402 Jul  5 16:05 kubernetes/gci-trusty/prometheus/alertmanager-configmap.yaml
-rw-r--r--  0 root   root     5440 Jul  5 16:05 kubernetes/gci-trusty/prometheus/prometheus-configmap.yaml
-rw-r--r--  0 root   root     1069 Jul  5 16:05 kubernetes/gci-trusty/prometheus/prometheus-rbac.yaml
-rw-r--r--  0 root   root      424 Jul  5 16:05 kubernetes/gci-trusty/prometheus/node-exporter-service.yaml
-rw-r--r--  0 root   root      347 Jul  5 16:05 kubernetes/gci-trusty/prometheus/prometheus-service.yaml
-rw-r--r--  0 root   root      506 Jul  5 16:05 kubernetes/gci-trusty/prometheus/kube-state-metrics-service.yaml
-rw-r--r--  0 root   root     2379 Jul  5 16:05 kubernetes/gci-trusty/prometheus/kube-state-metrics-deployment.yaml
-rw-r--r--  0 root   root     1577 Jul  5 16:05 kubernetes/gci-trusty/podsecuritypolicies/unprivileged-addon.yaml
-rw-r--r--  0 root   root      572 Jul  5 16:05 kubernetes/gci-trusty/podsecuritypolicies/persistent-volume-binder-role.yaml
-rw-r--r--  0 root   root      742 Jul  5 16:05 kubernetes/gci-trusty/podsecuritypolicies/node-binding.yaml
-rw-r--r--  0 root   root      535 Jul  5 16:05 kubernetes/gci-trusty/podsecuritypolicies/kube-system-binding.yaml
-rw-r--r--  0 root   root      643 Jul  5 16:05 kubernetes/gci-trusty/podsecuritypolicies/persistent-volume-binder-binding.yaml
-rw-r--r--  0 root   root      364 Jul  5 16:05 kubernetes/gci-trusty/podsecuritypolicies/unprivileged-addon-role.yaml
-rw-r--r--  0 root   root      330 Jul  5 16:05 kubernetes/gci-trusty/podsecuritypolicies/privileged-role.yaml
-rw-r--r--  0 root   root      791 Jul  5 16:05 kubernetes/gci-trusty/podsecuritypolicies/privileged.yaml
-rw-r--r--  0 root   root      425 Jul  5 16:05 kubernetes/gci-trusty/podsecuritypolicies/kube-proxy-binding.yaml
-rw-r--r--  0 root   root      878 Jul  5 16:05 kubernetes/gci-trusty/podsecuritypolicies/persistent-volume-binder.yaml
-rw-r--r--  0 root   root      551 Jul  5 16:05 kubernetes/gci-trusty/dashboard/dashboard-secret.yaml
-rw-r--r--  0 root   root      322 Jul  5 16:05 kubernetes/gci-trusty/dashboard/dashboard-service.yaml
-rw-r--r--  0 root   root     1353 Jul  5 16:05 kubernetes/gci-trusty/dashboard/dashboard-rbac.yaml
-rw-r--r--  0 root   root      264 Jul  5 16:05 kubernetes/gci-trusty/dashboard/dashboard-configmap.yaml
-rw-r--r--  0 root   root     1780 Jul  5 16:05 kubernetes/gci-trusty/dashboard/dashboard-controller.yaml
-rw-r--r--  0 root   root     3513 Jul  5 16:26 kubernetes/gci-trusty/cluster-autoscaler.manifest
-rw-r--r--  0 root   root      354 Jul  5 16:05 kubernetes/gci-trusty/fluentd-elasticsearch/kibana-service.yaml
-rw-r--r--  0 root   root      382 Jul  5 16:05 kubernetes/gci-trusty/fluentd-elasticsearch/es-service.yaml
-rw-r--r--  0 root   root     1321 Jul  5 16:05 kubernetes/gci-trusty/fluentd-elasticsearch/kibana-deployment.yaml
-rw-r--r--  0 root   root     2492 Jul  5 16:05 kubernetes/gci-trusty/fluentd-elasticsearch/fluentd-es-ds.yaml
-rw-r--r--  0 root   root      399 Jul  5 16:05 kubernetes/gci-trusty/fluentd-elasticsearch/podsecuritypolicies/es-psp-binding.yaml
-rw-r--r--  0 root   root     2624 Jul  5 16:05 kubernetes/gci-trusty/fluentd-elasticsearch/es-statefulset.yaml
-rw-r--r--  0 root   root    16124 Jul  5 16:05 kubernetes/gci-trusty/fluentd-elasticsearch/fluentd-es-configmap.yaml
-rw-r--r--  0 root   root      488 Jul  5 16:05 kubernetes/gci-trusty/kube-proxy/kube-proxy-rbac.yaml
-rw-r--r--  0 root   root     2060 Jul  5 16:05 kubernetes/gci-trusty/kube-proxy/kube-proxy-ds.yaml
-rw-r--r--  0 root   root     4512 Jul  5 16:05 kubernetes/gci-trusty/fluentd-gcp/fluentd-gcp-ds.yaml
-rw-r--r--  0 root   root     1015 Jul  5 16:05 kubernetes/gci-trusty/fluentd-gcp/scaler-rbac.yaml
-rw-r--r--  0 root   root      323 Jul  5 16:05 kubernetes/gci-trusty/fluentd-gcp/scaler-policy.yaml
-rw-r--r--  0 root   root    16675 Jul  5 16:05 kubernetes/gci-trusty/fluentd-gcp/fluentd-gcp-configmap-old.yaml
-rw-r--r--  0 root   root      190 Jul  5 16:05 kubernetes/gci-trusty/fluentd-gcp/fluentd-gcp-ds-sa.yaml
-rw-r--r--  0 root   root     1281 Jul  5 16:05 kubernetes/gci-trusty/fluentd-gcp/podsecuritypolicies/event-exporter-psp.yaml
-rw-r--r--  0 root   root      356 Jul  5 16:05 kubernetes/gci-trusty/fluentd-gcp/podsecuritypolicies/event-exporter-psp-role.yaml
-rw-r--r--  0 root   root     1189 Jul  5 16:05 kubernetes/gci-trusty/fluentd-gcp/podsecuritypolicies/fluentd-gcp-psp.yaml
-rw-r--r--  0 root   root      427 Jul  5 16:05 kubernetes/gci-trusty/fluentd-gcp/podsecuritypolicies/event-exporter-psp-binding.yaml
-rw-r--r--  0 root   root      350 Jul  5 16:05 kubernetes/gci-trusty/fluentd-gcp/podsecuritypolicies/fluentd-gcp-psp-role.yaml
-rw-r--r--  0 root   root      415 Jul  5 16:05 kubernetes/gci-trusty/fluentd-gcp/podsecuritypolicies/fluentd-gcp-psp-binding.yaml
-rw-r--r--  0 root   root      968 Jul  5 16:05 kubernetes/gci-trusty/fluentd-gcp/scaler-deployment.yaml
-rw-r--r--  0 root   root     2329 Jul  5 16:05 kubernetes/gci-trusty/fluentd-gcp/event-exporter.yaml
-rw-r--r--  0 root   root    18765 Jul  5 16:05 kubernetes/gci-trusty/fluentd-gcp/fluentd-gcp-configmap.yaml
-rw-r--r--  0 root   root     2698 Jul  5 16:26 kubernetes/gci-trusty/kube-controller-manager.manifest
-rw-r--r--  0 root   root      685 Jul  5 16:05 kubernetes/gci-trusty/metadata-agent/stackdriver/metadata-agent-rbac.yaml
-rw-r--r--  0 root   root      433 Jul  5 16:05 kubernetes/gci-trusty/metadata-agent/stackdriver/podsecuritypolicies/metadata-agent-psp-binding.yaml
-rw-r--r--  0 root   root     3493 Jul  5 16:05 kubernetes/gci-trusty/metadata-agent/stackdriver/metadata-agent.yaml
-rw-r--r--  0 root   root     2464 Jul  5 16:26 kubernetes/gci-trusty/glbc.manifest
-rw-r--r--  0 root   root     1493 Jul  5 16:26 kubernetes/gci-trusty/kube-scheduler.manifest
-rw-r--r--  0 root   root     1658 Jul  5 16:05 kubernetes/gci-trusty/device-plugins/nvidia-gpu/daemonset.yaml
-rw-r--r--  0 root   root     1315 Jul  5 16:26 kubernetes/gci-trusty/kube-addon-manager.yaml
-rw-r--r--  0 root   root   121915 Jul  5 16:26 kubernetes/gci-trusty/gci-configure-helper.sh
-rw-r--r--  0 root   root     2149 Jul  5 16:26 kubernetes/gci-trusty/kube-proxy.manifest
-rw-r--r--  0 root   root      261 Jul  5 16:05 kubernetes/gci-trusty/storage-class/aws/default.yaml
-rw-r--r--  0 root   root      269 Jul  5 16:05 kubernetes/gci-trusty/storage-class/local/default.yaml
-rw-r--r--  0 root   root      241 Jul  5 16:05 kubernetes/gci-trusty/storage-class/openstack/default.yaml
-rw-r--r--  0 root   root      245 Jul  5 16:05 kubernetes/gci-trusty/storage-class/azure/default.yaml
-rw-r--r--  0 root   root      228 Jul  5 16:05 kubernetes/gci-trusty/storage-class/vsphere/default.yaml
-rw-r--r--  0 root   root      300 Jul  5 16:05 kubernetes/gci-trusty/storage-class/gce/default.yaml
-rw-r--r--  0 root   root     4120 Jul  5 16:26 kubernetes/gci-trusty/health-monitor.sh
```

##### [Content of `kubernetes-test.tar.gz` (directories removed from list)](kubernetes-test_list-raw.txt)
##### [Content of `kubernetes-src.tar.gz`  (directories removed from list)](kubernetes-src_list-raw.txt)
##### [Content of `kubernetes.tar.gz`  (directories removed from list)](kubernetes_list-raw.txt)