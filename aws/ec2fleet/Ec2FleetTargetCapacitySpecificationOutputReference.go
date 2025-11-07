// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2fleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/ec2fleet/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2FleetTargetCapacitySpecificationOutputReference interface {
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
	DefaultTargetCapacityType() *string
	SetDefaultTargetCapacityType(val *string)
	DefaultTargetCapacityTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *Ec2FleetTargetCapacitySpecification
	SetInternalValue(val *Ec2FleetTargetCapacitySpecification)
	OnDemandTargetCapacity() *float64
	SetOnDemandTargetCapacity(val *float64)
	OnDemandTargetCapacityInput() *float64
	SpotTargetCapacity() *float64
	SetSpotTargetCapacity(val *float64)
	SpotTargetCapacityInput() *float64
	TargetCapacityUnitType() *string
	SetTargetCapacityUnitType(val *string)
	TargetCapacityUnitTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TotalTargetCapacity() *float64
	SetTotalTargetCapacity(val *float64)
	TotalTargetCapacityInput() *float64
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
	ResetOnDemandTargetCapacity()
	ResetSpotTargetCapacity()
	ResetTargetCapacityUnitType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2FleetTargetCapacitySpecificationOutputReference
type jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) DefaultTargetCapacityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTargetCapacityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) DefaultTargetCapacityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTargetCapacityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) InternalValue() *Ec2FleetTargetCapacitySpecification {
	var returns *Ec2FleetTargetCapacitySpecification
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) OnDemandTargetCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) OnDemandTargetCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandTargetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) SpotTargetCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) SpotTargetCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotTargetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) TargetCapacityUnitType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCapacityUnitType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) TargetCapacityUnitTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCapacityUnitTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) TotalTargetCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) TotalTargetCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalTargetCapacityInput",
		&returns,
	)
	return returns
}


func NewEc2FleetTargetCapacitySpecificationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Ec2FleetTargetCapacitySpecificationOutputReference {
	_init_.Initialize()

	if err := validateNewEc2FleetTargetCapacitySpecificationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ec2Fleet.Ec2FleetTargetCapacitySpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2FleetTargetCapacitySpecificationOutputReference_Override(e Ec2FleetTargetCapacitySpecificationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ec2Fleet.Ec2FleetTargetCapacitySpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference)SetDefaultTargetCapacityType(val *string) {
	if err := j.validateSetDefaultTargetCapacityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTargetCapacityType",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference)SetInternalValue(val *Ec2FleetTargetCapacitySpecification) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference)SetOnDemandTargetCapacity(val *float64) {
	if err := j.validateSetOnDemandTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandTargetCapacity",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference)SetSpotTargetCapacity(val *float64) {
	if err := j.validateSetSpotTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotTargetCapacity",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference)SetTargetCapacityUnitType(val *string) {
	if err := j.validateSetTargetCapacityUnitTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetCapacityUnitType",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference)SetTotalTargetCapacity(val *float64) {
	if err := j.validateSetTotalTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalTargetCapacity",
		val,
	)
}

func (e *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) ResetOnDemandTargetCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetOnDemandTargetCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) ResetSpotTargetCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetSpotTargetCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) ResetTargetCapacityUnitType() {
	_jsii_.InvokeVoid(
		e,
		"resetTargetCapacityUnitType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2FleetTargetCapacitySpecificationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

