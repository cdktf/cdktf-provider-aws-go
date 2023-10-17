// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/sagemakerdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerDomainDomainSettingsOutputReference interface {
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
	ExecutionRoleIdentityConfig() *string
	SetExecutionRoleIdentityConfig(val *string)
	ExecutionRoleIdentityConfigInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *SagemakerDomainDomainSettings
	SetInternalValue(val *SagemakerDomainDomainSettings)
	RStudioServerProDomainSettings() SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference
	RStudioServerProDomainSettingsInput() *SagemakerDomainDomainSettingsRStudioServerProDomainSettings
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutRStudioServerProDomainSettings(value *SagemakerDomainDomainSettingsRStudioServerProDomainSettings)
	ResetExecutionRoleIdentityConfig()
	ResetRStudioServerProDomainSettings()
	ResetSecurityGroupIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerDomainDomainSettingsOutputReference
type jsiiProxy_SagemakerDomainDomainSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) ExecutionRoleIdentityConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleIdentityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) ExecutionRoleIdentityConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleIdentityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) InternalValue() *SagemakerDomainDomainSettings {
	var returns *SagemakerDomainDomainSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) RStudioServerProDomainSettings() SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference {
	var returns SagemakerDomainDomainSettingsRStudioServerProDomainSettingsOutputReference
	_jsii_.Get(
		j,
		"rStudioServerProDomainSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) RStudioServerProDomainSettingsInput() *SagemakerDomainDomainSettingsRStudioServerProDomainSettings {
	var returns *SagemakerDomainDomainSettingsRStudioServerProDomainSettings
	_jsii_.Get(
		j,
		"rStudioServerProDomainSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerDomainDomainSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerDomainDomainSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerDomainDomainSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerDomainDomainSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerDomain.SagemakerDomainDomainSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerDomainDomainSettingsOutputReference_Override(s SagemakerDomainDomainSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerDomain.SagemakerDomainDomainSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsOutputReference)SetExecutionRoleIdentityConfig(val *string) {
	if err := j.validateSetExecutionRoleIdentityConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleIdentityConfig",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsOutputReference)SetInternalValue(val *SagemakerDomainDomainSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsOutputReference)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDomainSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) PutRStudioServerProDomainSettings(value *SagemakerDomainDomainSettingsRStudioServerProDomainSettings) {
	if err := s.validatePutRStudioServerProDomainSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRStudioServerProDomainSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) ResetExecutionRoleIdentityConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetExecutionRoleIdentityConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) ResetRStudioServerProDomainSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetRStudioServerProDomainSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		s,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerDomainDomainSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

