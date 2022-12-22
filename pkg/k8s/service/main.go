package service

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var AppService = &corev1.Service{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "app-service",
		Namespace: "default",
		Labels: map[string]string{
			"app": "ibm-cert-manager-webhook",
		},
	},
	Spec: corev1.ServiceSpec{
		Ports: []corev1.ServicePort{
			{
				Name:     "https",
				Port:     443,
				Protocol: "TCP",
				TargetPort: intstr.IntOrString{
					IntVal: 10250,
				},
			},
		},
		Selector: map[string]string{
			"app": "ibm-cert-manager-webhook",
		},
		Type: corev1.ServiceTypeClusterIP,
	},
}
