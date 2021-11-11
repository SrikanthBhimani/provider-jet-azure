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

type CompositeIndexIndexObservation struct {
}

type CompositeIndexIndexParameters struct {

	// +kubebuilder:validation:Required
	Order *string `json:"order" tf:"order,omitempty"`

	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`
}

type ExcludedPathObservation struct {
}

type ExcludedPathParameters struct {

	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`
}

type IncludedPathObservation struct {
}

type IncludedPathParameters struct {

	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`
}

type IndexingPolicyCompositeIndexObservation struct {
}

type IndexingPolicyCompositeIndexParameters struct {

	// +kubebuilder:validation:Required
	Index []CompositeIndexIndexParameters `json:"index" tf:"index,omitempty"`
}

type IndexingPolicyObservation struct {
}

type IndexingPolicyParameters struct {

	// +kubebuilder:validation:Optional
	CompositeIndex []IndexingPolicyCompositeIndexParameters `json:"compositeIndex,omitempty" tf:"composite_index,omitempty"`

	// +kubebuilder:validation:Optional
	ExcludedPath []ExcludedPathParameters `json:"excludedPath,omitempty" tf:"excluded_path,omitempty"`

	// +kubebuilder:validation:Optional
	IncludedPath []IncludedPathParameters `json:"includedPath,omitempty" tf:"included_path,omitempty"`

	// +kubebuilder:validation:Optional
	IndexingMode *string `json:"indexingMode,omitempty" tf:"indexing_mode,omitempty"`

	// +kubebuilder:validation:Optional
	SpatialIndex []IndexingPolicySpatialIndexParameters `json:"spatialIndex,omitempty" tf:"spatial_index,omitempty"`
}

type IndexingPolicySpatialIndexObservation struct {
	Types []*string `json:"types,omitempty" tf:"types,omitempty"`
}

type IndexingPolicySpatialIndexParameters struct {

	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`
}

type SqlContainerAutoscaleSettingsObservation struct {
}

type SqlContainerAutoscaleSettingsParameters struct {

	// +kubebuilder:validation:Optional
	MaxThroughput *int64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`
}

type SqlContainerConflictResolutionPolicyObservation struct {
}

type SqlContainerConflictResolutionPolicyParameters struct {

	// +kubebuilder:validation:Optional
	ConflictResolutionPath *string `json:"conflictResolutionPath,omitempty" tf:"conflict_resolution_path,omitempty"`

	// +kubebuilder:validation:Optional
	ConflictResolutionProcedure *string `json:"conflictResolutionProcedure,omitempty" tf:"conflict_resolution_procedure,omitempty"`

	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`
}

type SqlContainerObservation struct {
}

type SqlContainerParameters struct {

	// +crossplane:generate:reference:type=Account
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-azure/apis/rconfig.ExtractResourceName()
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AnalyticalStorageTTL *int64 `json:"analyticalStorageTtl,omitempty" tf:"analytical_storage_ttl,omitempty"`

	// +kubebuilder:validation:Optional
	AutoscaleSettings []SqlContainerAutoscaleSettingsParameters `json:"autoscaleSettings,omitempty" tf:"autoscale_settings,omitempty"`

	// +kubebuilder:validation:Optional
	ConflictResolutionPolicy []SqlContainerConflictResolutionPolicyParameters `json:"conflictResolutionPolicy,omitempty" tf:"conflict_resolution_policy,omitempty"`

	// +crossplane:generate:reference:type=SqlDatabase
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-azure/apis/rconfig.ExtractResourceName()
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Optional
	DatabaseNameRef *v1.Reference `json:"databaseNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DatabaseNameSelector *v1.Selector `json:"databaseNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DefaultTTL *int64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// +kubebuilder:validation:Optional
	IndexingPolicy []IndexingPolicyParameters `json:"indexingPolicy,omitempty" tf:"indexing_policy,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PartitionKeyPath *string `json:"partitionKeyPath" tf:"partition_key_path,omitempty"`

	// +kubebuilder:validation:Optional
	PartitionKeyVersion *int64 `json:"partitionKeyVersion,omitempty" tf:"partition_key_version,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-azure/apis/resource/v1alpha1.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-azure/apis/rconfig.ExtractResourceName()
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Throughput *int64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// +kubebuilder:validation:Optional
	UniqueKey []SqlContainerUniqueKeyParameters `json:"uniqueKey,omitempty" tf:"unique_key,omitempty"`
}

type SqlContainerUniqueKeyObservation struct {
}

type SqlContainerUniqueKeyParameters struct {

	// +kubebuilder:validation:Required
	Paths []*string `json:"paths" tf:"paths,omitempty"`
}

// SqlContainerSpec defines the desired state of SqlContainer
type SqlContainerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SqlContainerParameters `json:"forProvider"`
}

// SqlContainerStatus defines the observed state of SqlContainer.
type SqlContainerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SqlContainerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SqlContainer is the Schema for the SqlContainers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type SqlContainer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlContainerSpec   `json:"spec"`
	Status            SqlContainerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SqlContainerList contains a list of SqlContainers
type SqlContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlContainer `json:"items"`
}

// Repository type metadata.
var (
	SqlContainerKind             = "SqlContainer"
	SqlContainerGroupKind        = schema.GroupKind{Group: Group, Kind: SqlContainerKind}.String()
	SqlContainerKindAPIVersion   = SqlContainerKind + "." + GroupVersion.String()
	SqlContainerGroupVersionKind = GroupVersion.WithKind(SqlContainerKind)
)

func init() {
	SchemeBuilder.Register(&SqlContainer{}, &SqlContainerList{})
}
