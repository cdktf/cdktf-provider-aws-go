// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/codedeploydeploymentgroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference interface {
	cdktf.ComplexObject
	ActionOnTimeout() *string
	SetActionOnTimeout(val *string)
	ActionOnTimeoutInput() *string
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
	InternalValue() *CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption
	SetInternalValue(val *CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WaitTimeInMinutes() *float64
	SetWaitTimeInMinutes(val *float64)
	WaitTimeInMinutesInput() *float64
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
	ResetActionOnTimeout()
	ResetWaitTimeInMinutes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference
type jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) ActionOnTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionOnTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) ActionOnTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionOnTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) InternalValue() *CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption {
	var returns *CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) WaitTimeInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitTimeInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) WaitTimeInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitTimeInMinutesInput",
		&returns,
	)
	return returns
}


func NewCodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference {
	_init_.Initialize()

	if err := validateNewCodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.codedeployDeploymentGroup.CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference_Override(c CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.codedeployDeploymentGroup.CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference)SetActionOnTimeout(val *string) {
	if err := j.validateSetActionOnTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionOnTimeout",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference)SetInternalValue(val *CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference)SetWaitTimeInMinutes(val *float64) {
	if err := j.validateSetWaitTimeInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitTimeInMinutes",
		val,
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) ResetActionOnTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetActionOnTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) ResetWaitTimeInMinutes() {
	_jsii_.InvokeVoid(
		c,
		"resetWaitTimeInMinutes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

