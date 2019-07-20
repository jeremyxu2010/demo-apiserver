package demo

import (
	"github.com/kubernetes-incubator/apiserver-builder-alpha/pkg/builders"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	builders.Scheme.AddKnownTypes(SchemeGroupVersion, &FooBarOptions{})
	builders.ParameterScheme.AddKnownTypes(SchemeGroupVersion, &FooBarOptions{})
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooBarOptions FooBar request options
type FooBarOptions struct {
	metav1.TypeMeta
	Arg1 string `json:"arg1"`
}
