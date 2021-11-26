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

type ManagedDatabaseObservation struct {
}

type ManagedDatabaseParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	SQLManagedInstanceID *string `json:"sqlManagedInstanceId" tf:"sql_managed_instance_id,omitempty"`
}

// ManagedDatabaseSpec defines the desired state of ManagedDatabase
type ManagedDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedDatabaseParameters `json:"forProvider"`
}

// ManagedDatabaseStatus defines the observed state of ManagedDatabase.
type ManagedDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedDatabase is the Schema for the ManagedDatabases API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ManagedDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedDatabaseSpec   `json:"spec"`
	Status            ManagedDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedDatabaseList contains a list of ManagedDatabases
type ManagedDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedDatabase `json:"items"`
}

// Repository type metadata.
var (
	ManagedDatabase_Kind             = "ManagedDatabase"
	ManagedDatabase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagedDatabase_Kind}.String()
	ManagedDatabase_KindAPIVersion   = ManagedDatabase_Kind + "." + CRDGroupVersion.String()
	ManagedDatabase_GroupVersionKind = CRDGroupVersion.WithKind(ManagedDatabase_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagedDatabase{}, &ManagedDatabaseList{})
}