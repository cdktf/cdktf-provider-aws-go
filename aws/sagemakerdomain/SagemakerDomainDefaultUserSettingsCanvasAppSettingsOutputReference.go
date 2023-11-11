// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/sagemakerdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference interface {
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
	DirectDeploySettings() SagemakerDomainDefaultUserSettingsCanvasAppSettingsDirectDeploySettingsOutputReference
	DirectDeploySettingsInput() *SagemakerDomainDefaultUserSettingsCanvasAppSettingsDirectDeploySettings
	// Experimental.
	Fqn() *string
	IdentityProviderOauthSettings() SagemakerDomainDefaultUserSettingsCanvasAppSettingsIdentityProviderOauthSettingsList
	IdentityProviderOauthSettingsInput() interface{}
	InternalValue() *SagemakerDomainDefaultUserSettingsCanvasAppSettings
	SetInternalValue(val *SagemakerDomainDefaultUserSettingsCanvasAppSettings)
	KendraSettings() SagemakerDomainDefaultUserSettingsCanvasAppSettingsKendraSettingsOutputReference
	KendraSettingsInput() *SagemakerDomainDefaultUserSettingsCanvasAppSettingsKendraSettings
	ModelRegisterSettings() SagemakerDomainDefaultUserSettingsCanvasAppSettingsModelRegisterSettingsOutputReference
	ModelRegisterSettingsInput() *SagemakerDomainDefaultUserSettingsCanvasAppSettingsModelRegisterSettings
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeSeriesForecastingSettings() SagemakerDomainDefaultUserSettingsCanvasAppSettingsTimeSeriesForecastingSettingsOutputReference
	TimeSeriesForecastingSettingsInput() *SagemakerDomainDefaultUserSettingsCanvasAppSettingsTimeSeriesForecastingSettings
	WorkspaceSettings() SagemakerDomainDefaultUserSettingsCanvasAppSettingsWorkspaceSettingsOutputReference
	WorkspaceSettingsInput() *SagemakerDomainDefaultUserSettingsCanvasAppSettingsWorkspaceSettings
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
	PutDirectDeploySettings(value *SagemakerDomainDefaultUserSettingsCanvasAppSettingsDirectDeploySettings)
	PutIdentityProviderOauthSettings(value interface{})
	PutKendraSettings(value *SagemakerDomainDefaultUserSettingsCanvasAppSettingsKendraSettings)
	PutModelRegisterSettings(value *SagemakerDomainDefaultUserSettingsCanvasAppSettingsModelRegisterSettings)
	PutTimeSeriesForecastingSettings(value *SagemakerDomainDefaultUserSettingsCanvasAppSettingsTimeSeriesForecastingSettings)
	PutWorkspaceSettings(value *SagemakerDomainDefaultUserSettingsCanvasAppSettingsWorkspaceSettings)
	ResetDirectDeploySettings()
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

// The jsii proxy struct for SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference
type jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) DirectDeploySettings() SagemakerDomainDefaultUserSettingsCanvasAppSettingsDirectDeploySettingsOutputReference {
	var returns SagemakerDomainDefaultUserSettingsCanvasAppSettingsDirectDeploySettingsOutputReference
	_jsii_.Get(
		j,
		"directDeploySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) DirectDeploySettingsInput() *SagemakerDomainDefaultUserSettingsCanvasAppSettingsDirectDeploySettings {
	var returns *SagemakerDomainDefaultUserSettingsCanvasAppSettingsDirectDeploySettings
	_jsii_.Get(
		j,
		"directDeploySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) IdentityProviderOauthSettings() SagemakerDomainDefaultUserSettingsCanvasAppSettingsIdentityProviderOauthSettingsList {
	var returns SagemakerDomainDefaultUserSettingsCanvasAppSettingsIdentityProviderOauthSettingsList
	_jsii_.Get(
		j,
		"identityProviderOauthSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) IdentityProviderOauthSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identityProviderOauthSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) InternalValue() *SagemakerDomainDefaultUserSettingsCanvasAppSettings {
	var returns *SagemakerDomainDefaultUserSettingsCanvasAppSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) KendraSettings() SagemakerDomainDefaultUserSettingsCanvasAppSettingsKendraSettingsOutputReference {
	var returns SagemakerDomainDefaultUserSettingsCanvasAppSettingsKendraSettingsOutputReference
	_jsii_.Get(
		j,
		"kendraSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) KendraSettingsInput() *SagemakerDomainDefaultUserSettingsCanvasAppSettingsKendraSettings {
	var returns *SagemakerDomainDefaultUserSettingsCanvasAppSettingsKendraSettings
	_jsii_.Get(
		j,
		"kendraSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) ModelRegisterSettings() SagemakerDomainDefaultUserSettingsCanvasAppSettingsModelRegisterSettingsOutputReference {
	var returns SagemakerDomainDefaultUserSettingsCanvasAppSettingsModelRegisterSettingsOutputReference
	_jsii_.Get(
		j,
		"modelRegisterSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) ModelRegisterSettingsInput() *SagemakerDomainDefaultUserSettingsCanvasAppSettingsModelRegisterSettings {
	var returns *SagemakerDomainDefaultUserSettingsCanvasAppSettingsModelRegisterSettings
	_jsii_.Get(
		j,
		"modelRegisterSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) TimeSeriesForecastingSettings() SagemakerDomainDefaultUserSettingsCanvasAppSettingsTimeSeriesForecastingSettingsOutputReference {
	var returns SagemakerDomainDefaultUserSettingsCanvasAppSettingsTimeSeriesForecastingSettingsOutputReference
	_jsii_.Get(
		j,
		"timeSeriesForecastingSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) TimeSeriesForecastingSettingsInput() *SagemakerDomainDefaultUserSettingsCanvasAppSettingsTimeSeriesForecastingSettings {
	var returns *SagemakerDomainDefaultUserSettingsCanvasAppSettingsTimeSeriesForecastingSettings
	_jsii_.Get(
		j,
		"timeSeriesForecastingSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) WorkspaceSettings() SagemakerDomainDefaultUserSettingsCanvasAppSettingsWorkspaceSettingsOutputReference {
	var returns SagemakerDomainDefaultUserSettingsCanvasAppSettingsWorkspaceSettingsOutputReference
	_jsii_.Get(
		j,
		"workspaceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) WorkspaceSettingsInput() *SagemakerDomainDefaultUserSettingsCanvasAppSettingsWorkspaceSettings {
	var returns *SagemakerDomainDefaultUserSettingsCanvasAppSettingsWorkspaceSettings
	_jsii_.Get(
		j,
		"workspaceSettingsInput",
		&returns,
	)
	return returns
}


func NewSagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerDomain.SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference_Override(s SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerDomain.SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference)SetInternalValue(val *SagemakerDomainDefaultUserSettingsCanvasAppSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) PutDirectDeploySettings(value *SagemakerDomainDefaultUserSettingsCanvasAppSettingsDirectDeploySettings) {
	if err := s.validatePutDirectDeploySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDirectDeploySettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) PutIdentityProviderOauthSettings(value interface{}) {
	if err := s.validatePutIdentityProviderOauthSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIdentityProviderOauthSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) PutKendraSettings(value *SagemakerDomainDefaultUserSettingsCanvasAppSettingsKendraSettings) {
	if err := s.validatePutKendraSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putKendraSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) PutModelRegisterSettings(value *SagemakerDomainDefaultUserSettingsCanvasAppSettingsModelRegisterSettings) {
	if err := s.validatePutModelRegisterSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putModelRegisterSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) PutTimeSeriesForecastingSettings(value *SagemakerDomainDefaultUserSettingsCanvasAppSettingsTimeSeriesForecastingSettings) {
	if err := s.validatePutTimeSeriesForecastingSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeSeriesForecastingSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) PutWorkspaceSettings(value *SagemakerDomainDefaultUserSettingsCanvasAppSettingsWorkspaceSettings) {
	if err := s.validatePutWorkspaceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putWorkspaceSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) ResetDirectDeploySettings() {
	_jsii_.InvokeVoid(
		s,
		"resetDirectDeploySettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) ResetIdentityProviderOauthSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetIdentityProviderOauthSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) ResetKendraSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetKendraSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) ResetModelRegisterSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetModelRegisterSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) ResetTimeSeriesForecastingSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeSeriesForecastingSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) ResetWorkspaceSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetWorkspaceSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCanvasAppSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

