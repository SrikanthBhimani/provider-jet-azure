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

type TimeSeriesInsightsEventSourceIothubObservation struct {
}

type TimeSeriesInsightsEventSourceIothubParameters struct {

	// +kubebuilder:validation:Required
	ConsumerGroupName *string `json:"consumerGroupName" tf:"consumer_group_name,omitempty"`

	// +kubebuilder:validation:Required
	EnvironmentID *string `json:"environmentId" tf:"environment_id,omitempty"`

	// +kubebuilder:validation:Required
	EventSourceResourceID *string `json:"eventSourceResourceId" tf:"event_source_resource_id,omitempty"`

	// +kubebuilder:validation:Required
	IothubName *string `json:"iothubName" tf:"iothub_name,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	SharedAccessKey *string `json:"sharedAccessKey" tf:"shared_access_key,omitempty"`

	// +kubebuilder:validation:Required
	SharedAccessKeyName *string `json:"sharedAccessKeyName" tf:"shared_access_key_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TimestampPropertyName *string `json:"timestampPropertyName,omitempty" tf:"timestamp_property_name,omitempty"`
}

// TimeSeriesInsightsEventSourceIothubSpec defines the desired state of TimeSeriesInsightsEventSourceIothub
type TimeSeriesInsightsEventSourceIothubSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TimeSeriesInsightsEventSourceIothubParameters `json:"forProvider"`
}

// TimeSeriesInsightsEventSourceIothubStatus defines the observed state of TimeSeriesInsightsEventSourceIothub.
type TimeSeriesInsightsEventSourceIothubStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TimeSeriesInsightsEventSourceIothubObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TimeSeriesInsightsEventSourceIothub is the Schema for the TimeSeriesInsightsEventSourceIothubs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type TimeSeriesInsightsEventSourceIothub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TimeSeriesInsightsEventSourceIothubSpec   `json:"spec"`
	Status            TimeSeriesInsightsEventSourceIothubStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TimeSeriesInsightsEventSourceIothubList contains a list of TimeSeriesInsightsEventSourceIothubs
type TimeSeriesInsightsEventSourceIothubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TimeSeriesInsightsEventSourceIothub `json:"items"`
}

// Repository type metadata.
var (
	TimeSeriesInsightsEventSourceIothub_Kind             = "TimeSeriesInsightsEventSourceIothub"
	TimeSeriesInsightsEventSourceIothub_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TimeSeriesInsightsEventSourceIothub_Kind}.String()
	TimeSeriesInsightsEventSourceIothub_KindAPIVersion   = TimeSeriesInsightsEventSourceIothub_Kind + "." + CRDGroupVersion.String()
	TimeSeriesInsightsEventSourceIothub_GroupVersionKind = CRDGroupVersion.WithKind(TimeSeriesInsightsEventSourceIothub_Kind)
)

func init() {
	SchemeBuilder.Register(&TimeSeriesInsightsEventSourceIothub{}, &TimeSeriesInsightsEventSourceIothubList{})
}