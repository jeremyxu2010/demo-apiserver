
/*
Copyright 2017 The Kubernetes Authors.

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

// Package demo api is the internal version of the API.
package demo

import (
	"github.com/kubernetes-incubator/apiserver-builder-alpha/pkg/builders"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
)

var _ rest.CreaterUpdater = &FooREST{}
var _ rest.Patcher = &FooREST{}
var _ rest.Scoper = &FooREST{}

// FooREST FooREST
// +k8s:deepcopy-gen=false
type FooREST struct {
	*genericregistry.Store
}

// NewFooREST Custom REST storage that delegates to the generated standard Registry
func NewFooREST(optsGetter generic.RESTOptionsGetter) rest.Storage {
	groupResource := schema.GroupResource{
		Group:    "demo",
		Resource: "foos",
	}
	strategy := &FooStrategy{builders.StorageStrategySingleton}
	store := &genericregistry.Store{
		NewFunc:                  func() runtime.Object { return &Foo{} },
		NewListFunc:              func() runtime.Object { return &FooList{} },
		DefaultQualifiedResource: groupResource,

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter}
	if err := store.CompleteWithOptions(options); err != nil {
		panic(err)
	}
	return &FooREST{store}
}

// FooStandardStorageBuilder FooStandardStorageBuilder
// +k8s:deepcopy-gen=false
type FooStandardStorageBuilder struct {
	optsGetter generic.RESTOptionsGetter
}

// GetStandardStorage StandardStorageProvider's interface method
func (b *FooStandardStorageBuilder) GetStandardStorage() rest.StandardStorage {
	return NewFooREST(b.optsGetter).(rest.StandardStorage)
}
