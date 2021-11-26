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

type ManagementIdentityProviderAadb2CObservation struct {
}

type ManagementIdentityProviderAadb2CParameters struct {

	// +kubebuilder:validation:Required
	APIManagementName *string `json:"apiManagementName" tf:"api_management_name,omitempty"`

	// +kubebuilder:validation:Required
	AllowedTenant *string `json:"allowedTenant" tf:"allowed_tenant,omitempty"`

	// +kubebuilder:validation:Required
	Authority *string `json:"authority" tf:"authority,omitempty"`

	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	PasswordResetPolicy *string `json:"passwordResetPolicy,omitempty" tf:"password_reset_policy,omitempty"`

	// +kubebuilder:validation:Optional
	ProfileEditingPolicy *string `json:"profileEditingPolicy,omitempty" tf:"profile_editing_policy,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	SigninPolicy *string `json:"signinPolicy" tf:"signin_policy,omitempty"`

	// +kubebuilder:validation:Required
	SigninTenant *string `json:"signinTenant" tf:"signin_tenant,omitempty"`

	// +kubebuilder:validation:Required
	SignupPolicy *string `json:"signupPolicy" tf:"signup_policy,omitempty"`
}

// ManagementIdentityProviderAadb2CSpec defines the desired state of ManagementIdentityProviderAadb2C
type ManagementIdentityProviderAadb2CSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagementIdentityProviderAadb2CParameters `json:"forProvider"`
}

// ManagementIdentityProviderAadb2CStatus defines the observed state of ManagementIdentityProviderAadb2C.
type ManagementIdentityProviderAadb2CStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagementIdentityProviderAadb2CObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementIdentityProviderAadb2C is the Schema for the ManagementIdentityProviderAadb2Cs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ManagementIdentityProviderAadb2C struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagementIdentityProviderAadb2CSpec   `json:"spec"`
	Status            ManagementIdentityProviderAadb2CStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementIdentityProviderAadb2CList contains a list of ManagementIdentityProviderAadb2Cs
type ManagementIdentityProviderAadb2CList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagementIdentityProviderAadb2C `json:"items"`
}

// Repository type metadata.
var (
	ManagementIdentityProviderAadb2C_Kind             = "ManagementIdentityProviderAadb2C"
	ManagementIdentityProviderAadb2C_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagementIdentityProviderAadb2C_Kind}.String()
	ManagementIdentityProviderAadb2C_KindAPIVersion   = ManagementIdentityProviderAadb2C_Kind + "." + CRDGroupVersion.String()
	ManagementIdentityProviderAadb2C_GroupVersionKind = CRDGroupVersion.WithKind(ManagementIdentityProviderAadb2C_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagementIdentityProviderAadb2C{}, &ManagementIdentityProviderAadb2CList{})
}