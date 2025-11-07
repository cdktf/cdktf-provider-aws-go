// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbtargetgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/lbtargetgroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference interface {
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
	InternalValue() *LbTargetGroupTargetGroupHealthUnhealthyStateRouting
	SetInternalValue(val *LbTargetGroupTargetGroupHealthUnhealthyStateRouting)
	MinimumHealthyTargetsCount() *float64
	SetMinimumHealthyTargetsCount(val *float64)
	MinimumHealthyTargetsCountInput() *float64
	MinimumHealthyTargetsPercentage() *string
	SetMinimumHealthyTargetsPercentage(val *string)
	MinimumHealthyTargetsPercentageInput() *string
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
	ResetMinimumHealthyTargetsCount()
	ResetMinimumHealthyTargetsPercentage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference
type jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) InternalValue() *LbTargetGroupTargetGroupHealthUnhealthyStateRouting {
	var returns *LbTargetGroupTargetGroupHealthUnhealthyStateRouting
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) MinimumHealthyTargetsCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumHealthyTargetsCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) MinimumHealthyTargetsCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumHealthyTargetsCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) MinimumHealthyTargetsPercentage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumHealthyTargetsPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) MinimumHealthyTargetsPercentageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumHealthyTargetsPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference {
	_init_.Initialize()

	if err := validateNewLbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.lbTargetGroup.LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference_Override(l LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lbTargetGroup.LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference)SetInternalValue(val *LbTargetGroupTargetGroupHealthUnhealthyStateRouting) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference)SetMinimumHealthyTargetsCount(val *float64) {
	if err := j.validateSetMinimumHealthyTargetsCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumHealthyTargetsCount",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference)SetMinimumHealthyTargetsPercentage(val *string) {
	if err := j.validateSetMinimumHealthyTargetsPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumHealthyTargetsPercentage",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) ResetMinimumHealthyTargetsCount() {
	_jsii_.InvokeVoid(
		l,
		"resetMinimumHealthyTargetsCount",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) ResetMinimumHealthyTargetsPercentage() {
	_jsii_.InvokeVoid(
		l,
		"resetMinimumHealthyTargetsPercentage",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbTargetGroupTargetGroupHealthUnhealthyStateRoutingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

