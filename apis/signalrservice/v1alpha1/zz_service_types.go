/*
Copyright 2022 The Crossplane Authors.

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

type CorsObservation struct {
}

type CorsParameters struct {

	// +kubebuilder:validation:Required
	AllowedOrigins []*string `json:"allowedOrigins" tf:"allowed_origins,omitempty"`
}

type FeaturesObservation struct {
}

type FeaturesParameters struct {

	// +kubebuilder:validation:Required
	Flag *string `json:"flag" tf:"flag,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type ServiceObservation struct {
	HostName *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	PublicPort *float64 `json:"publicPort,omitempty" tf:"public_port,omitempty"`

	ServerPort *float64 `json:"serverPort,omitempty" tf:"server_port,omitempty"`
}

type ServiceParameters struct {

	// +kubebuilder:validation:Optional
	Cors []CorsParameters `json:"cors,omitempty" tf:"cors,omitempty"`

	// +kubebuilder:validation:Optional
	Features []FeaturesParameters `json:"features,omitempty" tf:"features,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Sku []SkuParameters `json:"sku" tf:"sku,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	UpstreamEndpoint []UpstreamEndpointParameters `json:"upstreamEndpoint,omitempty" tf:"upstream_endpoint,omitempty"`
}

type SkuObservation struct {
}

type SkuParameters struct {

	// +kubebuilder:validation:Required
	Capacity *float64 `json:"capacity" tf:"capacity,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type UpstreamEndpointObservation struct {
}

type UpstreamEndpointParameters struct {

	// +kubebuilder:validation:Required
	CategoryPattern []*string `json:"categoryPattern" tf:"category_pattern,omitempty"`

	// +kubebuilder:validation:Required
	EventPattern []*string `json:"eventPattern" tf:"event_pattern,omitempty"`

	// +kubebuilder:validation:Required
	HubPattern []*string `json:"hubPattern" tf:"hub_pattern,omitempty"`

	// +kubebuilder:validation:Required
	URLTemplate *string `json:"urlTemplate" tf:"url_template,omitempty"`
}

// ServiceSpec defines the desired state of Service
type ServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceParameters `json:"forProvider"`
}

// ServiceStatus defines the observed state of Service.
type ServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Service is the Schema for the Services API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type Service struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceSpec   `json:"spec"`
	Status            ServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceList contains a list of Services
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Service `json:"items"`
}

// Repository type metadata.
var (
	Service_Kind             = "Service"
	Service_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Service_Kind}.String()
	Service_KindAPIVersion   = Service_Kind + "." + CRDGroupVersion.String()
	Service_GroupVersionKind = CRDGroupVersion.WithKind(Service_Kind)
)

func init() {
	SchemeBuilder.Register(&Service{}, &ServiceList{})
}