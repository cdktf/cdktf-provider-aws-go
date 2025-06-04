// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakeruserprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/sagemakeruserprofile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DirectDeploySettings() SagemakerUserProfileUserSettingsCanvasAppSettingsDirectDeploySettingsOutputReference
	DirectDeploySettingsInput() *SagemakerUserProfileUserSettingsCanvasAppSettingsDirectDeploySettings
	EmrServerlessSettings() SagemakerUserProfileUserSettingsCanvasAppSettingsEmrServerlessSettingsOutputReference
	EmrServerlessSettingsInput() *SagemakerUserProfileUserSettingsCanvasAppSettingsEmrServerlessSettings
	// Experimental.
	Fqn() *string
	GenerativeAiSettings() SagemakerUserProfileUserSettingsCanvasAppSettingsGenerativeAiSettingsOutputReference
	GenerativeAiSettingsInput() *SagemakerUserProfileUserSettingsCanvasAppSettingsGenerativeAiSettings
	IdentityProviderOauthSettings() SagemakerUserProfileUserSettingsCanvasAppSettingsIdentityProviderOauthSettingsList
	IdentityProviderOauthSettingsInput() interface{}
	InternalValue() *SagemakerUserProfileUserSettingsCanvasAppSettings
	SetInternalValue(val *SagemakerUserProfileUserSettingsCanvasAppSettings)
	KendraSettings() SagemakerUserProfileUserSettingsCanvasAppSettingsKendraSettingsOutputReference
	KendraSettingsInput() *SagemakerUserProfileUserSettingsCanvasAppSettingsKendraSettings
	ModelRegisterSettings() SagemakerUserProfileUserSettingsCanvasAppSettingsModelRegisterSettingsOutputReference
	ModelRegisterSettingsInput() *SagemakerUserProfileUserSettingsCanvasAppSettingsModelRegisterSettings
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeSeriesForecastingSettings() SagemakerUserProfileUserSettingsCanvasAppSettingsTimeSeriesForecastingSettingsOutputReference
	TimeSeriesForecastingSettingsInput() *SagemakerUserProfileUserSettingsCanvasAppSettingsTimeSeriesForecastingSettings
	WorkspaceSettings() SagemakerUserProfileUserSettingsCanvasAppSettingsWorkspaceSettingsOutputReference
	WorkspaceSettingsInput() *SagemakerUserProfileUserSettingsCanvasAppSettingsWorkspaceSettings
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutDirectDeploySettings(value *SagemakerUserProfileUserSettingsCanvasAppSettingsDirectDeploySettings)
	PutEmrServerlessSettings(value *SagemakerUserProfileUserSettingsCanvasAppSettingsEmrServerlessSettings)
	PutGenerativeAiSettings(value *SagemakerUserProfileUserSettingsCanvasAppSettingsGenerativeAiSettings)
	PutIdentityProviderOauthSettings(value interface{})
	PutKendraSettings(value *SagemakerUserProfileUserSettingsCanvasAppSettingsKendraSettings)
	PutModelRegisterSettings(value *SagemakerUserProfileUserSettingsCanvasAppSettingsModelRegisterSettings)
	PutTimeSeriesForecastingSettings(value *SagemakerUserProfileUserSettingsCanvasAppSettingsTimeSeriesForecastingSettings)
	PutWorkspaceSettings(value *SagemakerUserProfileUserSettingsCanvasAppSettingsWorkspaceSettings)
	ResetDirectDeploySettings()
	ResetEmrServerlessSettings()
	ResetGenerativeAiSettings()
	ResetIdentityProviderOauthSettings()
	ResetKendraSettings()
	ResetModelRegisterSettings()
	ResetTimeSeriesForecastingSettings()
	ResetWorkspaceSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference
type jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) DirectDeploySettings() SagemakerUserProfileUserSettingsCanvasAppSettingsDirectDeploySettingsOutputReference {
	var returns SagemakerUserProfileUserSettingsCanvasAppSettingsDirectDeploySettingsOutputReference
	_jsii_.Get(
		j,
		"directDeploySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) DirectDeploySettingsInput() *SagemakerUserProfileUserSettingsCanvasAppSettingsDirectDeploySettings {
	var returns *SagemakerUserProfileUserSettingsCanvasAppSettingsDirectDeploySettings
	_jsii_.Get(
		j,
		"directDeploySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) EmrServerlessSettings() SagemakerUserProfileUserSettingsCanvasAppSettingsEmrServerlessSettingsOutputReference {
	var returns SagemakerUserProfileUserSettingsCanvasAppSettingsEmrServerlessSettingsOutputReference
	_jsii_.Get(
		j,
		"emrServerlessSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) EmrServerlessSettingsInput() *SagemakerUserProfileUserSettingsCanvasAppSettingsEmrServerlessSettings {
	var returns *SagemakerUserProfileUserSettingsCanvasAppSettingsEmrServerlessSettings
	_jsii_.Get(
		j,
		"emrServerlessSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) GenerativeAiSettings() SagemakerUserProfileUserSettingsCanvasAppSettingsGenerativeAiSettingsOutputReference {
	var returns SagemakerUserProfileUserSettingsCanvasAppSettingsGenerativeAiSettingsOutputReference
	_jsii_.Get(
		j,
		"generativeAiSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) GenerativeAiSettingsInput() *SagemakerUserProfileUserSettingsCanvasAppSettingsGenerativeAiSettings {
	var returns *SagemakerUserProfileUserSettingsCanvasAppSettingsGenerativeAiSettings
	_jsii_.Get(
		j,
		"generativeAiSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) IdentityProviderOauthSettings() SagemakerUserProfileUserSettingsCanvasAppSettingsIdentityProviderOauthSettingsList {
	var returns SagemakerUserProfileUserSettingsCanvasAppSettingsIdentityProviderOauthSettingsList
	_jsii_.Get(
		j,
		"identityProviderOauthSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) IdentityProviderOauthSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identityProviderOauthSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) InternalValue() *SagemakerUserProfileUserSettingsCanvasAppSettings {
	var returns *SagemakerUserProfileUserSettingsCanvasAppSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) KendraSettings() SagemakerUserProfileUserSettingsCanvasAppSettingsKendraSettingsOutputReference {
	var returns SagemakerUserProfileUserSettingsCanvasAppSettingsKendraSettingsOutputReference
	_jsii_.Get(
		j,
		"kendraSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) KendraSettingsInput() *SagemakerUserProfileUserSettingsCanvasAppSettingsKendraSettings {
	var returns *SagemakerUserProfileUserSettingsCanvasAppSettingsKendraSettings
	_jsii_.Get(
		j,
		"kendraSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) ModelRegisterSettings() SagemakerUserProfileUserSettingsCanvasAppSettingsModelRegisterSettingsOutputReference {
	var returns SagemakerUserProfileUserSettingsCanvasAppSettingsModelRegisterSettingsOutputReference
	_jsii_.Get(
		j,
		"modelRegisterSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) ModelRegisterSettingsInput() *SagemakerUserProfileUserSettingsCanvasAppSettingsModelRegisterSettings {
	var returns *SagemakerUserProfileUserSettingsCanvasAppSettingsModelRegisterSettings
	_jsii_.Get(
		j,
		"modelRegisterSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) TimeSeriesForecastingSettings() SagemakerUserProfileUserSettingsCanvasAppSettingsTimeSeriesForecastingSettingsOutputReference {
	var returns SagemakerUserProfileUserSettingsCanvasAppSettingsTimeSeriesForecastingSettingsOutputReference
	_jsii_.Get(
		j,
		"timeSeriesForecastingSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) TimeSeriesForecastingSettingsInput() *SagemakerUserProfileUserSettingsCanvasAppSettingsTimeSeriesForecastingSettings {
	var returns *SagemakerUserProfileUserSettingsCanvasAppSettingsTimeSeriesForecastingSettings
	_jsii_.Get(
		j,
		"timeSeriesForecastingSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) WorkspaceSettings() SagemakerUserProfileUserSettingsCanvasAppSettingsWorkspaceSettingsOutputReference {
	var returns SagemakerUserProfileUserSettingsCanvasAppSettingsWorkspaceSettingsOutputReference
	_jsii_.Get(
		j,
		"workspaceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) WorkspaceSettingsInput() *SagemakerUserProfileUserSettingsCanvasAppSettingsWorkspaceSettings {
	var returns *SagemakerUserProfileUserSettingsCanvasAppSettingsWorkspaceSettings
	_jsii_.Get(
		j,
		"workspaceSettingsInput",
		&returns,
	)
	return returns
}


func NewSagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerUserProfileUserSettingsCanvasAppSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerUserProfile.SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference_Override(s SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerUserProfile.SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference)SetInternalValue(val *SagemakerUserProfileUserSettingsCanvasAppSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) PutDirectDeploySettings(value *SagemakerUserProfileUserSettingsCanvasAppSettingsDirectDeploySettings) {
	if err := s.validatePutDirectDeploySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDirectDeploySettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) PutEmrServerlessSettings(value *SagemakerUserProfileUserSettingsCanvasAppSettingsEmrServerlessSettings) {
	if err := s.validatePutEmrServerlessSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEmrServerlessSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) PutGenerativeAiSettings(value *SagemakerUserProfileUserSettingsCanvasAppSettingsGenerativeAiSettings) {
	if err := s.validatePutGenerativeAiSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGenerativeAiSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) PutIdentityProviderOauthSettings(value interface{}) {
	if err := s.validatePutIdentityProviderOauthSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIdentityProviderOauthSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) PutKendraSettings(value *SagemakerUserProfileUserSettingsCanvasAppSettingsKendraSettings) {
	if err := s.validatePutKendraSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putKendraSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) PutModelRegisterSettings(value *SagemakerUserProfileUserSettingsCanvasAppSettingsModelRegisterSettings) {
	if err := s.validatePutModelRegisterSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putModelRegisterSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) PutTimeSeriesForecastingSettings(value *SagemakerUserProfileUserSettingsCanvasAppSettingsTimeSeriesForecastingSettings) {
	if err := s.validatePutTimeSeriesForecastingSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeSeriesForecastingSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) PutWorkspaceSettings(value *SagemakerUserProfileUserSettingsCanvasAppSettingsWorkspaceSettings) {
	if err := s.validatePutWorkspaceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putWorkspaceSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) ResetDirectDeploySettings() {
	_jsii_.InvokeVoid(
		s,
		"resetDirectDeploySettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) ResetEmrServerlessSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetEmrServerlessSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) ResetGenerativeAiSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetGenerativeAiSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) ResetIdentityProviderOauthSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetIdentityProviderOauthSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) ResetKendraSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetKendraSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) ResetModelRegisterSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetModelRegisterSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) ResetTimeSeriesForecastingSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeSeriesForecastingSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) ResetWorkspaceSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetWorkspaceSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerUserProfileUserSettingsCanvasAppSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

