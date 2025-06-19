// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configconfigrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/configconfigrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfigConfigRuleSourceCustomPolicyDetailsOutputReference interface {
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
	EnableDebugLogDelivery() interface{}
	SetEnableDebugLogDelivery(val interface{})
	EnableDebugLogDeliveryInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ConfigConfigRuleSourceCustomPolicyDetails
	SetInternalValue(val *ConfigConfigRuleSourceCustomPolicyDetails)
	PolicyRuntime() *string
	SetPolicyRuntime(val *string)
	PolicyRuntimeInput() *string
	PolicyText() *string
	SetPolicyText(val *string)
	PolicyTextInput() *string
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
	ResetEnableDebugLogDelivery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConfigConfigRuleSourceCustomPolicyDetailsOutputReference
type jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) EnableDebugLogDelivery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDebugLogDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) EnableDebugLogDeliveryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDebugLogDeliveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) InternalValue() *ConfigConfigRuleSourceCustomPolicyDetails {
	var returns *ConfigConfigRuleSourceCustomPolicyDetails
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) PolicyRuntime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) PolicyRuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) PolicyText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) PolicyTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConfigConfigRuleSourceCustomPolicyDetailsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ConfigConfigRuleSourceCustomPolicyDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewConfigConfigRuleSourceCustomPolicyDetailsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.configConfigRule.ConfigConfigRuleSourceCustomPolicyDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConfigConfigRuleSourceCustomPolicyDetailsOutputReference_Override(c ConfigConfigRuleSourceCustomPolicyDetailsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.configConfigRule.ConfigConfigRuleSourceCustomPolicyDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference)SetEnableDebugLogDelivery(val interface{}) {
	if err := j.validateSetEnableDebugLogDeliveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDebugLogDelivery",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference)SetInternalValue(val *ConfigConfigRuleSourceCustomPolicyDetails) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference)SetPolicyRuntime(val *string) {
	if err := j.validateSetPolicyRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyRuntime",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference)SetPolicyText(val *string) {
	if err := j.validateSetPolicyTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyText",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) ResetEnableDebugLogDelivery() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableDebugLogDelivery",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigRuleSourceCustomPolicyDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

