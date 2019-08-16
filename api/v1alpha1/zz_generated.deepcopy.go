// +build !ignore_autogenerated

/*
Copyright 2019 The xridge kubestone contributors.

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

// autogenerated by controller-gen object, do not modify manually

package v1alpha1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Fio) DeepCopyInto(out *Fio) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fio.
func (in *Fio) DeepCopy() *Fio {
	if in == nil {
		return nil
	}
	out := new(Fio)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Fio) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FioList) DeepCopyInto(out *FioList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Fio, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FioList.
func (in *FioList) DeepCopy() *FioList {
	if in == nil {
		return nil
	}
	out := new(FioList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FioList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FioSpec) DeepCopyInto(out *FioSpec) {
	*out = *in
	out.Image = in.Image
	if in.PersistentVolumeClaimName != nil {
		in, out := &in.PersistentVolumeClaimName, &out.PersistentVolumeClaimName
		*out = new(string)
		**out = **in
	}
	if in.PersistentVolumeClaim != nil {
		in, out := &in.PersistentVolumeClaim, &out.PersistentVolumeClaim
		*out = new(PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.BuiltinJobFiles != nil {
		in, out := &in.BuiltinJobFiles, &out.BuiltinJobFiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CustomJobFiles != nil {
		in, out := &in.CustomJobFiles, &out.CustomJobFiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FioSpec.
func (in *FioSpec) DeepCopy() *FioSpec {
	if in == nil {
		return nil
	}
	out := new(FioSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FioStatus) DeepCopyInto(out *FioStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FioStatus.
func (in *FioStatus) DeepCopy() *FioStatus {
	if in == nil {
		return nil
	}
	out := new(FioStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Iperf3) DeepCopyInto(out *Iperf3) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Iperf3.
func (in *Iperf3) DeepCopy() *Iperf3 {
	if in == nil {
		return nil
	}
	out := new(Iperf3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Iperf3) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Iperf3ConfigurationSpec) DeepCopyInto(out *Iperf3ConfigurationSpec) {
	*out = *in
	if in.PodLabels != nil {
		in, out := &in.PodLabels, &out.PodLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.PodScheduling.DeepCopyInto(&out.PodScheduling)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Iperf3ConfigurationSpec.
func (in *Iperf3ConfigurationSpec) DeepCopy() *Iperf3ConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(Iperf3ConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Iperf3List) DeepCopyInto(out *Iperf3List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Iperf3, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Iperf3List.
func (in *Iperf3List) DeepCopy() *Iperf3List {
	if in == nil {
		return nil
	}
	out := new(Iperf3List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Iperf3List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Iperf3Spec) DeepCopyInto(out *Iperf3Spec) {
	*out = *in
	out.Image = in.Image
	in.ServerConfiguration.DeepCopyInto(&out.ServerConfiguration)
	in.ClientConfiguration.DeepCopyInto(&out.ClientConfiguration)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Iperf3Spec.
func (in *Iperf3Spec) DeepCopy() *Iperf3Spec {
	if in == nil {
		return nil
	}
	out := new(Iperf3Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Iperf3Status) DeepCopyInto(out *Iperf3Status) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Iperf3Status.
func (in *Iperf3Status) DeepCopy() *Iperf3Status {
	if in == nil {
		return nil
	}
	out := new(Iperf3Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeClaimSpec) DeepCopyInto(out *PersistentVolumeClaimSpec) {
	*out = *in
	if in.AccessModes != nil {
		in, out := &in.AccessModes, &out.AccessModes
		*out = make([]PersistentVolumeAccessMode, len(*in))
		copy(*out, *in)
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	if in.VolumeMode != nil {
		in, out := &in.VolumeMode, &out.VolumeMode
		*out = new(PersistentVolumeMode)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeClaimSpec.
func (in *PersistentVolumeClaimSpec) DeepCopy() *PersistentVolumeClaimSpec {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeClaimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSchedulingSpec) DeepCopyInto(out *PodSchedulingSpec) {
	*out = *in
	in.Affinity.DeepCopyInto(&out.Affinity)
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSchedulingSpec.
func (in *PodSchedulingSpec) DeepCopy() *PodSchedulingSpec {
	if in == nil {
		return nil
	}
	out := new(PodSchedulingSpec)
	in.DeepCopyInto(out)
	return out
}
