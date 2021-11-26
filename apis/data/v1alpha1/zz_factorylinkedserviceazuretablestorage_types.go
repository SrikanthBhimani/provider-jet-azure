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

type FactoryLinkedServiceAzureTableStorageObservation struct {
}

type FactoryLinkedServiceAzureTableStorageParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Required
	ConnectionStringSecretRef v1.SecretKeySelector `json:"connectionStringSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	DataFactoryName *string `json:"dataFactoryName" tf:"data_factory_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

// FactoryLinkedServiceAzureTableStorageSpec defines the desired state of FactoryLinkedServiceAzureTableStorage
type FactoryLinkedServiceAzureTableStorageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FactoryLinkedServiceAzureTableStorageParameters `json:"forProvider"`
}

// FactoryLinkedServiceAzureTableStorageStatus defines the observed state of FactoryLinkedServiceAzureTableStorage.
type FactoryLinkedServiceAzureTableStorageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FactoryLinkedServiceAzureTableStorageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FactoryLinkedServiceAzureTableStorage is the Schema for the FactoryLinkedServiceAzureTableStorages API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type FactoryLinkedServiceAzureTableStorage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FactoryLinkedServiceAzureTableStorageSpec   `json:"spec"`
	Status            FactoryLinkedServiceAzureTableStorageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FactoryLinkedServiceAzureTableStorageList contains a list of FactoryLinkedServiceAzureTableStorages
type FactoryLinkedServiceAzureTableStorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FactoryLinkedServiceAzureTableStorage `json:"items"`
}

// Repository type metadata.
var (
	FactoryLinkedServiceAzureTableStorage_Kind             = "FactoryLinkedServiceAzureTableStorage"
	FactoryLinkedServiceAzureTableStorage_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FactoryLinkedServiceAzureTableStorage_Kind}.String()
	FactoryLinkedServiceAzureTableStorage_KindAPIVersion   = FactoryLinkedServiceAzureTableStorage_Kind + "." + CRDGroupVersion.String()
	FactoryLinkedServiceAzureTableStorage_GroupVersionKind = CRDGroupVersion.WithKind(FactoryLinkedServiceAzureTableStorage_Kind)
)

func init() {
	SchemeBuilder.Register(&FactoryLinkedServiceAzureTableStorage{}, &FactoryLinkedServiceAzureTableStorageList{})
}