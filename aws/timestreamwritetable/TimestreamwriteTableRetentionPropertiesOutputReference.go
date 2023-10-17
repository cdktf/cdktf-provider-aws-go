// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package timestreamwritetable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/timestreamwritetable/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TimestreamwriteTableRetentionPropertiesOutputReference interface {
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
	InternalValue() *TimestreamwriteTableRetentionProperties
	SetInternalValue(val *TimestreamwriteTableRetentionProperties)
	MagneticStoreRetentionPeriodInDays() *float64
	SetMagneticStoreRetentionPeriodInDays(val *float64)
	MagneticStoreRetentionPeriodInDaysInput() *float64
	MemoryStoreRetentionPeriodInHours() *float64
	SetMemoryStoreRetentionPeriodInHours(val *float64)
	MemoryStoreRetentionPeriodInHoursInput() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TimestreamwriteTableRetentionPropertiesOutputReference
type jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) InternalValue() *TimestreamwriteTableRetentionProperties {
	var returns *TimestreamwriteTableRetentionProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) MagneticStoreRetentionPeriodInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"magneticStoreRetentionPeriodInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) MagneticStoreRetentionPeriodInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"magneticStoreRetentionPeriodInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) MemoryStoreRetentionPeriodInHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryStoreRetentionPeriodInHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) MemoryStoreRetentionPeriodInHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryStoreRetentionPeriodInHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTimestreamwriteTableRetentionPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TimestreamwriteTableRetentionPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewTimestreamwriteTableRetentionPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.timestreamwriteTable.TimestreamwriteTableRetentionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTimestreamwriteTableRetentionPropertiesOutputReference_Override(t TimestreamwriteTableRetentionPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.timestreamwriteTable.TimestreamwriteTableRetentionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference)SetInternalValue(val *TimestreamwriteTableRetentionProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference)SetMagneticStoreRetentionPeriodInDays(val *float64) {
	if err := j.validateSetMagneticStoreRetentionPeriodInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"magneticStoreRetentionPeriodInDays",
		val,
	)
}

func (j *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference)SetMemoryStoreRetentionPeriodInHours(val *float64) {
	if err := j.validateSetMemoryStoreRetentionPeriodInHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryStoreRetentionPeriodInHours",
		val,
	)
}

func (j *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TimestreamwriteTableRetentionPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

