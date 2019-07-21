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
	aggregatedclientset "github.com/jeremyxu2010/demo-apiserver/pkg/client/clientset_generated/clientset"
	aggregatedinformerfactory "github.com/jeremyxu2010/demo-apiserver/pkg/client/informers_generated/externalversions"
	aggregatedadmission "github.com/jeremyxu2010/demo-apiserver/plugin/admission"
	"k8s.io/apiserver/pkg/admission"
	genericadmissioninitializer "k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
)

var _ admission.Interface = &fooPlugin{}
var _ admission.MutationInterface = &fooPlugin{}
var _ admission.ValidationInterface = &fooPlugin{}
var _ genericadmissioninitializer.WantsExternalKubeInformerFactory = &fooPlugin{}
var _ genericadmissioninitializer.WantsExternalKubeClientSet = &fooPlugin{}
var _ aggregatedadmission.WantsAggregatedResourceInformerFactory = &fooPlugin{}
var _ aggregatedadmission.WantsAggregatedResourceClientSet = &fooPlugin{}

// NewFooPlugin NewFooPlugin
func NewFooPlugin() admission.Interface {
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

func (p *fooPlugin) Admit(_ admission.Attributes) error {
	return nil
}

func (p *fooPlugin) Validate(_ admission.Attributes) error {
	return nil
}

func (p *fooPlugin) SetAggregatedResourceInformerFactory(aggregatedinformerfactory.SharedInformerFactory) {
}

func (p *fooPlugin) SetAggregatedResourceClientSet(aggregatedclientset.Interface) {}

func (p *fooPlugin) SetExternalKubeInformerFactory(informers.SharedInformerFactory) {}

func (p *fooPlugin) SetExternalKubeClientSet(kubernetes.Interface) {}
