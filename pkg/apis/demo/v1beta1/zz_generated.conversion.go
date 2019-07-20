// +build !ignore_autogenerated

/*
Copyright 2019 The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	unsafe "unsafe"

	demo "github.com/jeremyxu2010/demo-apiserver/pkg/apis/demo"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Foo)(nil), (*demo.Foo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Foo_To_demo_Foo(a.(*Foo), b.(*demo.Foo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*demo.Foo)(nil), (*Foo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_demo_Foo_To_v1beta1_Foo(a.(*demo.Foo), b.(*Foo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*FooBar)(nil), (*demo.FooBar)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_FooBar_To_demo_FooBar(a.(*FooBar), b.(*demo.FooBar), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*demo.FooBar)(nil), (*FooBar)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_demo_FooBar_To_v1beta1_FooBar(a.(*demo.FooBar), b.(*FooBar), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*FooBarList)(nil), (*demo.FooBarList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_FooBarList_To_demo_FooBarList(a.(*FooBarList), b.(*demo.FooBarList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*demo.FooBarList)(nil), (*FooBarList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_demo_FooBarList_To_v1beta1_FooBarList(a.(*demo.FooBarList), b.(*FooBarList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*FooBarOptions)(nil), (*demo.FooBarOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_FooBarOptions_To_demo_FooBarOptions(a.(*FooBarOptions), b.(*demo.FooBarOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*demo.FooBarOptions)(nil), (*FooBarOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_demo_FooBarOptions_To_v1beta1_FooBarOptions(a.(*demo.FooBarOptions), b.(*FooBarOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*FooList)(nil), (*demo.FooList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_FooList_To_demo_FooList(a.(*FooList), b.(*demo.FooList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*demo.FooList)(nil), (*FooList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_demo_FooList_To_v1beta1_FooList(a.(*demo.FooList), b.(*FooList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*FooSpec)(nil), (*demo.FooSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_FooSpec_To_demo_FooSpec(a.(*FooSpec), b.(*demo.FooSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*demo.FooSpec)(nil), (*FooSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_demo_FooSpec_To_v1beta1_FooSpec(a.(*demo.FooSpec), b.(*FooSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*FooStatus)(nil), (*demo.FooStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_FooStatus_To_demo_FooStatus(a.(*FooStatus), b.(*demo.FooStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*demo.FooStatus)(nil), (*FooStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_demo_FooStatus_To_v1beta1_FooStatus(a.(*demo.FooStatus), b.(*FooStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*FooBarOptions)(nil), (*demo.FooBarOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_FooBarOptions_To_demo_FooBarOptions(a.(*FooBarOptions), b.(*demo.FooBarOptions), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_Foo_To_demo_Foo(in *Foo, out *demo.Foo, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_FooSpec_To_demo_FooSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_FooStatus_To_demo_FooStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_Foo_To_demo_Foo is an autogenerated conversion function.
func Convert_v1beta1_Foo_To_demo_Foo(in *Foo, out *demo.Foo, s conversion.Scope) error {
	return autoConvert_v1beta1_Foo_To_demo_Foo(in, out, s)
}

func autoConvert_demo_Foo_To_v1beta1_Foo(in *demo.Foo, out *Foo, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_demo_FooSpec_To_v1beta1_FooSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_demo_FooStatus_To_v1beta1_FooStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_demo_Foo_To_v1beta1_Foo is an autogenerated conversion function.
func Convert_demo_Foo_To_v1beta1_Foo(in *demo.Foo, out *Foo, s conversion.Scope) error {
	return autoConvert_demo_Foo_To_v1beta1_Foo(in, out, s)
}

func autoConvert_v1beta1_FooBar_To_demo_FooBar(in *FooBar, out *demo.FooBar, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Message = in.Message
	return nil
}

// Convert_v1beta1_FooBar_To_demo_FooBar is an autogenerated conversion function.
func Convert_v1beta1_FooBar_To_demo_FooBar(in *FooBar, out *demo.FooBar, s conversion.Scope) error {
	return autoConvert_v1beta1_FooBar_To_demo_FooBar(in, out, s)
}

func autoConvert_demo_FooBar_To_v1beta1_FooBar(in *demo.FooBar, out *FooBar, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Message = in.Message
	return nil
}

// Convert_demo_FooBar_To_v1beta1_FooBar is an autogenerated conversion function.
func Convert_demo_FooBar_To_v1beta1_FooBar(in *demo.FooBar, out *FooBar, s conversion.Scope) error {
	return autoConvert_demo_FooBar_To_v1beta1_FooBar(in, out, s)
}

func autoConvert_v1beta1_FooBarList_To_demo_FooBarList(in *FooBarList, out *demo.FooBarList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]demo.FooBar)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_FooBarList_To_demo_FooBarList is an autogenerated conversion function.
func Convert_v1beta1_FooBarList_To_demo_FooBarList(in *FooBarList, out *demo.FooBarList, s conversion.Scope) error {
	return autoConvert_v1beta1_FooBarList_To_demo_FooBarList(in, out, s)
}

func autoConvert_demo_FooBarList_To_v1beta1_FooBarList(in *demo.FooBarList, out *FooBarList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]FooBar)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_demo_FooBarList_To_v1beta1_FooBarList is an autogenerated conversion function.
func Convert_demo_FooBarList_To_v1beta1_FooBarList(in *demo.FooBarList, out *FooBarList, s conversion.Scope) error {
	return autoConvert_demo_FooBarList_To_v1beta1_FooBarList(in, out, s)
}

func autoConvert_v1beta1_FooBarOptions_To_demo_FooBarOptions(in *FooBarOptions, out *demo.FooBarOptions, s conversion.Scope) error {
	// WARNING: in.ObjectMeta requires manual conversion: does not exist in peer-type
	out.Arg1 = in.Arg1
	return nil
}

func autoConvert_demo_FooBarOptions_To_v1beta1_FooBarOptions(in *demo.FooBarOptions, out *FooBarOptions, s conversion.Scope) error {
	out.Arg1 = in.Arg1
	return nil
}

// Convert_demo_FooBarOptions_To_v1beta1_FooBarOptions is an autogenerated conversion function.
func Convert_demo_FooBarOptions_To_v1beta1_FooBarOptions(in *demo.FooBarOptions, out *FooBarOptions, s conversion.Scope) error {
	return autoConvert_demo_FooBarOptions_To_v1beta1_FooBarOptions(in, out, s)
}

func autoConvert_v1beta1_FooList_To_demo_FooList(in *FooList, out *demo.FooList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]demo.Foo)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_FooList_To_demo_FooList is an autogenerated conversion function.
func Convert_v1beta1_FooList_To_demo_FooList(in *FooList, out *demo.FooList, s conversion.Scope) error {
	return autoConvert_v1beta1_FooList_To_demo_FooList(in, out, s)
}

func autoConvert_demo_FooList_To_v1beta1_FooList(in *demo.FooList, out *FooList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Foo)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_demo_FooList_To_v1beta1_FooList is an autogenerated conversion function.
func Convert_demo_FooList_To_v1beta1_FooList(in *demo.FooList, out *FooList, s conversion.Scope) error {
	return autoConvert_demo_FooList_To_v1beta1_FooList(in, out, s)
}

func autoConvert_v1beta1_FooSpec_To_demo_FooSpec(in *FooSpec, out *demo.FooSpec, s conversion.Scope) error {
	return nil
}

// Convert_v1beta1_FooSpec_To_demo_FooSpec is an autogenerated conversion function.
func Convert_v1beta1_FooSpec_To_demo_FooSpec(in *FooSpec, out *demo.FooSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_FooSpec_To_demo_FooSpec(in, out, s)
}

func autoConvert_demo_FooSpec_To_v1beta1_FooSpec(in *demo.FooSpec, out *FooSpec, s conversion.Scope) error {
	return nil
}

// Convert_demo_FooSpec_To_v1beta1_FooSpec is an autogenerated conversion function.
func Convert_demo_FooSpec_To_v1beta1_FooSpec(in *demo.FooSpec, out *FooSpec, s conversion.Scope) error {
	return autoConvert_demo_FooSpec_To_v1beta1_FooSpec(in, out, s)
}

func autoConvert_v1beta1_FooStatus_To_demo_FooStatus(in *FooStatus, out *demo.FooStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1beta1_FooStatus_To_demo_FooStatus is an autogenerated conversion function.
func Convert_v1beta1_FooStatus_To_demo_FooStatus(in *FooStatus, out *demo.FooStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_FooStatus_To_demo_FooStatus(in, out, s)
}

func autoConvert_demo_FooStatus_To_v1beta1_FooStatus(in *demo.FooStatus, out *FooStatus, s conversion.Scope) error {
	return nil
}

// Convert_demo_FooStatus_To_v1beta1_FooStatus is an autogenerated conversion function.
func Convert_demo_FooStatus_To_v1beta1_FooStatus(in *demo.FooStatus, out *FooStatus, s conversion.Scope) error {
	return autoConvert_demo_FooStatus_To_v1beta1_FooStatus(in, out, s)
}