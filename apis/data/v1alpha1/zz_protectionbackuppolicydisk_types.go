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

type CriteriaObservation struct {
}

type CriteriaParameters struct {

	// +kubebuilder:validation:Optional
	AbsoluteCriteria *string `json:"absoluteCriteria,omitempty" tf:"absolute_criteria,omitempty"`
}

type ProtectionBackupPolicyDiskObservation struct {
}

type ProtectionBackupPolicyDiskParameters struct {

	// +kubebuilder:validation:Required
	BackupRepeatingTimeIntervals []*string `json:"backupRepeatingTimeIntervals" tf:"backup_repeating_time_intervals,omitempty"`

	// +kubebuilder:validation:Required
	DefaultRetentionDuration *string `json:"defaultRetentionDuration" tf:"default_retention_duration,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionRule []RetentionRuleParameters `json:"retentionRule,omitempty" tf:"retention_rule,omitempty"`

	// +kubebuilder:validation:Required
	VaultID *string `json:"vaultId" tf:"vault_id,omitempty"`
}

type RetentionRuleObservation struct {
}

type RetentionRuleParameters struct {

	// +kubebuilder:validation:Required
	Criteria []CriteriaParameters `json:"criteria" tf:"criteria,omitempty"`

	// +kubebuilder:validation:Required
	Duration *string `json:"duration" tf:"duration,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Priority *int64 `json:"priority" tf:"priority,omitempty"`
}

// ProtectionBackupPolicyDiskSpec defines the desired state of ProtectionBackupPolicyDisk
type ProtectionBackupPolicyDiskSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProtectionBackupPolicyDiskParameters `json:"forProvider"`
}

// ProtectionBackupPolicyDiskStatus defines the observed state of ProtectionBackupPolicyDisk.
type ProtectionBackupPolicyDiskStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProtectionBackupPolicyDiskObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProtectionBackupPolicyDisk is the Schema for the ProtectionBackupPolicyDisks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ProtectionBackupPolicyDisk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProtectionBackupPolicyDiskSpec   `json:"spec"`
	Status            ProtectionBackupPolicyDiskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProtectionBackupPolicyDiskList contains a list of ProtectionBackupPolicyDisks
type ProtectionBackupPolicyDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProtectionBackupPolicyDisk `json:"items"`
}

// Repository type metadata.
var (
	ProtectionBackupPolicyDisk_Kind             = "ProtectionBackupPolicyDisk"
	ProtectionBackupPolicyDisk_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProtectionBackupPolicyDisk_Kind}.String()
	ProtectionBackupPolicyDisk_KindAPIVersion   = ProtectionBackupPolicyDisk_Kind + "." + CRDGroupVersion.String()
	ProtectionBackupPolicyDisk_GroupVersionKind = CRDGroupVersion.WithKind(ProtectionBackupPolicyDisk_Kind)
)

func init() {
	SchemeBuilder.Register(&ProtectionBackupPolicyDisk{}, &ProtectionBackupPolicyDiskList{})
}