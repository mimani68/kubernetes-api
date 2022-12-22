package container

import (
	corev1 "k8s.io/api/core/v1"
)

var SampleContainer = corev1.Container{
	Name:            "webserver",
	Image:           "nginx:latest",
	ImagePullPolicy: corev1.PullIfNotPresent,
	Args:            []string{"--config", "salam"},
	Env: []corev1.EnvVar{
		{
			Name: "POD_NAMESPACE",
			ValueFrom: &corev1.EnvVarSource{
				FieldRef: &corev1.ObjectFieldSelector{
					FieldPath: "metadata.namespace",
				},
			},
		},
		{
			Name:  "POD_RESTART",
			Value: "true",
		},
	},
	LivenessProbe: &corev1.Probe{
		// Handler: corev1.Handler{
		// 	Exec: &livenessExecActionController,
		// },
		InitialDelaySeconds: 100,
		TimeoutSeconds:      100,
		PeriodSeconds:       100,
		FailureThreshold:    100,
	},
	ReadinessProbe: &corev1.Probe{
		// Handler: corev1.Handler{
		// 	Exec: &readinessExecActionController,
		// },
		InitialDelaySeconds: 100,
		TimeoutSeconds:      100,
		PeriodSeconds:       100,
		FailureThreshold:    100,
	},
	SecurityContext: &corev1.SecurityContext{},
	Resources:       corev1.ResourceRequirements{},
}
