// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrmanagedscalingpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/emrmanagedscalingpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmrManagedScalingPolicyComputeLimitsOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaximumCapacityUnits() *float64
	SetMaximumCapacityUnits(val *float64)
	MaximumCapacityUnitsInput() *float64
	MaximumCoreCapacityUnits() *float64
	SetMaximumCoreCapacityUnits(val *float64)
	MaximumCoreCapacityUnitsInput() *float64
	MaximumOndemandCapacityUnits() *float64
	SetMaximumOndemandCapacityUnits(val *float64)
	MaximumOndemandCapacityUnitsInput() *float64
	MinimumCapacityUnits() *float64
	SetMinimumCapacityUnits(val *float64)
	MinimumCapacityUnitsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UnitType() *string
	SetUnitType(val *string)
	UnitTypeInput() *string
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
	ResetMaximumCoreCapacityUnits()
	ResetMaximumOndemandCapacityUnits()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EmrManagedScalingPolicyComputeLimitsOutputReference
type jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) MaximumCapacityUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumCapacityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) MaximumCapacityUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumCapacityUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) MaximumCoreCapacityUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumCoreCapacityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) MaximumCoreCapacityUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumCoreCapacityUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) MaximumOndemandCapacityUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumOndemandCapacityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) MaximumOndemandCapacityUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumOndemandCapacityUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) MinimumCapacityUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumCapacityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) MinimumCapacityUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumCapacityUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) UnitType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) UnitTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitTypeInput",
		&returns,
	)
	return returns
}


func NewEmrManagedScalingPolicyComputeLimitsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EmrManagedScalingPolicyComputeLimitsOutputReference {
	_init_.Initialize()

	if err := validateNewEmrManagedScalingPolicyComputeLimitsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.emrManagedScalingPolicy.EmrManagedScalingPolicyComputeLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEmrManagedScalingPolicyComputeLimitsOutputReference_Override(e EmrManagedScalingPolicyComputeLimitsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.emrManagedScalingPolicy.EmrManagedScalingPolicyComputeLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference)SetMaximumCapacityUnits(val *float64) {
	if err := j.validateSetMaximumCapacityUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumCapacityUnits",
		val,
	)
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference)SetMaximumCoreCapacityUnits(val *float64) {
	if err := j.validateSetMaximumCoreCapacityUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumCoreCapacityUnits",
		val,
	)
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference)SetMaximumOndemandCapacityUnits(val *float64) {
	if err := j.validateSetMaximumOndemandCapacityUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumOndemandCapacityUnits",
		val,
	)
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference)SetMinimumCapacityUnits(val *float64) {
	if err := j.validateSetMinimumCapacityUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumCapacityUnits",
		val,
	)
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference)SetUnitType(val *string) {
	if err := j.validateSetUnitTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unitType",
		val,
	)
}

func (e *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) ResetMaximumCoreCapacityUnits() {
	_jsii_.InvokeVoid(
		e,
		"resetMaximumCoreCapacityUnits",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) ResetMaximumOndemandCapacityUnits() {
	_jsii_.InvokeVoid(
		e,
		"resetMaximumOndemandCapacityUnits",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EmrManagedScalingPolicyComputeLimitsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

