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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this TwinsEndpointEventgrid.
func (mg *TwinsEndpointEventgrid) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this TwinsEndpointEventgrid.
func (mg *TwinsEndpointEventgrid) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this TwinsEndpointEventgrid.
func (mg *TwinsEndpointEventgrid) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this TwinsEndpointEventgrid.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *TwinsEndpointEventgrid) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this TwinsEndpointEventgrid.
func (mg *TwinsEndpointEventgrid) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this TwinsEndpointEventgrid.
func (mg *TwinsEndpointEventgrid) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this TwinsEndpointEventgrid.
func (mg *TwinsEndpointEventgrid) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this TwinsEndpointEventgrid.
func (mg *TwinsEndpointEventgrid) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this TwinsEndpointEventgrid.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *TwinsEndpointEventgrid) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this TwinsEndpointEventgrid.
func (mg *TwinsEndpointEventgrid) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this TwinsEndpointEventhub.
func (mg *TwinsEndpointEventhub) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this TwinsEndpointEventhub.
func (mg *TwinsEndpointEventhub) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this TwinsEndpointEventhub.
func (mg *TwinsEndpointEventhub) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this TwinsEndpointEventhub.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *TwinsEndpointEventhub) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this TwinsEndpointEventhub.
func (mg *TwinsEndpointEventhub) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this TwinsEndpointEventhub.
func (mg *TwinsEndpointEventhub) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this TwinsEndpointEventhub.
func (mg *TwinsEndpointEventhub) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this TwinsEndpointEventhub.
func (mg *TwinsEndpointEventhub) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this TwinsEndpointEventhub.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *TwinsEndpointEventhub) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this TwinsEndpointEventhub.
func (mg *TwinsEndpointEventhub) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this TwinsEndpointServicebus.
func (mg *TwinsEndpointServicebus) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this TwinsEndpointServicebus.
func (mg *TwinsEndpointServicebus) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this TwinsEndpointServicebus.
func (mg *TwinsEndpointServicebus) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this TwinsEndpointServicebus.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *TwinsEndpointServicebus) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this TwinsEndpointServicebus.
func (mg *TwinsEndpointServicebus) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this TwinsEndpointServicebus.
func (mg *TwinsEndpointServicebus) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this TwinsEndpointServicebus.
func (mg *TwinsEndpointServicebus) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this TwinsEndpointServicebus.
func (mg *TwinsEndpointServicebus) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this TwinsEndpointServicebus.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *TwinsEndpointServicebus) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this TwinsEndpointServicebus.
func (mg *TwinsEndpointServicebus) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this TwinsInstance.
func (mg *TwinsInstance) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this TwinsInstance.
func (mg *TwinsInstance) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this TwinsInstance.
func (mg *TwinsInstance) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this TwinsInstance.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *TwinsInstance) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this TwinsInstance.
func (mg *TwinsInstance) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this TwinsInstance.
func (mg *TwinsInstance) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this TwinsInstance.
func (mg *TwinsInstance) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this TwinsInstance.
func (mg *TwinsInstance) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this TwinsInstance.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *TwinsInstance) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this TwinsInstance.
func (mg *TwinsInstance) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}