
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



package v1beta1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/jeremyxu2010/demo-apiserver/pkg/apis/demo/v1beta1"
	. "github.com/jeremyxu2010/demo-apiserver/pkg/client/clientset_generated/clientset/typed/demo/v1beta1"
)

var _ = Describe("Foo", func() {
	var instance Foo
	var expected Foo
	var client FooInterface

	BeforeEach(func() {
		instance = Foo{}
		instance.Name = "instance-1"

		expected = instance
	})

	AfterEach(func() {
		client.Delete(instance.Name, &metav1.DeleteOptions{})
	})

	Describe("when sending a bar request", func() {
		It("should return success", func() {
			client = cs.DemoV1beta1().Foos("foo-test-bar")
			_, err := client.Create(&instance)
			Expect(err).ShouldNot(HaveOccurred())

			bar := &FooBar{}
			bar.Name = instance.Name
			restClient := cs.DemoV1beta1().RESTClient()
			err = restClient.Post().Namespace("foo-test-bar").
				Name(instance.Name).
				Resource("foos").
				SubResource("bar").
				Body(bar).Do().Error()
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
