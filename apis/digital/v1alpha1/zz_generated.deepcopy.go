// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinsEndpointEventgrid) DeepCopyInto(out *TwinsEndpointEventgrid) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinsEndpointEventgrid.
func (in *TwinsEndpointEventgrid) DeepCopy() *TwinsEndpointEventgrid {
	if in == nil {
		return nil
	}
	out := new(TwinsEndpointEventgrid)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TwinsEndpointEventgrid) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinsEndpointEventgridList) DeepCopyInto(out *TwinsEndpointEventgridList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TwinsEndpointEventgrid, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinsEndpointEventgridList.
func (in *TwinsEndpointEventgridList) DeepCopy() *TwinsEndpointEventgridList {
	if in == nil {
		return nil
	}
	out := new(TwinsEndpointEventgridList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TwinsEndpointEventgridList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinsEndpointEventgridObservation) DeepCopyInto(out *TwinsEndpointEventgridObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinsEndpointEventgridObservation.
func (in *TwinsEndpointEventgridObservation) DeepCopy() *TwinsEndpointEventgridObservation {
	if in == nil {
		return nil
	}
	out := new(TwinsEndpointEventgridObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinsEndpointEventgridParameters) DeepCopyInto(out *TwinsEndpointEventgridParameters) {
	*out = *in
	if in.DeadLetterStorageSecret != nil {
		in, out := &in.DeadLetterStorageSecret, &out.DeadLetterStorageSecret
		*out = new(string)
		**out = **in
	}
	if in.DigitalTwinsID != nil {
		in, out := &in.DigitalTwinsID, &out.DigitalTwinsID
		*out = new(string)
		**out = **in
	}
	if in.EventgridTopicEndpoint != nil {
		in, out := &in.EventgridTopicEndpoint, &out.EventgridTopicEndpoint
		*out = new(string)
		**out = **in
	}
	if in.EventgridTopicPrimaryAccessKey != nil {
		in, out := &in.EventgridTopicPrimaryAccessKey, &out.EventgridTopicPrimaryAccessKey
		*out = new(string)
		**out = **in
	}
	if in.EventgridTopicSecondaryAccessKey != nil {
		in, out := &in.EventgridTopicSecondaryAccessKey, &out.EventgridTopicSecondaryAccessKey
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinsEndpointEventgridParameters.
func (in *TwinsEndpointEventgridParameters) DeepCopy() *TwinsEndpointEventgridParameters {
	if in == nil {
		return nil
	}
	out := new(TwinsEndpointEventgridParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinsEndpointEventgridSpec) DeepCopyInto(out *TwinsEndpointEventgridSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinsEndpointEventgridSpec.
func (in *TwinsEndpointEventgridSpec) DeepCopy() *TwinsEndpointEventgridSpec {
	if in == nil {
		return nil
	}
	out := new(TwinsEndpointEventgridSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinsEndpointEventgridStatus) DeepCopyInto(out *TwinsEndpointEventgridStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinsEndpointEventgridStatus.
func (in *TwinsEndpointEventgridStatus) DeepCopy() *TwinsEndpointEventgridStatus {
	if in == nil {
		return nil
	}
	out := new(TwinsEndpointEventgridStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinsEndpointEventhub) DeepCopyInto(out *TwinsEndpointEventhub) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinsEndpointEventhub.
func (in *TwinsEndpointEventhub) DeepCopy() *TwinsEndpointEventhub {
	if in == nil {
		return nil
	}
	out := new(TwinsEndpointEventhub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TwinsEndpointEventhub) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinsEndpointEventhubList) DeepCopyInto(out *TwinsEndpointEventhubList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TwinsEndpointEventhub, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinsEndpointEventhubList.
func (in *TwinsEndpointEventhubList) DeepCopy() *TwinsEndpointEventhubList {
	if in == nil {
		return nil
	}
	out := new(TwinsEndpointEventhubList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TwinsEndpointEventhubList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinsEndpointEventhubObservation) DeepCopyInto(out *TwinsEndpointEventhubObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinsEndpointEventhubObservation.
func (in *TwinsEndpointEventhubObservation) DeepCopy() *TwinsEndpointEventhubObservation {
	if in == nil {
		return nil
	}
	out := new(TwinsEndpointEventhubObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinsEndpointEventhubParameters) DeepCopyInto(out *TwinsEndpointEventhubParameters) {
	*out = *in
	if in.DeadLetterStorageSecretSecretRef != nil {
		in, out := &in.DeadLetterStorageSecretSecretRef, &out.DeadLetterStorageSecretSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.DigitalTwinsID != nil {
		in, out := &in.DigitalTwinsID, &out.DigitalTwinsID
		*out = new(string)
		**out = **in
	}
	out.EventhubPrimaryConnectionStringSecretRef = in.EventhubPrimaryConnectionStringSecretRef
	out.EventhubSecondaryConnectionStringSecretRef = in.EventhubSecondaryConnectionStringSecretRef
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinsEndpointEventhubParameters.
func (in *TwinsEndpointEventhubParameters) DeepCopy() *TwinsEndpointEventhubParameters {
	if in == nil {
		return nil
	}
	out := new(TwinsEndpointEventhubParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinsEndpointEventhubSpec) DeepCopyInto(out *TwinsEndpointEventhubSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinsEndpointEventhubSpec.
func (in *TwinsEndpointEventhubSpec) DeepCopy() *TwinsEndpointEventhubSpec {
	if in == nil {
		return nil
	}
	out := new(TwinsEndpointEventhubSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinsEndpointEventhubStatus) DeepCopyInto(out *TwinsEndpointEventhubStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinsEndpointEventhubStatus.
func (in *TwinsEndpointEventhubStatus) DeepCopy() *TwinsEndpointEventhubStatus {
	if in == nil {
		return nil
	}
	out := new(TwinsEndpointEventhubStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinsEndpointServicebus) DeepCopyInto(out *TwinsEndpointServicebus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinsEndpointServicebus.
func (in *TwinsEndpointServicebus) DeepCopy() *TwinsEndpointServicebus {
	if in == nil {
		return nil
	}
	out := new(TwinsEndpointServicebus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TwinsEndpointServicebus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinsEndpointServicebusList) DeepCopyInto(out *TwinsEndpointServicebusList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TwinsEndpointServicebus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinsEndpointServicebusList.
func (in *TwinsEndpointServicebusList) DeepCopy() *TwinsEndpointServicebusList {
	if in == nil {
		return nil
	}
	out := new(TwinsEndpointServicebusList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TwinsEndpointServicebusList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinsEndpointServicebusObservation) DeepCopyInto(out *TwinsEndpointServicebusObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinsEndpointServicebusObservation.
func (in *TwinsEndpointServicebusObservation) DeepCopy() *TwinsEndpointServicebusObservation {
	if in == nil {
		return nil
	}
	out := new(TwinsEndpointServicebusObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinsEndpointServicebusParameters) DeepCopyInto(out *TwinsEndpointServicebusParameters) {
	*out = *in
	if in.DeadLetterStorageSecretSecretRef != nil {
		in, out := &in.DeadLetterStorageSecretSecretRef, &out.DeadLetterStorageSecretSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.DigitalTwinsID != nil {
		in, out := &in.DigitalTwinsID, &out.DigitalTwinsID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	out.ServicebusPrimaryConnectionStringSecretRef = in.ServicebusPrimaryConnectionStringSecretRef
	out.ServicebusSecondaryConnectionStringSecretRef = in.ServicebusSecondaryConnectionStringSecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinsEndpointServicebusParameters.
func (in *TwinsEndpointServicebusParameters) DeepCopy() *TwinsEndpointServicebusParameters {
	if in == nil {
		return nil
	}
	out := new(TwinsEndpointServicebusParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinsEndpointServicebusSpec) DeepCopyInto(out *TwinsEndpointServicebusSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinsEndpointServicebusSpec.
func (in *TwinsEndpointServicebusSpec) DeepCopy() *TwinsEndpointServicebusSpec {
	if in == nil {
		return nil
	}
	out := new(TwinsEndpointServicebusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinsEndpointServicebusStatus) DeepCopyInto(out *TwinsEndpointServicebusStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinsEndpointServicebusStatus.
func (in *TwinsEndpointServicebusStatus) DeepCopy() *TwinsEndpointServicebusStatus {
	if in == nil {
		return nil
	}
	out := new(TwinsEndpointServicebusStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinsInstance) DeepCopyInto(out *TwinsInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinsInstance.
func (in *TwinsInstance) DeepCopy() *TwinsInstance {
	if in == nil {
		return nil
	}
	out := new(TwinsInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TwinsInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinsInstanceList) DeepCopyInto(out *TwinsInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TwinsInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinsInstanceList.
func (in *TwinsInstanceList) DeepCopy() *TwinsInstanceList {
	if in == nil {
		return nil
	}
	out := new(TwinsInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TwinsInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinsInstanceObservation) DeepCopyInto(out *TwinsInstanceObservation) {
	*out = *in
	if in.HostName != nil {
		in, out := &in.HostName, &out.HostName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinsInstanceObservation.
func (in *TwinsInstanceObservation) DeepCopy() *TwinsInstanceObservation {
	if in == nil {
		return nil
	}
	out := new(TwinsInstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinsInstanceParameters) DeepCopyInto(out *TwinsInstanceParameters) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinsInstanceParameters.
func (in *TwinsInstanceParameters) DeepCopy() *TwinsInstanceParameters {
	if in == nil {
		return nil
	}
	out := new(TwinsInstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinsInstanceSpec) DeepCopyInto(out *TwinsInstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinsInstanceSpec.
func (in *TwinsInstanceSpec) DeepCopy() *TwinsInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(TwinsInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TwinsInstanceStatus) DeepCopyInto(out *TwinsInstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TwinsInstanceStatus.
func (in *TwinsInstanceStatus) DeepCopy() *TwinsInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(TwinsInstanceStatus)
	in.DeepCopyInto(out)
	return out
}