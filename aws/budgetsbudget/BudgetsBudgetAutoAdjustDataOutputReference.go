// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package budgetsbudget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/budgetsbudget/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BudgetsBudgetAutoAdjustDataOutputReference interface {
	cdktf.ComplexObject
	AutoAdjustType() *string
	SetAutoAdjustType(val *string)
	AutoAdjustTypeInput() *string
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
	HistoricalOptions() BudgetsBudgetAutoAdjustDataHistoricalOptionsOutputReference
	HistoricalOptionsInput() *BudgetsBudgetAutoAdjustDataHistoricalOptions
	InternalValue() *BudgetsBudgetAutoAdjustData
	SetInternalValue(val *BudgetsBudgetAutoAdjustData)
	LastAutoAdjustTime() *string
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
	PutHistoricalOptions(value *BudgetsBudgetAutoAdjustDataHistoricalOptions)
	ResetHistoricalOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BudgetsBudgetAutoAdjustDataOutputReference
type jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) AutoAdjustType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoAdjustType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) AutoAdjustTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoAdjustTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) HistoricalOptions() BudgetsBudgetAutoAdjustDataHistoricalOptionsOutputReference {
	var returns BudgetsBudgetAutoAdjustDataHistoricalOptionsOutputReference
	_jsii_.Get(
		j,
		"historicalOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) HistoricalOptionsInput() *BudgetsBudgetAutoAdjustDataHistoricalOptions {
	var returns *BudgetsBudgetAutoAdjustDataHistoricalOptions
	_jsii_.Get(
		j,
		"historicalOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) InternalValue() *BudgetsBudgetAutoAdjustData {
	var returns *BudgetsBudgetAutoAdjustData
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) LastAutoAdjustTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastAutoAdjustTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBudgetsBudgetAutoAdjustDataOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BudgetsBudgetAutoAdjustDataOutputReference {
	_init_.Initialize()

	if err := validateNewBudgetsBudgetAutoAdjustDataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudgetAutoAdjustDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBudgetsBudgetAutoAdjustDataOutputReference_Override(b BudgetsBudgetAutoAdjustDataOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.budgetsBudget.BudgetsBudgetAutoAdjustDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference)SetAutoAdjustType(val *string) {
	if err := j.validateSetAutoAdjustTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoAdjustType",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference)SetInternalValue(val *BudgetsBudgetAutoAdjustData) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) PutHistoricalOptions(value *BudgetsBudgetAutoAdjustDataHistoricalOptions) {
	if err := b.validatePutHistoricalOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putHistoricalOptions",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) ResetHistoricalOptions() {
	_jsii_.InvokeVoid(
		b,
		"resetHistoricalOptions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetAutoAdjustDataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

