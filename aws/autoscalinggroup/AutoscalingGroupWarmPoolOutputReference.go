// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/autoscalinggroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutoscalingGroupWarmPoolOutputReference interface {
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
	InstanceReusePolicy() AutoscalingGroupWarmPoolInstanceReusePolicyOutputReference
	InstanceReusePolicyInput() *AutoscalingGroupWarmPoolInstanceReusePolicy
	InternalValue() *AutoscalingGroupWarmPool
	SetInternalValue(val *AutoscalingGroupWarmPool)
	MaxGroupPreparedCapacity() *float64
	SetMaxGroupPreparedCapacity(val *float64)
	MaxGroupPreparedCapacityInput() *float64
	MinSize() *float64
	SetMinSize(val *float64)
	MinSizeInput() *float64
	PoolState() *string
	SetPoolState(val *string)
	PoolStateInput() *string
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
	PutInstanceReusePolicy(value *AutoscalingGroupWarmPoolInstanceReusePolicy)
	ResetInstanceReusePolicy()
	ResetMaxGroupPreparedCapacity()
	ResetMinSize()
	ResetPoolState()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutoscalingGroupWarmPoolOutputReference
type jsiiProxy_AutoscalingGroupWarmPoolOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) InstanceReusePolicy() AutoscalingGroupWarmPoolInstanceReusePolicyOutputReference {
	var returns AutoscalingGroupWarmPoolInstanceReusePolicyOutputReference
	_jsii_.Get(
		j,
		"instanceReusePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) InstanceReusePolicyInput() *AutoscalingGroupWarmPoolInstanceReusePolicy {
	var returns *AutoscalingGroupWarmPoolInstanceReusePolicy
	_jsii_.Get(
		j,
		"instanceReusePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) InternalValue() *AutoscalingGroupWarmPool {
	var returns *AutoscalingGroupWarmPool
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) MaxGroupPreparedCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxGroupPreparedCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) MaxGroupPreparedCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxGroupPreparedCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) MinSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) PoolState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) PoolStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutoscalingGroupWarmPoolOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AutoscalingGroupWarmPoolOutputReference {
	_init_.Initialize()

	if err := validateNewAutoscalingGroupWarmPoolOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoscalingGroupWarmPoolOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.autoscalingGroup.AutoscalingGroupWarmPoolOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAutoscalingGroupWarmPoolOutputReference_Override(a AutoscalingGroupWarmPoolOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.autoscalingGroup.AutoscalingGroupWarmPoolOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AutoscalingGroupWarmPoolOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupWarmPoolOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupWarmPoolOutputReference)SetInternalValue(val *AutoscalingGroupWarmPool) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupWarmPoolOutputReference)SetMaxGroupPreparedCapacity(val *float64) {
	if err := j.validateSetMaxGroupPreparedCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxGroupPreparedCapacity",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupWarmPoolOutputReference)SetMinSize(val *float64) {
	if err := j.validateSetMinSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupWarmPoolOutputReference)SetPoolState(val *string) {
	if err := j.validateSetPoolStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"poolState",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupWarmPoolOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupWarmPoolOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) PutInstanceReusePolicy(value *AutoscalingGroupWarmPoolInstanceReusePolicy) {
	if err := a.validatePutInstanceReusePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInstanceReusePolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) ResetInstanceReusePolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceReusePolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) ResetMaxGroupPreparedCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxGroupPreparedCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) ResetMinSize() {
	_jsii_.InvokeVoid(
		a,
		"resetMinSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) ResetPoolState() {
	_jsii_.InvokeVoid(
		a,
		"resetPoolState",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupWarmPoolOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

