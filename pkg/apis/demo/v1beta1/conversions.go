package v1beta1

import (
	"github.com/jeremyxu2010/demo-apiserver/pkg/apis/demo"
	conversion "k8s.io/apimachinery/pkg/conversion"
)

//revive:disable
func Convert_v1beta1_FooBarOptions_To_demo_FooBarOptions(in *FooBarOptions, out *demo.FooBarOptions, s conversion.Scope) error {
	return autoConvert_v1beta1_FooBarOptions_To_demo_FooBarOptions(in, out, s)
}

//revive:enable
