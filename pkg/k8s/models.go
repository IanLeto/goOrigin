package k8s

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ModelObjectMeta struct {
	Name         string
	GenerateName string
	Namespace    string
	Labels       map[string]string
}

func (m *ModelObjectMeta) ToK8sModel() *metav1.ObjectMeta {
	return &metav1.ObjectMeta{
		Name:         m.Name,
		GenerateName: m.GenerateName,
		Namespace:    m.Namespace,
		Labels:       m.Labels,
	}
}

type ModelObjectRef struct {
	Kind      string
	Namespace string
	Name      string
}

func (m *ModelObjectRef) ToK8sModel() *corev1.ObjectReference {
	// obj——ref 一种资源的引用，通过 ns + name + kind 来确定唯一资源
	return &corev1.ObjectReference{
		Kind:      m.Kind,
		Namespace: m.Namespace,
		Name:      m.Name,
	}
}

type ModelLocalObjectRef struct {
	Name string
}

func (m *ModelLocalObjectRef) ToK8sModel() *corev1.LocalObjectReference {
	return &corev1.LocalObjectReference{
		Name: m.Name,
	}
}

type ModelRole struct {
}

type ModelServiceAccount struct {
	Kind     string
	Metadata *ModelObjectMeta
}

type ModelRoleBinding struct {
}
