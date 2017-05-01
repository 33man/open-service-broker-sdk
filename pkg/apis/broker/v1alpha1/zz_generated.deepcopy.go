// +build !ignore_autogenerated

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceInstance, InType: reflect.TypeOf(&ServiceInstance{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceInstanceCondition, InType: reflect.TypeOf(&ServiceInstanceCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceInstanceList, InType: reflect.TypeOf(&ServiceInstanceList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceInstanceSpec, InType: reflect.TypeOf(&ServiceInstanceSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceInstanceStatus, InType: reflect.TypeOf(&ServiceInstanceStatus{})},
	)
}

// DeepCopy_v1alpha1_ServiceInstance is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ServiceInstance(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceInstance)
		out := out.(*ServiceInstance)
		*out = *in
		newVal, err := c.DeepCopy(&in.ObjectMeta)
		if err != nil {
			return err
		}
		out.ObjectMeta = *newVal.(*v1.ObjectMeta)

		if err := DeepCopy_v1alpha1_ServiceInstanceStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ServiceInstanceCondition is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ServiceInstanceCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceInstanceCondition)
		out := out.(*ServiceInstanceCondition)
		*out = *in
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		return nil
	}
}

// DeepCopy_v1alpha1_ServiceInstanceList is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ServiceInstanceList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceInstanceList)
		out := out.(*ServiceInstanceList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ServiceInstance, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_ServiceInstance(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1alpha1_ServiceInstanceSpec is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ServiceInstanceSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceInstanceSpec)
		out := out.(*ServiceInstanceSpec)
		*out = *in
		return nil
	}
}

// DeepCopy_v1alpha1_ServiceInstanceStatus is an autogenerated deepcopy function.
func DeepCopy_v1alpha1_ServiceInstanceStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceInstanceStatus)
		out := out.(*ServiceInstanceStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]ServiceInstanceCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_ServiceInstanceCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}
