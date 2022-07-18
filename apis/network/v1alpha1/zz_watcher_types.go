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

type WatcherObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type WatcherParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

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
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// WatcherSpec defines the desired state of Watcher
type WatcherSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WatcherParameters `json:"forProvider"`
}

// WatcherStatus defines the observed state of Watcher.
type WatcherStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WatcherObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Watcher is the Schema for the Watchers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type Watcher struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WatcherSpec   `json:"spec"`
	Status            WatcherStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WatcherList contains a list of Watchers
type WatcherList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Watcher `json:"items"`
}

// Repository type metadata.
var (
	Watcher_Kind             = "Watcher"
	Watcher_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Watcher_Kind}.String()
	Watcher_KindAPIVersion   = Watcher_Kind + "." + CRDGroupVersion.String()
	Watcher_GroupVersionKind = CRDGroupVersion.WithKind(Watcher_Kind)
)

func init() {
	SchemeBuilder.Register(&Watcher{}, &WatcherList{})
}