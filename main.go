package main

import (
	"flag"

	"app.io/pkg/k8s"
	apiv1 "k8s.io/api/core/v1"
)

var kubeconfig *string

func init() {
	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	flag.Parse()
}

func main() {

	client, err := k8s.GetKubernetesClient(*kubeconfig)
	if err != nil {
		panic(err)
	}

	k8s.DeploymentApp(*client, "default", "server-backend", "echo", "hub.dckr.ir/library/nginx:1.23", 1)
	k8s.DeploymentList(*client, apiv1.NamespaceAll)

}
