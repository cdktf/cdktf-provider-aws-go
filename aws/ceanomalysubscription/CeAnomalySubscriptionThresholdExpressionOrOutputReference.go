// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ceanomalysubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/ceanomalysubscription/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CeAnomalySubscriptionThresholdExpressionOrOutputReference interface {
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
	CostCategory() CeAnomalySubscriptionThresholdExpressionOrCostCategoryOutputReference
	CostCategoryInput() *CeAnomalySubscriptionThresholdExpressionOrCostCategory
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Dimension() CeAnomalySubscriptionThresholdExpressionOrDimensionOutputReference
	DimensionInput() *CeAnomalySubscriptionThresholdExpressionOrDimension
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Tags() CeAnomalySubscriptionThresholdExpressionOrTagsOutputReference
	TagsInput() *CeAnomalySubscriptionThresholdExpressionOrTags
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
	PutCostCategory(value *CeAnomalySubscriptionThresholdExpressionOrCostCategory)
	PutDimension(value *CeAnomalySubscriptionThresholdExpressionOrDimension)
	PutTags(value *CeAnomalySubscriptionThresholdExpressionOrTags)
	ResetCostCategory()
	ResetDimension()
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

// The jsii proxy struct for CeAnomalySubscriptionThresholdExpressionOrOutputReference
type jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) CostCategory() CeAnomalySubscriptionThresholdExpressionOrCostCategoryOutputReference {
	var returns CeAnomalySubscriptionThresholdExpressionOrCostCategoryOutputReference
	_jsii_.Get(
		j,
		"costCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) CostCategoryInput() *CeAnomalySubscriptionThresholdExpressionOrCostCategory {
	var returns *CeAnomalySubscriptionThresholdExpressionOrCostCategory
	_jsii_.Get(
		j,
		"costCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) Dimension() CeAnomalySubscriptionThresholdExpressionOrDimensionOutputReference {
	var returns CeAnomalySubscriptionThresholdExpressionOrDimensionOutputReference
	_jsii_.Get(
		j,
		"dimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) DimensionInput() *CeAnomalySubscriptionThresholdExpressionOrDimension {
	var returns *CeAnomalySubscriptionThresholdExpressionOrDimension
	_jsii_.Get(
		j,
		"dimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) Tags() CeAnomalySubscriptionThresholdExpressionOrTagsOutputReference {
	var returns CeAnomalySubscriptionThresholdExpressionOrTagsOutputReference
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) TagsInput() *CeAnomalySubscriptionThresholdExpressionOrTags {
	var returns *CeAnomalySubscriptionThresholdExpressionOrTags
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCeAnomalySubscriptionThresholdExpressionOrOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CeAnomalySubscriptionThresholdExpressionOrOutputReference {
	_init_.Initialize()

	if err := validateNewCeAnomalySubscriptionThresholdExpressionOrOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ceAnomalySubscription.CeAnomalySubscriptionThresholdExpressionOrOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCeAnomalySubscriptionThresholdExpressionOrOutputReference_Override(c CeAnomalySubscriptionThresholdExpressionOrOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ceAnomalySubscription.CeAnomalySubscriptionThresholdExpressionOrOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) PutCostCategory(value *CeAnomalySubscriptionThresholdExpressionOrCostCategory) {
	if err := c.validatePutCostCategoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCostCategory",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) PutDimension(value *CeAnomalySubscriptionThresholdExpressionOrDimension) {
	if err := c.validatePutDimensionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDimension",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) PutTags(value *CeAnomalySubscriptionThresholdExpressionOrTags) {
	if err := c.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTags",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) ResetCostCategory() {
	_jsii_.InvokeVoid(
		c,
		"resetCostCategory",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) ResetDimension() {
	_jsii_.InvokeVoid(
		c,
		"resetDimension",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CeAnomalySubscriptionThresholdExpressionOrOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

