// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecscapacityprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/ecscapacityprovider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference interface {
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
	InstanceWarmupPeriod() *float64
	SetInstanceWarmupPeriod(val *float64)
	InstanceWarmupPeriodInput() *float64
	InternalValue() *EcsCapacityProviderAutoScalingGroupProviderManagedScaling
	SetInternalValue(val *EcsCapacityProviderAutoScalingGroupProviderManagedScaling)
	MaximumScalingStepSize() *float64
	SetMaximumScalingStepSize(val *float64)
	MaximumScalingStepSizeInput() *float64
	MinimumScalingStepSize() *float64
	SetMinimumScalingStepSize(val *float64)
	MinimumScalingStepSizeInput() *float64
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	TargetCapacity() *float64
	SetTargetCapacity(val *float64)
	TargetCapacityInput() *float64
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
	ResetInstanceWarmupPeriod()
	ResetMaximumScalingStepSize()
	ResetMinimumScalingStepSize()
	ResetStatus()
	ResetTargetCapacity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference
type jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) InstanceWarmupPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceWarmupPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) InstanceWarmupPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceWarmupPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) InternalValue() *EcsCapacityProviderAutoScalingGroupProviderManagedScaling {
	var returns *EcsCapacityProviderAutoScalingGroupProviderManagedScaling
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) MaximumScalingStepSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumScalingStepSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) MaximumScalingStepSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumScalingStepSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) MinimumScalingStepSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumScalingStepSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) MinimumScalingStepSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumScalingStepSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) TargetCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) TargetCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference {
	_init_.Initialize()

	if err := validateNewEcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ecsCapacityProvider.EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference_Override(e EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ecsCapacityProvider.EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference)SetInstanceWarmupPeriod(val *float64) {
	if err := j.validateSetInstanceWarmupPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceWarmupPeriod",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference)SetInternalValue(val *EcsCapacityProviderAutoScalingGroupProviderManagedScaling) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference)SetMaximumScalingStepSize(val *float64) {
	if err := j.validateSetMaximumScalingStepSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumScalingStepSize",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference)SetMinimumScalingStepSize(val *float64) {
	if err := j.validateSetMinimumScalingStepSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumScalingStepSize",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference)SetTargetCapacity(val *float64) {
	if err := j.validateSetTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetCapacity",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) ResetInstanceWarmupPeriod() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceWarmupPeriod",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) ResetMaximumScalingStepSize() {
	_jsii_.InvokeVoid(
		e,
		"resetMaximumScalingStepSize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) ResetMinimumScalingStepSize() {
	_jsii_.InvokeVoid(
		e,
		"resetMinimumScalingStepSize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		e,
		"resetStatus",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) ResetTargetCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetTargetCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

