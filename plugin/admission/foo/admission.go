
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



package fooadmission

import (
	aggregatedadmission "github.com/jeremyxu2010/demo-apiserver/plugin/admission"
	aggregatedinformerfactory "github.com/jeremyxu2010/demo-apiserver/pkg/client/informers_generated/externalversions"
	aggregatedclientset "github.com/jeremyxu2010/demo-apiserver/pkg/client/clientset_generated/clientset"
	genericadmissioninitializer "k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/apiserver/pkg/admission"
)

var _ admission.Interface 											= &fooPlugin{}
var _ admission.MutationInterface 									= &fooPlugin{}
var _ admission.ValidationInterface 								= &fooPlugin{}
var _ genericadmissioninitializer.WantsExternalKubeInformerFactory 	= &fooPlugin{}
var _ genericadmissioninitializer.WantsExternalKubeClientSet 		= &fooPlugin{}
var _ aggregatedadmission.WantsAggregatedResourceInformerFactory 	= &fooPlugin{}
var _ aggregatedadmission.WantsAggregatedResourceClientSet 			= &fooPlugin{}

func NewFooPlugin() *fooPlugin {
	return &fooPlugin{
		Handler: admission.NewHandler(admission.Create, admission.Update),
	}
}

type fooPlugin struct {
	*admission.Handler
}

func (p *fooPlugin) ValidateInitialization() error {
	return nil
}

func (p *fooPlugin) Admit(a admission.Attributes) error {
	return nil
}

func (p *fooPlugin) Validate(a admission.Attributes) error {
	return nil
}

func (p *fooPlugin) SetAggregatedResourceInformerFactory(aggregatedinformerfactory.SharedInformerFactory) {}

func (p *fooPlugin) SetAggregatedResourceClientSet(aggregatedclientset.Interface) {}

func (p *fooPlugin) SetExternalKubeInformerFactory(informers.SharedInformerFactory) {}

func (p *fooPlugin) SetExternalKubeClientSet(kubernetes.Interface) {}
