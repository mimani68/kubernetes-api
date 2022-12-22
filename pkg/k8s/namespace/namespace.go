package namespace

import (
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func CreateNamespace(client kubernetes.Clientset, namespace string) {
	if namespace == "" {
		namespace = apiv1.NamespaceDefault
	}
	namespaceClient := client.CoreV1().Namespaces()

}
