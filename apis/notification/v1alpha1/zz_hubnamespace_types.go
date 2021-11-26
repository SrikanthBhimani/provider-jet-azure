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

type HubNamespaceObservation struct {
	ServicebusEndpoint *string `json:"servicebusEndpoint,omitempty" tf:"servicebus_endpoint,omitempty"`
}

type HubNamespaceParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	NamespaceType *string `json:"namespaceType" tf:"namespace_type,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	SkuName *string `json:"skuName" tf:"sku_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// HubNamespaceSpec defines the desired state of HubNamespace
type HubNamespaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HubNamespaceParameters `json:"forProvider"`
}

// HubNamespaceStatus defines the observed state of HubNamespace.
type HubNamespaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HubNamespaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HubNamespace is the Schema for the HubNamespaces API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type HubNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HubNamespaceSpec   `json:"spec"`
	Status            HubNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HubNamespaceList contains a list of HubNamespaces
type HubNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HubNamespace `json:"items"`
}

// Repository type metadata.
var (
	HubNamespace_Kind             = "HubNamespace"
	HubNamespace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HubNamespace_Kind}.String()
	HubNamespace_KindAPIVersion   = HubNamespace_Kind + "." + CRDGroupVersion.String()
	HubNamespace_GroupVersionKind = CRDGroupVersion.WithKind(HubNamespace_Kind)
)

func init() {
	SchemeBuilder.Register(&HubNamespace{}, &HubNamespaceList{})
}