// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/sagemakerdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference interface {
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
	IdleTimeoutInMinutes() *float64
	SetIdleTimeoutInMinutes(val *float64)
	IdleTimeoutInMinutesInput() *float64
	InternalValue() *SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettings
	SetInternalValue(val *SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettings)
	LifecycleManagement() *string
	SetLifecycleManagement(val *string)
	LifecycleManagementInput() *string
	MaxIdleTimeoutInMinutes() *float64
	SetMaxIdleTimeoutInMinutes(val *float64)
	MaxIdleTimeoutInMinutesInput() *float64
	MinIdleTimeoutInMinutes() *float64
	SetMinIdleTimeoutInMinutes(val *float64)
	MinIdleTimeoutInMinutesInput() *float64
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
	ResetIdleTimeoutInMinutes()
	ResetLifecycleManagement()
	ResetMaxIdleTimeoutInMinutes()
	ResetMinIdleTimeoutInMinutes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference
type jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) IdleTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) IdleTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) InternalValue() *SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettings {
	var returns *SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) LifecycleManagement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) LifecycleManagementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) MaxIdleTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) MaxIdleTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) MinIdleTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIdleTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) MinIdleTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIdleTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerDomain.SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference_Override(s SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sagemakerDomain.SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference)SetIdleTimeoutInMinutes(val *float64) {
	if err := j.validateSetIdleTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference)SetInternalValue(val *SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference)SetLifecycleManagement(val *string) {
	if err := j.validateSetLifecycleManagementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleManagement",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference)SetMaxIdleTimeoutInMinutes(val *float64) {
	if err := j.validateSetMaxIdleTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIdleTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference)SetMinIdleTimeoutInMinutes(val *float64) {
	if err := j.validateSetMinIdleTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minIdleTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) ResetIdleTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		s,
		"resetIdleTimeoutInMinutes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) ResetLifecycleManagement() {
	_jsii_.InvokeVoid(
		s,
		"resetLifecycleManagement",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) ResetMaxIdleTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxIdleTimeoutInMinutes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) ResetMinIdleTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		s,
		"resetMinIdleTimeoutInMinutes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

