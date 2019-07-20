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

package main

import (
	"github.com/kubernetes-incubator/apiserver-builder-alpha/pkg/apiserver"
	"github.com/kubernetes-incubator/apiserver-builder-alpha/pkg/builders"
	openapinamer "k8s.io/apiserver/pkg/endpoints/openapi"
	genericapiserver "k8s.io/apiserver/pkg/server"

	// Make sure dep tools picks up these dependencies
	_ "github.com/go-openapi/loads"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kubernetes-incubator/apiserver-builder-alpha/pkg/cmd/server"
	_ "k8s.io/client-go/plugin/pkg/client/auth" // Enable cloud provider auth

	"github.com/jeremyxu2010/demo-apiserver/pkg/apis"
	"github.com/jeremyxu2010/demo-apiserver/pkg/openapi"
	_ "github.com/jeremyxu2010/demo-apiserver/plugin/admission/install"
)

func main() {
	version := "v0"
	openapidefs := openapi.GetOpenAPIDefinitions
	server.StartApiServer("/registry/jeremyxu2010.me", apis.GetAllApiBuilders(), openapidefs, "Api", version, func(apiServerConfig *apiserver.Config) error {
		apiServerConfig.RecommendedConfig.CorsAllowedOriginList = []string{"editor.swagger.io"}
		apiServerConfig.RecommendedConfig.OpenAPIConfig = genericapiserver.DefaultOpenAPIConfig(openapidefs, openapinamer.NewDefinitionNamer(builders.Scheme))
		apiServerConfig.RecommendedConfig.OpenAPIConfig.Info.Title = "Api"
		apiServerConfig.RecommendedConfig.OpenAPIConfig.Info.Version = version
		// apiServerConfig.RecommendedConfig.EnableSwaggerUI = true
		// apiServerConfig.RecommendedConfig.SwaggerConfig = genericapiserver.DefaultSwaggerConfig()
		return nil
	})
}
