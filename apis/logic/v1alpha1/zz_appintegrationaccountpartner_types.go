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

type AppIntegrationAccountPartnerObservation struct {
}

type AppIntegrationAccountPartnerParameters struct {

	// +kubebuilder:validation:Required
	BusinessIdentity []BusinessIdentityParameters `json:"businessIdentity" tf:"business_identity,omitempty"`

	// +kubebuilder:validation:Required
	IntegrationAccountName *string `json:"integrationAccountName" tf:"integration_account_name,omitempty"`

	// +kubebuilder:validation:Optional
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

type BusinessIdentityObservation struct {
}

type BusinessIdentityParameters struct {

	// +kubebuilder:validation:Required
	Qualifier *string `json:"qualifier" tf:"qualifier,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// AppIntegrationAccountPartnerSpec defines the desired state of AppIntegrationAccountPartner
type AppIntegrationAccountPartnerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppIntegrationAccountPartnerParameters `json:"forProvider"`
}

// AppIntegrationAccountPartnerStatus defines the observed state of AppIntegrationAccountPartner.
type AppIntegrationAccountPartnerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppIntegrationAccountPartnerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppIntegrationAccountPartner is the Schema for the AppIntegrationAccountPartners API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type AppIntegrationAccountPartner struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppIntegrationAccountPartnerSpec   `json:"spec"`
	Status            AppIntegrationAccountPartnerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppIntegrationAccountPartnerList contains a list of AppIntegrationAccountPartners
type AppIntegrationAccountPartnerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppIntegrationAccountPartner `json:"items"`
}

// Repository type metadata.
var (
	AppIntegrationAccountPartner_Kind             = "AppIntegrationAccountPartner"
	AppIntegrationAccountPartner_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppIntegrationAccountPartner_Kind}.String()
	AppIntegrationAccountPartner_KindAPIVersion   = AppIntegrationAccountPartner_Kind + "." + CRDGroupVersion.String()
	AppIntegrationAccountPartner_GroupVersionKind = CRDGroupVersion.WithKind(AppIntegrationAccountPartner_Kind)
)

func init() {
	SchemeBuilder.Register(&AppIntegrationAccountPartner{}, &AppIntegrationAccountPartnerList{})
}