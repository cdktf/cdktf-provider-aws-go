// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubconfigurationpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/securityhubconfigurationpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference interface {
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
	DisabledControlIdentifiers() *[]*string
	SetDisabledControlIdentifiers(val *[]*string)
	DisabledControlIdentifiersInput() *[]*string
	EnabledControlIdentifiers() *[]*string
	SetEnabledControlIdentifiers(val *[]*string)
	EnabledControlIdentifiersInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfiguration
	SetInternalValue(val *SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfiguration)
	SecurityControlCustomParameter() SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameterList
	SecurityControlCustomParameterInput() interface{}
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
	PutSecurityControlCustomParameter(value interface{})
	ResetDisabledControlIdentifiers()
	ResetEnabledControlIdentifiers()
	ResetSecurityControlCustomParameter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference
type jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) DisabledControlIdentifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disabledControlIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) DisabledControlIdentifiersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disabledControlIdentifiersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) EnabledControlIdentifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledControlIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) EnabledControlIdentifiersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledControlIdentifiersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) InternalValue() *SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfiguration {
	var returns *SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) SecurityControlCustomParameter() SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameterList {
	var returns SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameterList
	_jsii_.Get(
		j,
		"securityControlCustomParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) SecurityControlCustomParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"securityControlCustomParameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewSecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.securityhubConfigurationPolicy.SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference_Override(s SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.securityhubConfigurationPolicy.SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference)SetDisabledControlIdentifiers(val *[]*string) {
	if err := j.validateSetDisabledControlIdentifiersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabledControlIdentifiers",
		val,
	)
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference)SetEnabledControlIdentifiers(val *[]*string) {
	if err := j.validateSetEnabledControlIdentifiersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledControlIdentifiers",
		val,
	)
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference)SetInternalValue(val *SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) PutSecurityControlCustomParameter(value interface{}) {
	if err := s.validatePutSecurityControlCustomParameterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSecurityControlCustomParameter",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) ResetDisabledControlIdentifiers() {
	_jsii_.InvokeVoid(
		s,
		"resetDisabledControlIdentifiers",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) ResetEnabledControlIdentifiers() {
	_jsii_.InvokeVoid(
		s,
		"resetEnabledControlIdentifiers",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) ResetSecurityControlCustomParameter() {
	_jsii_.InvokeVoid(
		s,
		"resetSecurityControlCustomParameter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

