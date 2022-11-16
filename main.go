package main

import (
	"flag"

	"app.io/pkg/k8s"
	k8sDeployment "app.io/pkg/k8s/deployment"
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

	k8sDeployment.List(*client, apiv1.NamespaceAll)
	k8sDeployment.PerformDeployment(*client, "default", "server-backend", "echo", "hub.dckr.ir/library/nginx:1.23", 1)
	k8sDeployment.Scale(*client, "server-backend-deployment", "default", 5)
	// k8sDeployment.ChangeImage(*client, "server-backend-deployment", "default", "nginx:latest")
	// k8sDeployment.Delete(*client, "server-backend-deployment", "default")

}
