// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsexpressgatewayservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/ecsexpressgatewayservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EcsExpressGatewayServiceScalingTargetOutputReference interface {
	cdktf.ComplexObject
	AutoScalingMetric() *string
	SetAutoScalingMetric(val *string)
	AutoScalingMetricInput() *string
	AutoScalingTargetValue() *float64
	SetAutoScalingTargetValue(val *float64)
	AutoScalingTargetValueInput() *float64
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
	MaxTaskCount() *float64
	SetMaxTaskCount(val *float64)
	MaxTaskCountInput() *float64
	MinTaskCount() *float64
	SetMinTaskCount(val *float64)
	MinTaskCountInput() *float64
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetAutoScalingMetric()
	ResetAutoScalingTargetValue()
	ResetMaxTaskCount()
	ResetMinTaskCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsExpressGatewayServiceScalingTargetOutputReference
type jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) AutoScalingMetric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) AutoScalingMetricInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingMetricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) AutoScalingTargetValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoScalingTargetValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) AutoScalingTargetValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoScalingTargetValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) MaxTaskCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTaskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) MaxTaskCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTaskCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) MinTaskCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTaskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) MinTaskCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTaskCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEcsExpressGatewayServiceScalingTargetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EcsExpressGatewayServiceScalingTargetOutputReference {
	_init_.Initialize()

	if err := validateNewEcsExpressGatewayServiceScalingTargetOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ecsExpressGatewayService.EcsExpressGatewayServiceScalingTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEcsExpressGatewayServiceScalingTargetOutputReference_Override(e EcsExpressGatewayServiceScalingTargetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ecsExpressGatewayService.EcsExpressGatewayServiceScalingTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference)SetAutoScalingMetric(val *string) {
	if err := j.validateSetAutoScalingMetricParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoScalingMetric",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference)SetAutoScalingTargetValue(val *float64) {
	if err := j.validateSetAutoScalingTargetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoScalingTargetValue",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference)SetMaxTaskCount(val *float64) {
	if err := j.validateSetMaxTaskCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTaskCount",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference)SetMinTaskCount(val *float64) {
	if err := j.validateSetMinTaskCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTaskCount",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) ResetAutoScalingMetric() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoScalingMetric",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) ResetAutoScalingTargetValue() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoScalingTargetValue",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) ResetMaxTaskCount() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxTaskCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) ResetMinTaskCount() {
	_jsii_.InvokeVoid(
		e,
		"resetMinTaskCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsExpressGatewayServiceScalingTargetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

