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

type GroupSubscriptionAssociationObservation struct {
}

type GroupSubscriptionAssociationParameters struct {

	// +kubebuilder:validation:Required
	ManagementGroupID *string `json:"managementGroupId" tf:"management_group_id,omitempty"`

	// +kubebuilder:validation:Required
	SubscriptionID *string `json:"subscriptionId" tf:"subscription_id,omitempty"`
}

// GroupSubscriptionAssociationSpec defines the desired state of GroupSubscriptionAssociation
type GroupSubscriptionAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupSubscriptionAssociationParameters `json:"forProvider"`
}

// GroupSubscriptionAssociationStatus defines the observed state of GroupSubscriptionAssociation.
type GroupSubscriptionAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupSubscriptionAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GroupSubscriptionAssociation is the Schema for the GroupSubscriptionAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type GroupSubscriptionAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupSubscriptionAssociationSpec   `json:"spec"`
	Status            GroupSubscriptionAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupSubscriptionAssociationList contains a list of GroupSubscriptionAssociations
type GroupSubscriptionAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupSubscriptionAssociation `json:"items"`
}

// Repository type metadata.
var (
	GroupSubscriptionAssociation_Kind             = "GroupSubscriptionAssociation"
	GroupSubscriptionAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GroupSubscriptionAssociation_Kind}.String()
	GroupSubscriptionAssociation_KindAPIVersion   = GroupSubscriptionAssociation_Kind + "." + CRDGroupVersion.String()
	GroupSubscriptionAssociation_GroupVersionKind = CRDGroupVersion.WithKind(GroupSubscriptionAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&GroupSubscriptionAssociation{}, &GroupSubscriptionAssociationList{})
}