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

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package install

import (
	aggregatedclientset "github.com/jeremyxu2010/demo-apiserver/pkg/client/clientset_generated/clientset"
	aggregatedinformerfactory "github.com/jeremyxu2010/demo-apiserver/pkg/client/informers_generated/externalversions"
	intializer "github.com/jeremyxu2010/demo-apiserver/plugin/admission"
	. "github.com/jeremyxu2010/demo-apiserver/plugin/admission/foo"
	"github.com/kubernetes-incubator/apiserver-builder-alpha/pkg/cmd/server"
	"k8s.io/apiserver/pkg/admission"
	genericserver "k8s.io/apiserver/pkg/server"
	"k8s.io/client-go/rest"
)

func init() {
	server.AggregatedAdmissionInitializerGetter = GetAggregatedResourceAdmissionControllerInitializer
	server.AggregatedAdmissionPlugins["Foo"] = NewFooPlugin()
}

func GetAggregatedResourceAdmissionControllerInitializer(config *rest.Config) (admission.PluginInitializer, genericserver.PostStartHookFunc) {
	// init aggregated resource clients
	aggregatedResourceClient := aggregatedclientset.NewForConfigOrDie(config)
	aggregatedInformerFactory := aggregatedinformerfactory.NewSharedInformerFactory(aggregatedResourceClient, 0)
	aggregatedResourceInitializer := intializer.New(aggregatedResourceClient, aggregatedInformerFactory)

	return aggregatedResourceInitializer, func(context genericserver.PostStartHookContext) error {
		aggregatedInformerFactory.Start(context.StopCh)
		return nil
	}
}
