//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The Kubernetes Authors.

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

package v1alpha1

import (
	unsafe "unsafe"

	settings "github.com/kubernetes-sigs/service-catalog/pkg/apis/settings"
	v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*PodPreset)(nil), (*settings.PodPreset)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PodPreset_To_settings_PodPreset(a.(*PodPreset), b.(*settings.PodPreset), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*settings.PodPreset)(nil), (*PodPreset)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_settings_PodPreset_To_v1alpha1_PodPreset(a.(*settings.PodPreset), b.(*PodPreset), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PodPresetList)(nil), (*settings.PodPresetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PodPresetList_To_settings_PodPresetList(a.(*PodPresetList), b.(*settings.PodPresetList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*settings.PodPresetList)(nil), (*PodPresetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_settings_PodPresetList_To_v1alpha1_PodPresetList(a.(*settings.PodPresetList), b.(*PodPresetList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PodPresetSpec)(nil), (*settings.PodPresetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PodPresetSpec_To_settings_PodPresetSpec(a.(*PodPresetSpec), b.(*settings.PodPresetSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*settings.PodPresetSpec)(nil), (*PodPresetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_settings_PodPresetSpec_To_v1alpha1_PodPresetSpec(a.(*settings.PodPresetSpec), b.(*PodPresetSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_PodPreset_To_settings_PodPreset(in *PodPreset, out *settings.PodPreset, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_PodPresetSpec_To_settings_PodPresetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_PodPreset_To_settings_PodPreset is an autogenerated conversion function.
func Convert_v1alpha1_PodPreset_To_settings_PodPreset(in *PodPreset, out *settings.PodPreset, s conversion.Scope) error {
	return autoConvert_v1alpha1_PodPreset_To_settings_PodPreset(in, out, s)
}

func autoConvert_settings_PodPreset_To_v1alpha1_PodPreset(in *settings.PodPreset, out *PodPreset, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_settings_PodPresetSpec_To_v1alpha1_PodPresetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_settings_PodPreset_To_v1alpha1_PodPreset is an autogenerated conversion function.
func Convert_settings_PodPreset_To_v1alpha1_PodPreset(in *settings.PodPreset, out *PodPreset, s conversion.Scope) error {
	return autoConvert_settings_PodPreset_To_v1alpha1_PodPreset(in, out, s)
}

func autoConvert_v1alpha1_PodPresetList_To_settings_PodPresetList(in *PodPresetList, out *settings.PodPresetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]settings.PodPreset)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_PodPresetList_To_settings_PodPresetList is an autogenerated conversion function.
func Convert_v1alpha1_PodPresetList_To_settings_PodPresetList(in *PodPresetList, out *settings.PodPresetList, s conversion.Scope) error {
	return autoConvert_v1alpha1_PodPresetList_To_settings_PodPresetList(in, out, s)
}

func autoConvert_settings_PodPresetList_To_v1alpha1_PodPresetList(in *settings.PodPresetList, out *PodPresetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]PodPreset)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_settings_PodPresetList_To_v1alpha1_PodPresetList is an autogenerated conversion function.
func Convert_settings_PodPresetList_To_v1alpha1_PodPresetList(in *settings.PodPresetList, out *PodPresetList, s conversion.Scope) error {
	return autoConvert_settings_PodPresetList_To_v1alpha1_PodPresetList(in, out, s)
}

func autoConvert_v1alpha1_PodPresetSpec_To_settings_PodPresetSpec(in *PodPresetSpec, out *settings.PodPresetSpec, s conversion.Scope) error {
	out.Selector = in.Selector
	out.Env = *(*[]v1.EnvVar)(unsafe.Pointer(&in.Env))
	out.EnvFrom = *(*[]v1.EnvFromSource)(unsafe.Pointer(&in.EnvFrom))
	out.Volumes = *(*[]v1.Volume)(unsafe.Pointer(&in.Volumes))
	out.VolumeMounts = *(*[]v1.VolumeMount)(unsafe.Pointer(&in.VolumeMounts))
	return nil
}

// Convert_v1alpha1_PodPresetSpec_To_settings_PodPresetSpec is an autogenerated conversion function.
func Convert_v1alpha1_PodPresetSpec_To_settings_PodPresetSpec(in *PodPresetSpec, out *settings.PodPresetSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_PodPresetSpec_To_settings_PodPresetSpec(in, out, s)
}

func autoConvert_settings_PodPresetSpec_To_v1alpha1_PodPresetSpec(in *settings.PodPresetSpec, out *PodPresetSpec, s conversion.Scope) error {
	out.Selector = in.Selector
	out.Env = *(*[]v1.EnvVar)(unsafe.Pointer(&in.Env))
	out.EnvFrom = *(*[]v1.EnvFromSource)(unsafe.Pointer(&in.EnvFrom))
	out.Volumes = *(*[]v1.Volume)(unsafe.Pointer(&in.Volumes))
	out.VolumeMounts = *(*[]v1.VolumeMount)(unsafe.Pointer(&in.VolumeMounts))
	return nil
}

// Convert_settings_PodPresetSpec_To_v1alpha1_PodPresetSpec is an autogenerated conversion function.
func Convert_settings_PodPresetSpec_To_v1alpha1_PodPresetSpec(in *settings.PodPresetSpec, out *PodPresetSpec, s conversion.Scope) error {
	return autoConvert_settings_PodPresetSpec_To_v1alpha1_PodPresetSpec(in, out, s)
}
