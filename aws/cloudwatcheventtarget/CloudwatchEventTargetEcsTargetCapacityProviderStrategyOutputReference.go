// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatcheventtarget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/cloudwatcheventtarget/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference interface {
	cdktf.ComplexObject
	Base() *float64
	SetBase(val *float64)
	BaseInput() *float64
	CapacityProvider() *string
	SetCapacityProvider(val *string)
	CapacityProviderInput() *string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Weight() *float64
	SetWeight(val *float64)
	WeightInput() *float64
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
	ResetBase()
	ResetWeight()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference
type jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) Base() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"base",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) BaseInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"baseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) CapacityProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) CapacityProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) WeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightInput",
		&returns,
	)
	return returns
}


func NewCloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference {
	_init_.Initialize()

	if err := validateNewCloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cloudwatchEventTarget.CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference_Override(c CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cloudwatchEventTarget.CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference)SetBase(val *float64) {
	if err := j.validateSetBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"base",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference)SetCapacityProvider(val *string) {
	if err := j.validateSetCapacityProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityProvider",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference)SetWeight(val *float64) {
	if err := j.validateSetWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weight",
		val,
	)
}

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) ResetBase() {
	_jsii_.InvokeVoid(
		c,
		"resetBase",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) ResetWeight() {
	_jsii_.InvokeVoid(
		c,
		"resetWeight",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudwatchEventTargetEcsTargetCapacityProviderStrategyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

