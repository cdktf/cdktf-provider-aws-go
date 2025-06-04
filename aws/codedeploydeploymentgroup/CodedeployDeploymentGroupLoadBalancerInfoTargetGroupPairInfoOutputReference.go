// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/codedeploydeploymentgroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference interface {
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
	InternalValue() *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfo
	SetInternalValue(val *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfo)
	ProdTrafficRoute() CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference
	ProdTrafficRouteInput() *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute
	TargetGroup() CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroupList
	TargetGroupInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TestTrafficRoute() CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference
	TestTrafficRouteInput() *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute
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
	PutProdTrafficRoute(value *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute)
	PutTargetGroup(value interface{})
	PutTestTrafficRoute(value *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute)
	ResetTestTrafficRoute()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference
type jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) InternalValue() *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfo {
	var returns *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) ProdTrafficRoute() CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference {
	var returns CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference
	_jsii_.Get(
		j,
		"prodTrafficRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) ProdTrafficRouteInput() *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute {
	var returns *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute
	_jsii_.Get(
		j,
		"prodTrafficRouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) TargetGroup() CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroupList {
	var returns CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroupList
	_jsii_.Get(
		j,
		"targetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) TargetGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) TestTrafficRoute() CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference {
	var returns CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference
	_jsii_.Get(
		j,
		"testTrafficRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) TestTrafficRouteInput() *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute {
	var returns *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute
	_jsii_.Get(
		j,
		"testTrafficRouteInput",
		&returns,
	)
	return returns
}


func NewCodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference {
	_init_.Initialize()

	if err := validateNewCodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.codedeployDeploymentGroup.CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference_Override(c CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.codedeployDeploymentGroup.CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference)SetInternalValue(val *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) PutProdTrafficRoute(value *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute) {
	if err := c.validatePutProdTrafficRouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putProdTrafficRoute",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) PutTargetGroup(value interface{}) {
	if err := c.validatePutTargetGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTargetGroup",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) PutTestTrafficRoute(value *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute) {
	if err := c.validatePutTestTrafficRouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTestTrafficRoute",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) ResetTestTrafficRoute() {
	_jsii_.InvokeVoid(
		c,
		"resetTestTrafficRoute",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

