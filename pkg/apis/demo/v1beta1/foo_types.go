/*
Copyright YEAR The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Foo Foo
// +k8s:openapi-gen=true
// +resource:path=foos,strategy=FooStrategy,rest=FooREST
// +subresource:request=FooBar,path=bar,kind=FooBar,rest=FooBarREST
type Foo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FooSpec   `json:"spec,omitempty"`
	Status FooStatus `json:"status,omitempty"`
}

// FooSpec defines the desired state of Foo
type FooSpec struct {
}

// FooStatus defines the observed state of Foo
type FooStatus struct {
}

// // DefaultingFunction sets default Foo field values
// func (FooSchemeFns) DefaultingFunction(o interface{}) {
// 	obj := o.(*Foo)
// 	// set default field values here
// 	log.Printf("Defaulting fields for Foo %s\n", obj.Name)
// }
