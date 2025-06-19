// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/codedeploydeploymentgroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodedeployDeploymentGroupLoadBalancerInfoOutputReference interface {
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
	ElbInfo() CodedeployDeploymentGroupLoadBalancerInfoElbInfoList
	ElbInfoInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *CodedeployDeploymentGroupLoadBalancerInfo
	SetInternalValue(val *CodedeployDeploymentGroupLoadBalancerInfo)
	TargetGroupInfo() CodedeployDeploymentGroupLoadBalancerInfoTargetGroupInfoList
	TargetGroupInfoInput() interface{}
	TargetGroupPairInfo() CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference
	TargetGroupPairInfoInput() *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfo
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
	PutElbInfo(value interface{})
	PutTargetGroupInfo(value interface{})
	PutTargetGroupPairInfo(value *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfo)
	ResetElbInfo()
	ResetTargetGroupInfo()
	ResetTargetGroupPairInfo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CodedeployDeploymentGroupLoadBalancerInfoOutputReference
type jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) ElbInfo() CodedeployDeploymentGroupLoadBalancerInfoElbInfoList {
	var returns CodedeployDeploymentGroupLoadBalancerInfoElbInfoList
	_jsii_.Get(
		j,
		"elbInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) ElbInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elbInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) InternalValue() *CodedeployDeploymentGroupLoadBalancerInfo {
	var returns *CodedeployDeploymentGroupLoadBalancerInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) TargetGroupInfo() CodedeployDeploymentGroupLoadBalancerInfoTargetGroupInfoList {
	var returns CodedeployDeploymentGroupLoadBalancerInfoTargetGroupInfoList
	_jsii_.Get(
		j,
		"targetGroupInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) TargetGroupInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetGroupInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) TargetGroupPairInfo() CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference {
	var returns CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference
	_jsii_.Get(
		j,
		"targetGroupPairInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) TargetGroupPairInfoInput() *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfo {
	var returns *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfo
	_jsii_.Get(
		j,
		"targetGroupPairInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCodedeployDeploymentGroupLoadBalancerInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CodedeployDeploymentGroupLoadBalancerInfoOutputReference {
	_init_.Initialize()

	if err := validateNewCodedeployDeploymentGroupLoadBalancerInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.codedeployDeploymentGroup.CodedeployDeploymentGroupLoadBalancerInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCodedeployDeploymentGroupLoadBalancerInfoOutputReference_Override(c CodedeployDeploymentGroupLoadBalancerInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.codedeployDeploymentGroup.CodedeployDeploymentGroupLoadBalancerInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference)SetInternalValue(val *CodedeployDeploymentGroupLoadBalancerInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) PutElbInfo(value interface{}) {
	if err := c.validatePutElbInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putElbInfo",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) PutTargetGroupInfo(value interface{}) {
	if err := c.validatePutTargetGroupInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTargetGroupInfo",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) PutTargetGroupPairInfo(value *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfo) {
	if err := c.validatePutTargetGroupPairInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTargetGroupPairInfo",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) ResetElbInfo() {
	_jsii_.InvokeVoid(
		c,
		"resetElbInfo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) ResetTargetGroupInfo() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetGroupInfo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) ResetTargetGroupPairInfo() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetGroupPairInfo",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

