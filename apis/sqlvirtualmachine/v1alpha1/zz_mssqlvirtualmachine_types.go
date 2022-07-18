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

type AutoBackupObservation struct {
}

type AutoBackupParameters struct {

	// +kubebuilder:validation:Optional
	EncryptionEnabled *bool `json:"encryptionEnabled,omitempty" tf:"encryption_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	EncryptionPasswordSecretRef *v1.SecretKeySelector `json:"encryptionPasswordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ManualSchedule []ManualScheduleParameters `json:"manualSchedule,omitempty" tf:"manual_schedule,omitempty"`

	// +kubebuilder:validation:Required
	RetentionPeriodInDays *float64 `json:"retentionPeriodInDays" tf:"retention_period_in_days,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountAccessKey *string `json:"storageAccountAccessKey" tf:"storage_account_access_key,omitempty"`

	// +kubebuilder:validation:Required
	StorageBlobEndpoint *string `json:"storageBlobEndpoint" tf:"storage_blob_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	SystemDatabasesBackupEnabled *bool `json:"systemDatabasesBackupEnabled,omitempty" tf:"system_databases_backup_enabled,omitempty"`
}

type AutoPatchingObservation struct {
}

type AutoPatchingParameters struct {

	// +kubebuilder:validation:Required
	DayOfWeek *string `json:"dayOfWeek" tf:"day_of_week,omitempty"`

	// +kubebuilder:validation:Required
	MaintenanceWindowDurationInMinutes *float64 `json:"maintenanceWindowDurationInMinutes" tf:"maintenance_window_duration_in_minutes,omitempty"`

	// +kubebuilder:validation:Required
	MaintenanceWindowStartingHour *float64 `json:"maintenanceWindowStartingHour" tf:"maintenance_window_starting_hour,omitempty"`
}

type DataSettingsObservation struct {
}

type DataSettingsParameters struct {

	// +kubebuilder:validation:Required
	DefaultFilePath *string `json:"defaultFilePath" tf:"default_file_path,omitempty"`

	// +kubebuilder:validation:Required
	Luns []*float64 `json:"luns" tf:"luns,omitempty"`
}

type KeyVaultCredentialObservation struct {
}

type KeyVaultCredentialParameters struct {

	// +kubebuilder:validation:Required
	KeyVaultURLSecretRef v1.SecretKeySelector `json:"keyVaultUrlSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ServicePrincipalNameSecretRef v1.SecretKeySelector `json:"servicePrincipalNameSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	ServicePrincipalSecretSecretRef v1.SecretKeySelector `json:"servicePrincipalSecretSecretRef" tf:"-"`
}

type LogSettingsObservation struct {
}

type LogSettingsParameters struct {

	// +kubebuilder:validation:Required
	DefaultFilePath *string `json:"defaultFilePath" tf:"default_file_path,omitempty"`

	// +kubebuilder:validation:Required
	Luns []*float64 `json:"luns" tf:"luns,omitempty"`
}

type MSSQLVirtualMachineObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MSSQLVirtualMachineParameters struct {

	// +kubebuilder:validation:Optional
	AutoBackup []AutoBackupParameters `json:"autoBackup,omitempty" tf:"auto_backup,omitempty"`

	// +kubebuilder:validation:Optional
	AutoPatching []AutoPatchingParameters `json:"autoPatching,omitempty" tf:"auto_patching,omitempty"`

	// +kubebuilder:validation:Optional
	KeyVaultCredential []KeyVaultCredentialParameters `json:"keyVaultCredential,omitempty" tf:"key_vault_credential,omitempty"`

	// +kubebuilder:validation:Optional
	RServicesEnabled *bool `json:"rServicesEnabled,omitempty" tf:"r_services_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	SQLConnectivityPort *float64 `json:"sqlConnectivityPort,omitempty" tf:"sql_connectivity_port,omitempty"`

	// +kubebuilder:validation:Optional
	SQLConnectivityType *string `json:"sqlConnectivityType,omitempty" tf:"sql_connectivity_type,omitempty"`

	// +kubebuilder:validation:Optional
	SQLConnectivityUpdatePasswordSecretRef *v1.SecretKeySelector `json:"sqlConnectivityUpdatePasswordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SQLConnectivityUpdateUsernameSecretRef *v1.SecretKeySelector `json:"sqlConnectivityUpdateUsernameSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	SQLLicenseType *string `json:"sqlLicenseType" tf:"sql_license_type,omitempty"`

	// +kubebuilder:validation:Optional
	StorageConfiguration []StorageConfigurationParameters `json:"storageConfiguration,omitempty" tf:"storage_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	VirtualMachineID *string `json:"virtualMachineId" tf:"virtual_machine_id,omitempty"`
}

type ManualScheduleObservation struct {
}

type ManualScheduleParameters struct {

	// +kubebuilder:validation:Required
	FullBackupFrequency *string `json:"fullBackupFrequency" tf:"full_backup_frequency,omitempty"`

	// +kubebuilder:validation:Required
	FullBackupStartHour *float64 `json:"fullBackupStartHour" tf:"full_backup_start_hour,omitempty"`

	// +kubebuilder:validation:Required
	FullBackupWindowInHours *float64 `json:"fullBackupWindowInHours" tf:"full_backup_window_in_hours,omitempty"`

	// +kubebuilder:validation:Required
	LogBackupFrequencyInMinutes *float64 `json:"logBackupFrequencyInMinutes" tf:"log_backup_frequency_in_minutes,omitempty"`
}

type StorageConfigurationObservation struct {
}

type StorageConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	DataSettings []DataSettingsParameters `json:"dataSettings,omitempty" tf:"data_settings,omitempty"`

	// +kubebuilder:validation:Required
	DiskType *string `json:"diskType" tf:"disk_type,omitempty"`

	// +kubebuilder:validation:Optional
	LogSettings []LogSettingsParameters `json:"logSettings,omitempty" tf:"log_settings,omitempty"`

	// +kubebuilder:validation:Required
	StorageWorkloadType *string `json:"storageWorkloadType" tf:"storage_workload_type,omitempty"`

	// +kubebuilder:validation:Optional
	TempDBSettings []TempDBSettingsParameters `json:"tempDbSettings,omitempty" tf:"temp_db_settings,omitempty"`
}

type TempDBSettingsObservation struct {
}

type TempDBSettingsParameters struct {

	// +kubebuilder:validation:Required
	DefaultFilePath *string `json:"defaultFilePath" tf:"default_file_path,omitempty"`

	// +kubebuilder:validation:Required
	Luns []*float64 `json:"luns" tf:"luns,omitempty"`
}

// MSSQLVirtualMachineSpec defines the desired state of MSSQLVirtualMachine
type MSSQLVirtualMachineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MSSQLVirtualMachineParameters `json:"forProvider"`
}

// MSSQLVirtualMachineStatus defines the observed state of MSSQLVirtualMachine.
type MSSQLVirtualMachineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MSSQLVirtualMachineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLVirtualMachine is the Schema for the MSSQLVirtualMachines API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type MSSQLVirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MSSQLVirtualMachineSpec   `json:"spec"`
	Status            MSSQLVirtualMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLVirtualMachineList contains a list of MSSQLVirtualMachines
type MSSQLVirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSSQLVirtualMachine `json:"items"`
}

// Repository type metadata.
var (
	MSSQLVirtualMachine_Kind             = "MSSQLVirtualMachine"
	MSSQLVirtualMachine_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MSSQLVirtualMachine_Kind}.String()
	MSSQLVirtualMachine_KindAPIVersion   = MSSQLVirtualMachine_Kind + "." + CRDGroupVersion.String()
	MSSQLVirtualMachine_GroupVersionKind = CRDGroupVersion.WithKind(MSSQLVirtualMachine_Kind)
)

func init() {
	SchemeBuilder.Register(&MSSQLVirtualMachine{}, &MSSQLVirtualMachineList{})
}