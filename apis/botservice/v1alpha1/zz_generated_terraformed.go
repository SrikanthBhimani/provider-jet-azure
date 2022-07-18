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
	"github.com/pkg/errors"

	"github.com/crossplane/terrajet/pkg/resource"
	"github.com/crossplane/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this BotChannelAlexa
func (mg *BotChannelAlexa) GetTerraformResourceType() string {
	return "azurerm_bot_channel_alexa"
}

// GetConnectionDetailsMapping for this BotChannelAlexa
func (tr *BotChannelAlexa) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this BotChannelAlexa
func (tr *BotChannelAlexa) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BotChannelAlexa
func (tr *BotChannelAlexa) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BotChannelAlexa
func (tr *BotChannelAlexa) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BotChannelAlexa
func (tr *BotChannelAlexa) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BotChannelAlexa
func (tr *BotChannelAlexa) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BotChannelAlexa using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BotChannelAlexa) LateInitialize(attrs []byte) (bool, error) {
	params := &BotChannelAlexaParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BotChannelAlexa) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BotChannelDirectLineSpeech
func (mg *BotChannelDirectLineSpeech) GetTerraformResourceType() string {
	return "azurerm_bot_channel_direct_line_speech"
}

// GetConnectionDetailsMapping for this BotChannelDirectLineSpeech
func (tr *BotChannelDirectLineSpeech) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"cognitive_service_access_key": "spec.forProvider.cognitiveServiceAccessKeySecretRef"}
}

// GetObservation of this BotChannelDirectLineSpeech
func (tr *BotChannelDirectLineSpeech) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BotChannelDirectLineSpeech
func (tr *BotChannelDirectLineSpeech) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BotChannelDirectLineSpeech
func (tr *BotChannelDirectLineSpeech) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BotChannelDirectLineSpeech
func (tr *BotChannelDirectLineSpeech) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BotChannelDirectLineSpeech
func (tr *BotChannelDirectLineSpeech) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BotChannelDirectLineSpeech using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BotChannelDirectLineSpeech) LateInitialize(attrs []byte) (bool, error) {
	params := &BotChannelDirectLineSpeechParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BotChannelDirectLineSpeech) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BotChannelDirectLine
func (mg *BotChannelDirectLine) GetTerraformResourceType() string {
	return "azurerm_bot_channel_directline"
}

// GetConnectionDetailsMapping for this BotChannelDirectLine
func (tr *BotChannelDirectLine) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"site[*].key": "status.atProvider.site[*].key", "site[*].key2": "status.atProvider.site[*].key2"}
}

// GetObservation of this BotChannelDirectLine
func (tr *BotChannelDirectLine) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BotChannelDirectLine
func (tr *BotChannelDirectLine) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BotChannelDirectLine
func (tr *BotChannelDirectLine) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BotChannelDirectLine
func (tr *BotChannelDirectLine) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BotChannelDirectLine
func (tr *BotChannelDirectLine) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BotChannelDirectLine using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BotChannelDirectLine) LateInitialize(attrs []byte) (bool, error) {
	params := &BotChannelDirectLineParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BotChannelDirectLine) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BotChannelEmail
func (mg *BotChannelEmail) GetTerraformResourceType() string {
	return "azurerm_bot_channel_email"
}

// GetConnectionDetailsMapping for this BotChannelEmail
func (tr *BotChannelEmail) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"email_password": "spec.forProvider.emailPasswordSecretRef"}
}

// GetObservation of this BotChannelEmail
func (tr *BotChannelEmail) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BotChannelEmail
func (tr *BotChannelEmail) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BotChannelEmail
func (tr *BotChannelEmail) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BotChannelEmail
func (tr *BotChannelEmail) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BotChannelEmail
func (tr *BotChannelEmail) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BotChannelEmail using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BotChannelEmail) LateInitialize(attrs []byte) (bool, error) {
	params := &BotChannelEmailParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BotChannelEmail) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BotChannelFacebook
func (mg *BotChannelFacebook) GetTerraformResourceType() string {
	return "azurerm_bot_channel_facebook"
}

// GetConnectionDetailsMapping for this BotChannelFacebook
func (tr *BotChannelFacebook) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"facebook_application_secret": "spec.forProvider.facebookApplicationSecretSecretRef", "page[*].access_token": "spec.forProvider.page[*].accessTokenSecretRef"}
}

// GetObservation of this BotChannelFacebook
func (tr *BotChannelFacebook) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BotChannelFacebook
func (tr *BotChannelFacebook) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BotChannelFacebook
func (tr *BotChannelFacebook) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BotChannelFacebook
func (tr *BotChannelFacebook) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BotChannelFacebook
func (tr *BotChannelFacebook) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BotChannelFacebook using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BotChannelFacebook) LateInitialize(attrs []byte) (bool, error) {
	params := &BotChannelFacebookParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BotChannelFacebook) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BotChannelLine
func (mg *BotChannelLine) GetTerraformResourceType() string {
	return "azurerm_bot_channel_line"
}

// GetConnectionDetailsMapping for this BotChannelLine
func (tr *BotChannelLine) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"line_channel[*].access_token": "spec.forProvider.lineChannel[*].accessTokenSecretRef", "line_channel[*].secret": "spec.forProvider.lineChannel[*].secretSecretRef"}
}

// GetObservation of this BotChannelLine
func (tr *BotChannelLine) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BotChannelLine
func (tr *BotChannelLine) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BotChannelLine
func (tr *BotChannelLine) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BotChannelLine
func (tr *BotChannelLine) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BotChannelLine
func (tr *BotChannelLine) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BotChannelLine using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BotChannelLine) LateInitialize(attrs []byte) (bool, error) {
	params := &BotChannelLineParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BotChannelLine) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BotChannelMSTeams
func (mg *BotChannelMSTeams) GetTerraformResourceType() string {
	return "azurerm_bot_channel_ms_teams"
}

// GetConnectionDetailsMapping for this BotChannelMSTeams
func (tr *BotChannelMSTeams) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this BotChannelMSTeams
func (tr *BotChannelMSTeams) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BotChannelMSTeams
func (tr *BotChannelMSTeams) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BotChannelMSTeams
func (tr *BotChannelMSTeams) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BotChannelMSTeams
func (tr *BotChannelMSTeams) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BotChannelMSTeams
func (tr *BotChannelMSTeams) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BotChannelMSTeams using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BotChannelMSTeams) LateInitialize(attrs []byte) (bool, error) {
	params := &BotChannelMSTeamsParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BotChannelMSTeams) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BotChannelSlack
func (mg *BotChannelSlack) GetTerraformResourceType() string {
	return "azurerm_bot_channel_slack"
}

// GetConnectionDetailsMapping for this BotChannelSlack
func (tr *BotChannelSlack) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"client_secret": "spec.forProvider.clientSecretSecretRef", "signing_secret": "spec.forProvider.signingSecretSecretRef", "verification_token": "spec.forProvider.verificationTokenSecretRef"}
}

// GetObservation of this BotChannelSlack
func (tr *BotChannelSlack) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BotChannelSlack
func (tr *BotChannelSlack) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BotChannelSlack
func (tr *BotChannelSlack) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BotChannelSlack
func (tr *BotChannelSlack) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BotChannelSlack
func (tr *BotChannelSlack) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BotChannelSlack using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BotChannelSlack) LateInitialize(attrs []byte) (bool, error) {
	params := &BotChannelSlackParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BotChannelSlack) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BotChannelSMS
func (mg *BotChannelSMS) GetTerraformResourceType() string {
	return "azurerm_bot_channel_sms"
}

// GetConnectionDetailsMapping for this BotChannelSMS
func (tr *BotChannelSMS) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"sms_channel_auth_token": "spec.forProvider.smsChannelAuthTokenSecretRef"}
}

// GetObservation of this BotChannelSMS
func (tr *BotChannelSMS) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BotChannelSMS
func (tr *BotChannelSMS) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BotChannelSMS
func (tr *BotChannelSMS) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BotChannelSMS
func (tr *BotChannelSMS) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BotChannelSMS
func (tr *BotChannelSMS) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BotChannelSMS using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BotChannelSMS) LateInitialize(attrs []byte) (bool, error) {
	params := &BotChannelSMSParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BotChannelSMS) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BotChannelWebChat
func (mg *BotChannelWebChat) GetTerraformResourceType() string {
	return "azurerm_bot_channel_web_chat"
}

// GetConnectionDetailsMapping for this BotChannelWebChat
func (tr *BotChannelWebChat) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this BotChannelWebChat
func (tr *BotChannelWebChat) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BotChannelWebChat
func (tr *BotChannelWebChat) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BotChannelWebChat
func (tr *BotChannelWebChat) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BotChannelWebChat
func (tr *BotChannelWebChat) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BotChannelWebChat
func (tr *BotChannelWebChat) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BotChannelWebChat using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BotChannelWebChat) LateInitialize(attrs []byte) (bool, error) {
	params := &BotChannelWebChatParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BotChannelWebChat) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BotChannelsRegistration
func (mg *BotChannelsRegistration) GetTerraformResourceType() string {
	return "azurerm_bot_channels_registration"
}

// GetConnectionDetailsMapping for this BotChannelsRegistration
func (tr *BotChannelsRegistration) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"developer_app_insights_api_key": "spec.forProvider.developerAppInsightsApiKeySecretRef"}
}

// GetObservation of this BotChannelsRegistration
func (tr *BotChannelsRegistration) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BotChannelsRegistration
func (tr *BotChannelsRegistration) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BotChannelsRegistration
func (tr *BotChannelsRegistration) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BotChannelsRegistration
func (tr *BotChannelsRegistration) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BotChannelsRegistration
func (tr *BotChannelsRegistration) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BotChannelsRegistration using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BotChannelsRegistration) LateInitialize(attrs []byte) (bool, error) {
	params := &BotChannelsRegistrationParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BotChannelsRegistration) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BotConnection
func (mg *BotConnection) GetTerraformResourceType() string {
	return "azurerm_bot_connection"
}

// GetConnectionDetailsMapping for this BotConnection
func (tr *BotConnection) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"client_secret": "spec.forProvider.clientSecretSecretRef"}
}

// GetObservation of this BotConnection
func (tr *BotConnection) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BotConnection
func (tr *BotConnection) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BotConnection
func (tr *BotConnection) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BotConnection
func (tr *BotConnection) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BotConnection
func (tr *BotConnection) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BotConnection using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BotConnection) LateInitialize(attrs []byte) (bool, error) {
	params := &BotConnectionParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BotConnection) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BotWebApp
func (mg *BotWebApp) GetTerraformResourceType() string {
	return "azurerm_bot_web_app"
}

// GetConnectionDetailsMapping for this BotWebApp
func (tr *BotWebApp) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"developer_app_insights_api_key": "spec.forProvider.developerAppInsightsApiKeySecretRef", "luis_key": "spec.forProvider.luisKeySecretRef"}
}

// GetObservation of this BotWebApp
func (tr *BotWebApp) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BotWebApp
func (tr *BotWebApp) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BotWebApp
func (tr *BotWebApp) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BotWebApp
func (tr *BotWebApp) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BotWebApp
func (tr *BotWebApp) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BotWebApp using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BotWebApp) LateInitialize(attrs []byte) (bool, error) {
	params := &BotWebAppParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BotWebApp) GetTerraformSchemaVersion() int {
	return 0
}