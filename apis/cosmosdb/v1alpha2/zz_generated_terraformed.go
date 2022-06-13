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

package v1alpha2

import (
	"github.com/pkg/errors"

	"github.com/crossplane/terrajet/pkg/resource"
	"github.com/crossplane/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this Account
func (mg *Account) GetTerraformResourceType() string {
	return "azurerm_cosmosdb_account"
}

// GetConnectionDetailsMapping for this Account
func (tr *Account) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"connection_strings[*]": "status.atProvider.connectionStrings[*]", "primary_key": "status.atProvider.primaryKey", "primary_master_key": "status.atProvider.primaryMasterKey", "primary_readonly_key": "status.atProvider.primaryReadonlyKey", "primary_readonly_master_key": "status.atProvider.primaryReadonlyMasterKey", "secondary_key": "status.atProvider.secondaryKey", "secondary_master_key": "status.atProvider.secondaryMasterKey", "secondary_readonly_key": "status.atProvider.secondaryReadonlyKey", "secondary_readonly_master_key": "status.atProvider.secondaryReadonlyMasterKey"}
}

// GetObservation of this Account
func (tr *Account) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Account
func (tr *Account) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Account
func (tr *Account) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Account
func (tr *Account) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Account
func (tr *Account) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Account using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Account) LateInitialize(attrs []byte) (bool, error) {
	params := &AccountParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Account) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this CassandraKeySpace
func (mg *CassandraKeySpace) GetTerraformResourceType() string {
	return "azurerm_cosmosdb_cassandra_keyspace"
}

// GetConnectionDetailsMapping for this CassandraKeySpace
func (tr *CassandraKeySpace) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this CassandraKeySpace
func (tr *CassandraKeySpace) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this CassandraKeySpace
func (tr *CassandraKeySpace) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this CassandraKeySpace
func (tr *CassandraKeySpace) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this CassandraKeySpace
func (tr *CassandraKeySpace) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this CassandraKeySpace
func (tr *CassandraKeySpace) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this CassandraKeySpace using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *CassandraKeySpace) LateInitialize(attrs []byte) (bool, error) {
	params := &CassandraKeySpaceParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *CassandraKeySpace) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this CassandraTable
func (mg *CassandraTable) GetTerraformResourceType() string {
	return "azurerm_cosmosdb_cassandra_table"
}

// GetConnectionDetailsMapping for this CassandraTable
func (tr *CassandraTable) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this CassandraTable
func (tr *CassandraTable) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this CassandraTable
func (tr *CassandraTable) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this CassandraTable
func (tr *CassandraTable) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this CassandraTable
func (tr *CassandraTable) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this CassandraTable
func (tr *CassandraTable) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this CassandraTable using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *CassandraTable) LateInitialize(attrs []byte) (bool, error) {
	params := &CassandraTableParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *CassandraTable) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this GremlinDatabase
func (mg *GremlinDatabase) GetTerraformResourceType() string {
	return "azurerm_cosmosdb_gremlin_database"
}

// GetConnectionDetailsMapping for this GremlinDatabase
func (tr *GremlinDatabase) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this GremlinDatabase
func (tr *GremlinDatabase) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this GremlinDatabase
func (tr *GremlinDatabase) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this GremlinDatabase
func (tr *GremlinDatabase) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this GremlinDatabase
func (tr *GremlinDatabase) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this GremlinDatabase
func (tr *GremlinDatabase) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this GremlinDatabase using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *GremlinDatabase) LateInitialize(attrs []byte) (bool, error) {
	params := &GremlinDatabaseParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *GremlinDatabase) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this GremlinGraph
func (mg *GremlinGraph) GetTerraformResourceType() string {
	return "azurerm_cosmosdb_gremlin_graph"
}

// GetConnectionDetailsMapping for this GremlinGraph
func (tr *GremlinGraph) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this GremlinGraph
func (tr *GremlinGraph) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this GremlinGraph
func (tr *GremlinGraph) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this GremlinGraph
func (tr *GremlinGraph) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this GremlinGraph
func (tr *GremlinGraph) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this GremlinGraph
func (tr *GremlinGraph) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this GremlinGraph using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *GremlinGraph) LateInitialize(attrs []byte) (bool, error) {
	params := &GremlinGraphParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *GremlinGraph) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this MongoCollection
func (mg *MongoCollection) GetTerraformResourceType() string {
	return "azurerm_cosmosdb_mongo_collection"
}

// GetConnectionDetailsMapping for this MongoCollection
func (tr *MongoCollection) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this MongoCollection
func (tr *MongoCollection) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this MongoCollection
func (tr *MongoCollection) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this MongoCollection
func (tr *MongoCollection) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this MongoCollection
func (tr *MongoCollection) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this MongoCollection
func (tr *MongoCollection) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this MongoCollection using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *MongoCollection) LateInitialize(attrs []byte) (bool, error) {
	params := &MongoCollectionParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *MongoCollection) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this MongoDatabase
func (mg *MongoDatabase) GetTerraformResourceType() string {
	return "azurerm_cosmosdb_mongo_database"
}

// GetConnectionDetailsMapping for this MongoDatabase
func (tr *MongoDatabase) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this MongoDatabase
func (tr *MongoDatabase) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this MongoDatabase
func (tr *MongoDatabase) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this MongoDatabase
func (tr *MongoDatabase) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this MongoDatabase
func (tr *MongoDatabase) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this MongoDatabase
func (tr *MongoDatabase) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this MongoDatabase using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *MongoDatabase) LateInitialize(attrs []byte) (bool, error) {
	params := &MongoDatabaseParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *MongoDatabase) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this NotebookWorkspace
func (mg *NotebookWorkspace) GetTerraformResourceType() string {
	return "azurerm_cosmosdb_notebook_workspace"
}

// GetConnectionDetailsMapping for this NotebookWorkspace
func (tr *NotebookWorkspace) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this NotebookWorkspace
func (tr *NotebookWorkspace) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this NotebookWorkspace
func (tr *NotebookWorkspace) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this NotebookWorkspace
func (tr *NotebookWorkspace) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this NotebookWorkspace
func (tr *NotebookWorkspace) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this NotebookWorkspace
func (tr *NotebookWorkspace) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this NotebookWorkspace using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *NotebookWorkspace) LateInitialize(attrs []byte) (bool, error) {
	params := &NotebookWorkspaceParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *NotebookWorkspace) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SQLContainer
func (mg *SQLContainer) GetTerraformResourceType() string {
	return "azurerm_cosmosdb_sql_container"
}

// GetConnectionDetailsMapping for this SQLContainer
func (tr *SQLContainer) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SQLContainer
func (tr *SQLContainer) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SQLContainer
func (tr *SQLContainer) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SQLContainer
func (tr *SQLContainer) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SQLContainer
func (tr *SQLContainer) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SQLContainer
func (tr *SQLContainer) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SQLContainer using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SQLContainer) LateInitialize(attrs []byte) (bool, error) {
	params := &SQLContainerParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SQLContainer) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this SQLDatabase
func (mg *SQLDatabase) GetTerraformResourceType() string {
	return "azurerm_cosmosdb_sql_database"
}

// GetConnectionDetailsMapping for this SQLDatabase
func (tr *SQLDatabase) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SQLDatabase
func (tr *SQLDatabase) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SQLDatabase
func (tr *SQLDatabase) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SQLDatabase
func (tr *SQLDatabase) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SQLDatabase
func (tr *SQLDatabase) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SQLDatabase
func (tr *SQLDatabase) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SQLDatabase using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SQLDatabase) LateInitialize(attrs []byte) (bool, error) {
	params := &SQLDatabaseParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SQLDatabase) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this SQLFunction
func (mg *SQLFunction) GetTerraformResourceType() string {
	return "azurerm_cosmosdb_sql_function"
}

// GetConnectionDetailsMapping for this SQLFunction
func (tr *SQLFunction) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SQLFunction
func (tr *SQLFunction) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SQLFunction
func (tr *SQLFunction) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SQLFunction
func (tr *SQLFunction) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SQLFunction
func (tr *SQLFunction) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SQLFunction
func (tr *SQLFunction) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SQLFunction using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SQLFunction) LateInitialize(attrs []byte) (bool, error) {
	params := &SQLFunctionParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SQLFunction) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SQLStoredProcedure
func (mg *SQLStoredProcedure) GetTerraformResourceType() string {
	return "azurerm_cosmosdb_sql_stored_procedure"
}

// GetConnectionDetailsMapping for this SQLStoredProcedure
func (tr *SQLStoredProcedure) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SQLStoredProcedure
func (tr *SQLStoredProcedure) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SQLStoredProcedure
func (tr *SQLStoredProcedure) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SQLStoredProcedure
func (tr *SQLStoredProcedure) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SQLStoredProcedure
func (tr *SQLStoredProcedure) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SQLStoredProcedure
func (tr *SQLStoredProcedure) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SQLStoredProcedure using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SQLStoredProcedure) LateInitialize(attrs []byte) (bool, error) {
	params := &SQLStoredProcedureParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SQLStoredProcedure) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SQLTrigger
func (mg *SQLTrigger) GetTerraformResourceType() string {
	return "azurerm_cosmosdb_sql_trigger"
}

// GetConnectionDetailsMapping for this SQLTrigger
func (tr *SQLTrigger) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SQLTrigger
func (tr *SQLTrigger) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SQLTrigger
func (tr *SQLTrigger) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SQLTrigger
func (tr *SQLTrigger) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SQLTrigger
func (tr *SQLTrigger) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SQLTrigger
func (tr *SQLTrigger) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SQLTrigger using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SQLTrigger) LateInitialize(attrs []byte) (bool, error) {
	params := &SQLTriggerParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SQLTrigger) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Table
func (mg *Table) GetTerraformResourceType() string {
	return "azurerm_cosmosdb_table"
}

// GetConnectionDetailsMapping for this Table
func (tr *Table) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Table
func (tr *Table) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Table
func (tr *Table) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Table
func (tr *Table) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Table
func (tr *Table) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Table
func (tr *Table) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Table using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Table) LateInitialize(attrs []byte) (bool, error) {
	params := &TableParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Table) GetTerraformSchemaVersion() int {
	return 1
}
