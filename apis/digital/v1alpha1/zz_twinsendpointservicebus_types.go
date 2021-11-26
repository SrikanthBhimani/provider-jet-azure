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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TwinsEndpointServicebusObservation struct {
}

type TwinsEndpointServicebusParameters struct {

	// +kubebuilder:validation:Optional
	DeadLetterStorageSecretSecretRef *v1.SecretKeySelector `json:"deadLetterStorageSecretSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	DigitalTwinsID *string `json:"digitalTwinsId" tf:"digital_twins_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ServicebusPrimaryConnectionStringSecretRef v1.SecretKeySelector `json:"servicebusPrimaryConnectionStringSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	ServicebusSecondaryConnectionStringSecretRef v1.SecretKeySelector `json:"servicebusSecondaryConnectionStringSecretRef" tf:"-"`
}

// TwinsEndpointServicebusSpec defines the desired state of TwinsEndpointServicebus
type TwinsEndpointServicebusSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TwinsEndpointServicebusParameters `json:"forProvider"`
}

// TwinsEndpointServicebusStatus defines the observed state of TwinsEndpointServicebus.
type TwinsEndpointServicebusStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TwinsEndpointServicebusObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TwinsEndpointServicebus is the Schema for the TwinsEndpointServicebuss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type TwinsEndpointServicebus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TwinsEndpointServicebusSpec   `json:"spec"`
	Status            TwinsEndpointServicebusStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TwinsEndpointServicebusList contains a list of TwinsEndpointServicebuss
type TwinsEndpointServicebusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TwinsEndpointServicebus `json:"items"`
}

// Repository type metadata.
var (
	TwinsEndpointServicebus_Kind             = "TwinsEndpointServicebus"
	TwinsEndpointServicebus_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TwinsEndpointServicebus_Kind}.String()
	TwinsEndpointServicebus_KindAPIVersion   = TwinsEndpointServicebus_Kind + "." + CRDGroupVersion.String()
	TwinsEndpointServicebus_GroupVersionKind = CRDGroupVersion.WithKind(TwinsEndpointServicebus_Kind)
)

func init() {
	SchemeBuilder.Register(&TwinsEndpointServicebus{}, &TwinsEndpointServicebusList{})
}