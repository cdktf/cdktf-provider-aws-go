// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/kinesisanalyticsv2application/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference interface {
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
	InputStartingPosition() *string
	SetInputStartingPosition(val *string)
	InputStartingPositionInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	ResetInputStartingPosition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference
type jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference) InputStartingPosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputStartingPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference) InputStartingPositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputStartingPositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisanalyticsv2Application.Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewKinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference_Override(k Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisanalyticsv2Application.Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		k,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference)SetInputStartingPosition(val *string) {
	if err := j.validateSetInputStartingPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputStartingPosition",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference) ResetInputStartingPosition() {
	_jsii_.InvokeVoid(
		k,
		"resetInputStartingPosition",
		nil, // no parameters
	)
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputStartingPositionConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

