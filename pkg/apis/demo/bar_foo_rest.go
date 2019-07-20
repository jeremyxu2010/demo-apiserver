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

package demo

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
)

var _ rest.Connecter = &FooBarREST{}

// FooBarREST FooBar custom REST
// +k8s:deepcopy-gen=false
type FooBarREST struct {
	Registry FooRegistry
}

// Connect returns an http.Handler that will handle the request/response for a given API invocation.
// The provided responder may be used for common API responses. The responder will write both status
// code and body, so the ServeHTTP method should exit after invoking the responder. The Handler will
// be used for a single API request and then discarded. The Responder is guaranteed to write to the
// same http.ResponseWriter passed to ServeHTTP.
func (r *FooBarREST) Connect(ctx context.Context, name string, opts runtime.Object, responder rest.Responder) (http.Handler, error) {
	fooBarOptions, ok := opts.(*FooBarOptions)
	if !ok {
		err := fmt.Errorf("invalid options object: %#v", opts)
		return &errorHTTHandler{Err: err}, err
	}
	foo, err := r.Registry.GetFoo(ctx, name, &metav1.GetOptions{})
	if err != nil {
		return &errorHTTHandler{Err: err}, err
	}

	// Execute real business logic using foo and fooBarOptions
	fmt.Fprintf(ioutil.Discard, "%v, %v\n", fooBarOptions, foo)

	// Send Response
	fooBar := &FooBar{Message: "xxx"}
	return newFooBarHTTHandler(responder, fooBar), nil
}

// +k8s:deepcopy-gen=false
type errorHTTHandler struct {
	Responder rest.Responder
	Err       error
}

func (h *errorHTTHandler) ServeHTTP(http.ResponseWriter, *http.Request) {
	h.Responder.Error(h.Err)
}

// +k8s:deepcopy-gen=false
type fooBarHTTHandler struct {
	Responder rest.Responder
	FooBar    *FooBar
}

func (h *fooBarHTTHandler) ServeHTTP(http.ResponseWriter, *http.Request) {
	h.Responder.Object(200, h.FooBar)
}

func newFooBarHTTHandler(responder rest.Responder, fooBar *FooBar) http.Handler {
	return &fooBarHTTHandler{Responder: responder, FooBar: fooBar}
}

// NewConnectOptions returns an empty options object that will be used to pass
// options to the Connect method. If nil, then a nil options object is passed to
// Connect. It may return a bool and a string. If true, the value of the request
// path below the object will be included as the named string in the serialization
// of the runtime object.
func (r *FooBarREST) NewConnectOptions() (runtime.Object, bool, string) {
	return &FooBarOptions{}, false, ""
}

// ConnectMethods returns the list of HTTP methods handled by Connect
func (r *FooBarREST) ConnectMethods() []string {
	return []string{"POST"}
}

// New returns an empty object that can be used with Create and Update after request data has been put into it.
// This object must be a pointer type for use with Codec.DecodeInto([]byte, runtime.Object)
func (r *FooBarREST) New() runtime.Object {
	return &FooBarOptions{}
}

// NewFooBarREST create a new FooBarREST
func NewFooBarREST(optsGetter generic.RESTOptionsGetter) rest.Storage {
	return &FooBarREST{Registry: NewFooRegistry(&FooStandardStorageBuilder{optsGetter: optsGetter})}
}
