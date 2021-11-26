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

type LakeAnalyticsFirewallRuleObservation struct {
}

type LakeAnalyticsFirewallRuleParameters struct {

	// +kubebuilder:validation:Required
	AccountName *string `json:"accountName" tf:"account_name,omitempty"`

	// +kubebuilder:validation:Required
	EndIPAddress *string `json:"endIpAddress" tf:"end_ip_address,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	StartIPAddress *string `json:"startIpAddress" tf:"start_ip_address,omitempty"`
}

// LakeAnalyticsFirewallRuleSpec defines the desired state of LakeAnalyticsFirewallRule
type LakeAnalyticsFirewallRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LakeAnalyticsFirewallRuleParameters `json:"forProvider"`
}

// LakeAnalyticsFirewallRuleStatus defines the observed state of LakeAnalyticsFirewallRule.
type LakeAnalyticsFirewallRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LakeAnalyticsFirewallRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LakeAnalyticsFirewallRule is the Schema for the LakeAnalyticsFirewallRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type LakeAnalyticsFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LakeAnalyticsFirewallRuleSpec   `json:"spec"`
	Status            LakeAnalyticsFirewallRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LakeAnalyticsFirewallRuleList contains a list of LakeAnalyticsFirewallRules
type LakeAnalyticsFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LakeAnalyticsFirewallRule `json:"items"`
}

// Repository type metadata.
var (
	LakeAnalyticsFirewallRule_Kind             = "LakeAnalyticsFirewallRule"
	LakeAnalyticsFirewallRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LakeAnalyticsFirewallRule_Kind}.String()
	LakeAnalyticsFirewallRule_KindAPIVersion   = LakeAnalyticsFirewallRule_Kind + "." + CRDGroupVersion.String()
	LakeAnalyticsFirewallRule_GroupVersionKind = CRDGroupVersion.WithKind(LakeAnalyticsFirewallRule_Kind)
)

func init() {
	SchemeBuilder.Register(&LakeAnalyticsFirewallRule{}, &LakeAnalyticsFirewallRuleList{})
}