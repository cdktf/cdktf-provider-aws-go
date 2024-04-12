// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubconfigurationpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/securityhubconfigurationpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityhubConfigurationPolicyConfigurationPolicyOutputReference interface {
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
	EnabledStandardArns() *[]*string
	SetEnabledStandardArns(val *[]*string)
	EnabledStandardArnsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *SecurityhubConfigurationPolicyConfigurationPolicy
	SetInternalValue(val *SecurityhubConfigurationPolicyConfigurationPolicy)
	SecurityControlsConfiguration() SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference
	SecurityControlsConfigurationInput() *SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfiguration
	ServiceEnabled() interface{}
	SetServiceEnabled(val interface{})
	ServiceEnabledInput() interface{}
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
	PutSecurityControlsConfiguration(value *SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfiguration)
	ResetEnabledStandardArns()
	ResetSecurityControlsConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecurityhubConfigurationPolicyConfigurationPolicyOutputReference
type jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) EnabledStandardArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledStandardArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) EnabledStandardArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledStandardArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) InternalValue() *SecurityhubConfigurationPolicyConfigurationPolicy {
	var returns *SecurityhubConfigurationPolicyConfigurationPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) SecurityControlsConfiguration() SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference {
	var returns SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationOutputReference
	_jsii_.Get(
		j,
		"securityControlsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) SecurityControlsConfigurationInput() *SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfiguration {
	var returns *SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfiguration
	_jsii_.Get(
		j,
		"securityControlsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) ServiceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) ServiceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSecurityhubConfigurationPolicyConfigurationPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SecurityhubConfigurationPolicyConfigurationPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewSecurityhubConfigurationPolicyConfigurationPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.securityhubConfigurationPolicy.SecurityhubConfigurationPolicyConfigurationPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSecurityhubConfigurationPolicyConfigurationPolicyOutputReference_Override(s SecurityhubConfigurationPolicyConfigurationPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.securityhubConfigurationPolicy.SecurityhubConfigurationPolicyConfigurationPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference)SetEnabledStandardArns(val *[]*string) {
	if err := j.validateSetEnabledStandardArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledStandardArns",
		val,
	)
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference)SetInternalValue(val *SecurityhubConfigurationPolicyConfigurationPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference)SetServiceEnabled(val interface{}) {
	if err := j.validateSetServiceEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceEnabled",
		val,
	)
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) PutSecurityControlsConfiguration(value *SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfiguration) {
	if err := s.validatePutSecurityControlsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSecurityControlsConfiguration",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) ResetEnabledStandardArns() {
	_jsii_.InvokeVoid(
		s,
		"resetEnabledStandardArns",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) ResetSecurityControlsConfiguration() {
	_jsii_.InvokeVoid(
		s,
		"resetSecurityControlsConfiguration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SecurityhubConfigurationPolicyConfigurationPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

