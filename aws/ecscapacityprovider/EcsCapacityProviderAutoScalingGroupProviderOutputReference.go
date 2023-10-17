// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecscapacityprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/ecscapacityprovider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EcsCapacityProviderAutoScalingGroupProviderOutputReference interface {
	cdktf.ComplexObject
	AutoScalingGroupArn() *string
	SetAutoScalingGroupArn(val *string)
	AutoScalingGroupArnInput() *string
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
	InternalValue() *EcsCapacityProviderAutoScalingGroupProvider
	SetInternalValue(val *EcsCapacityProviderAutoScalingGroupProvider)
	ManagedScaling() EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference
	ManagedScalingInput() *EcsCapacityProviderAutoScalingGroupProviderManagedScaling
	ManagedTerminationProtection() *string
	SetManagedTerminationProtection(val *string)
	ManagedTerminationProtectionInput() *string
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
	PutManagedScaling(value *EcsCapacityProviderAutoScalingGroupProviderManagedScaling)
	ResetManagedScaling()
	ResetManagedTerminationProtection()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsCapacityProviderAutoScalingGroupProviderOutputReference
type jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) AutoScalingGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) AutoScalingGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) InternalValue() *EcsCapacityProviderAutoScalingGroupProvider {
	var returns *EcsCapacityProviderAutoScalingGroupProvider
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) ManagedScaling() EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference {
	var returns EcsCapacityProviderAutoScalingGroupProviderManagedScalingOutputReference
	_jsii_.Get(
		j,
		"managedScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) ManagedScalingInput() *EcsCapacityProviderAutoScalingGroupProviderManagedScaling {
	var returns *EcsCapacityProviderAutoScalingGroupProviderManagedScaling
	_jsii_.Get(
		j,
		"managedScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) ManagedTerminationProtection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedTerminationProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) ManagedTerminationProtectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedTerminationProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEcsCapacityProviderAutoScalingGroupProviderOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EcsCapacityProviderAutoScalingGroupProviderOutputReference {
	_init_.Initialize()

	if err := validateNewEcsCapacityProviderAutoScalingGroupProviderOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ecsCapacityProvider.EcsCapacityProviderAutoScalingGroupProviderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEcsCapacityProviderAutoScalingGroupProviderOutputReference_Override(e EcsCapacityProviderAutoScalingGroupProviderOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ecsCapacityProvider.EcsCapacityProviderAutoScalingGroupProviderOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference)SetAutoScalingGroupArn(val *string) {
	if err := j.validateSetAutoScalingGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoScalingGroupArn",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference)SetInternalValue(val *EcsCapacityProviderAutoScalingGroupProvider) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference)SetManagedTerminationProtection(val *string) {
	if err := j.validateSetManagedTerminationProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedTerminationProtection",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) PutManagedScaling(value *EcsCapacityProviderAutoScalingGroupProviderManagedScaling) {
	if err := e.validatePutManagedScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putManagedScaling",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) ResetManagedScaling() {
	_jsii_.InvokeVoid(
		e,
		"resetManagedScaling",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) ResetManagedTerminationProtection() {
	_jsii_.InvokeVoid(
		e,
		"resetManagedTerminationProtection",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EcsCapacityProviderAutoScalingGroupProviderOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

