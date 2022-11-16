package deployment

import (
	"context"
	"fmt"

	"app.io/pkg/k8s"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/retry"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func PerformDeployment(client kubernetes.Clientset, namespace, appName, image string, replica int) {
	if namespace == "" {
		namespace = apiv1.NamespaceDefault
	}
	deploymentsClient := client.AppsV1().Deployments(namespace)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: appName,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: k8s.Int32Ptr(replica),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": appName,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": appName,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: image,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	fmt.Println("Creating deployment...")
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
}

func update(client kubernetes.Clientset, namespace, deploymentName string, newResource *appsv1.Deployment) {
	if namespace == "" {
		namespace = apiv1.NamespaceDefault
	}
	deploymentsClient := client.AppsV1().Deployments(namespace)
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		_, updateErr := deploymentsClient.Update(context.TODO(), newResource, metav1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		panic(fmt.Errorf("Update failed: %v", retryErr))
	}
	fmt.Println("Updated deployment...")
}

func Scale(client kubernetes.Clientset, namespace, deploymentName string, number int) {
	deploymentsClient := client.AppsV1().Deployments(namespace)
	updatedResource, getErr := deploymentsClient.Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if getErr != nil {
		panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
	}
	updatedResource.Spec.Replicas = k8s.Int32Ptr(number)
	update(client, namespace, deploymentName, updatedResource)
}

func ChangeImage(client kubernetes.Clientset, namespace, deploymentName string, imageTag string) {
	deploymentsClient := client.AppsV1().Deployments(namespace)
	updatedResource, getErr := deploymentsClient.Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if getErr != nil {
		panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
	}
	updatedResource.Spec.Template.Spec.Containers[0].Image = imageTag
	update(client, namespace, deploymentName, updatedResource)
}

func Delete(client kubernetes.Clientset, namespace, deploymentName string) {
	if namespace == "" {
		namespace = apiv1.NamespaceDefault
	}
	deploymentsClient := client.AppsV1().Deployments(namespace)
	deletePolicy := metav1.DeletePropagationForeground
	if err := deploymentsClient.Delete(context.TODO(), deploymentName, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	fmt.Println("Deleted deployment.")
}

func List(client kubernetes.Clientset, namespace string) {
	if namespace == "" {
		namespace = apiv1.NamespaceAll
	}
	deploymentsClient := client.AppsV1().Deployments(namespace)

	fmt.Printf("Listing deployments in namespace %s:\n", namespace)
	list, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range list.Items {
		fmt.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
	}
}
