package main

import (
	"flag"

	k8sPkg "app.io/pkg/k8s"
)

var kubeconfig *string

func init() {
	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	flag.Parse()
}

func main() {

	client, err := k8sPkg.GetKubernetesClient(*kubeconfig)
	if err != nil {
		panic(err)
	}

	k8sPkg.DeploymentList(*client)

}
