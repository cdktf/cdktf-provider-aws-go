// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/sagemakerdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference interface {
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
	// Experimental.
	Fqn() *string
	IdleSettings() SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsOutputReference
	IdleSettingsInput() *SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettings
	InternalValue() *SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagement
	SetInternalValue(val *SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagement)
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
	PutIdleSettings(value *SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettings)
	ResetIdleSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference
type jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) IdleSettings() SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsOutputReference {
	var returns SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettingsOutputReference
	_jsii_.Get(
		j,
		"idleSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) IdleSettingsInput() *SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettings {
	var returns *SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettings
	_jsii_.Get(
		j,
		"idleSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) InternalValue() *SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagement {
	var returns *SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagement
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerDomain.SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference_Override(s SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerDomain.SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference)SetInternalValue(val *SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagement) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) PutIdleSettings(value *SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementIdleSettings) {
	if err := s.validatePutIdleSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIdleSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) ResetIdleSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetIdleSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsAppLifecycleManagementOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

