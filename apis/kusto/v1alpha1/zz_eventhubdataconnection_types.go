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

type EventHubDataConnectionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EventHubDataConnectionParameters struct {

	// +kubebuilder:validation:Required
	ClusterName *string `json:"clusterName" tf:"cluster_name,omitempty"`

	// +kubebuilder:validation:Optional
	Compression *string `json:"compression,omitempty" tf:"compression,omitempty"`

	// +kubebuilder:validation:Required
	ConsumerGroup *string `json:"consumerGroup" tf:"consumer_group,omitempty"`

	// +kubebuilder:validation:Optional
	DataFormat *string `json:"dataFormat,omitempty" tf:"data_format,omitempty"`

	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Required
	EventHubID *string `json:"eventhubId" tf:"eventhub_id,omitempty"`

	// +kubebuilder:validation:Optional
	EventSystemProperties []*string `json:"eventSystemProperties,omitempty" tf:"event_system_properties,omitempty"`

	// +kubebuilder:validation:Optional
	IdentityID *string `json:"identityId,omitempty" tf:"identity_id,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	MappingRuleName *string `json:"mappingRuleName,omitempty" tf:"mapping_rule_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`
}

// EventHubDataConnectionSpec defines the desired state of EventHubDataConnection
type EventHubDataConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EventHubDataConnectionParameters `json:"forProvider"`
}

// EventHubDataConnectionStatus defines the observed state of EventHubDataConnection.
type EventHubDataConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EventHubDataConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EventHubDataConnection is the Schema for the EventHubDataConnections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type EventHubDataConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventHubDataConnectionSpec   `json:"spec"`
	Status            EventHubDataConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventHubDataConnectionList contains a list of EventHubDataConnections
type EventHubDataConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventHubDataConnection `json:"items"`
}

// Repository type metadata.
var (
	EventHubDataConnection_Kind             = "EventHubDataConnection"
	EventHubDataConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EventHubDataConnection_Kind}.String()
	EventHubDataConnection_KindAPIVersion   = EventHubDataConnection_Kind + "." + CRDGroupVersion.String()
	EventHubDataConnection_GroupVersionKind = CRDGroupVersion.WithKind(EventHubDataConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&EventHubDataConnection{}, &EventHubDataConnectionList{})
}