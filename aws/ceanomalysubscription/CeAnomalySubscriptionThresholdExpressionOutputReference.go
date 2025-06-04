// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ceanomalysubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/ceanomalysubscription/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CeAnomalySubscriptionThresholdExpressionOutputReference interface {
	cdktf.ComplexObject
	And() CeAnomalySubscriptionThresholdExpressionAndList
	AndInput() interface{}
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
	CostCategory() CeAnomalySubscriptionThresholdExpressionCostCategoryOutputReference
	CostCategoryInput() *CeAnomalySubscriptionThresholdExpressionCostCategory
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Dimension() CeAnomalySubscriptionThresholdExpressionDimensionOutputReference
	DimensionInput() *CeAnomalySubscriptionThresholdExpressionDimension
	// Experimental.
	Fqn() *string
	InternalValue() *CeAnomalySubscriptionThresholdExpression
	SetInternalValue(val *CeAnomalySubscriptionThresholdExpression)
	Not() CeAnomalySubscriptionThresholdExpressionNotOutputReference
	NotInput() *CeAnomalySubscriptionThresholdExpressionNot
	Or() CeAnomalySubscriptionThresholdExpressionOrList
	OrInput() interface{}
	Tags() CeAnomalySubscriptionThresholdExpressionTagsOutputReference
	TagsInput() *CeAnomalySubscriptionThresholdExpressionTags
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
	PutAnd(value interface{})
	PutCostCategory(value *CeAnomalySubscriptionThresholdExpressionCostCategory)
	PutDimension(value *CeAnomalySubscriptionThresholdExpressionDimension)
	PutNot(value *CeAnomalySubscriptionThresholdExpressionNot)
	PutOr(value interface{})
	PutTags(value *CeAnomalySubscriptionThresholdExpressionTags)
	ResetAnd()
	ResetCostCategory()
	ResetDimension()
	ResetNot()
	ResetOr()
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CeAnomalySubscriptionThresholdExpressionOutputReference
type jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) And() CeAnomalySubscriptionThresholdExpressionAndList {
	var returns CeAnomalySubscriptionThresholdExpressionAndList
	_jsii_.Get(
		j,
		"and",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) AndInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"andInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) CostCategory() CeAnomalySubscriptionThresholdExpressionCostCategoryOutputReference {
	var returns CeAnomalySubscriptionThresholdExpressionCostCategoryOutputReference
	_jsii_.Get(
		j,
		"costCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) CostCategoryInput() *CeAnomalySubscriptionThresholdExpressionCostCategory {
	var returns *CeAnomalySubscriptionThresholdExpressionCostCategory
	_jsii_.Get(
		j,
		"costCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) Dimension() CeAnomalySubscriptionThresholdExpressionDimensionOutputReference {
	var returns CeAnomalySubscriptionThresholdExpressionDimensionOutputReference
	_jsii_.Get(
		j,
		"dimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) DimensionInput() *CeAnomalySubscriptionThresholdExpressionDimension {
	var returns *CeAnomalySubscriptionThresholdExpressionDimension
	_jsii_.Get(
		j,
		"dimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) InternalValue() *CeAnomalySubscriptionThresholdExpression {
	var returns *CeAnomalySubscriptionThresholdExpression
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) Not() CeAnomalySubscriptionThresholdExpressionNotOutputReference {
	var returns CeAnomalySubscriptionThresholdExpressionNotOutputReference
	_jsii_.Get(
		j,
		"not",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) NotInput() *CeAnomalySubscriptionThresholdExpressionNot {
	var returns *CeAnomalySubscriptionThresholdExpressionNot
	_jsii_.Get(
		j,
		"notInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) Or() CeAnomalySubscriptionThresholdExpressionOrList {
	var returns CeAnomalySubscriptionThresholdExpressionOrList
	_jsii_.Get(
		j,
		"or",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) OrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) Tags() CeAnomalySubscriptionThresholdExpressionTagsOutputReference {
	var returns CeAnomalySubscriptionThresholdExpressionTagsOutputReference
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) TagsInput() *CeAnomalySubscriptionThresholdExpressionTags {
	var returns *CeAnomalySubscriptionThresholdExpressionTags
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCeAnomalySubscriptionThresholdExpressionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CeAnomalySubscriptionThresholdExpressionOutputReference {
	_init_.Initialize()

	if err := validateNewCeAnomalySubscriptionThresholdExpressionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ceAnomalySubscription.CeAnomalySubscriptionThresholdExpressionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCeAnomalySubscriptionThresholdExpressionOutputReference_Override(c CeAnomalySubscriptionThresholdExpressionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ceAnomalySubscription.CeAnomalySubscriptionThresholdExpressionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference)SetInternalValue(val *CeAnomalySubscriptionThresholdExpression) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) PutAnd(value interface{}) {
	if err := c.validatePutAndParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAnd",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) PutCostCategory(value *CeAnomalySubscriptionThresholdExpressionCostCategory) {
	if err := c.validatePutCostCategoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCostCategory",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) PutDimension(value *CeAnomalySubscriptionThresholdExpressionDimension) {
	if err := c.validatePutDimensionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDimension",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) PutNot(value *CeAnomalySubscriptionThresholdExpressionNot) {
	if err := c.validatePutNotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNot",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) PutOr(value interface{}) {
	if err := c.validatePutOrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOr",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) PutTags(value *CeAnomalySubscriptionThresholdExpressionTags) {
	if err := c.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTags",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) ResetAnd() {
	_jsii_.InvokeVoid(
		c,
		"resetAnd",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) ResetCostCategory() {
	_jsii_.InvokeVoid(
		c,
		"resetCostCategory",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) ResetDimension() {
	_jsii_.InvokeVoid(
		c,
		"resetDimension",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) ResetNot() {
	_jsii_.InvokeVoid(
		c,
		"resetNot",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) ResetOr() {
	_jsii_.InvokeVoid(
		c,
		"resetOr",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

