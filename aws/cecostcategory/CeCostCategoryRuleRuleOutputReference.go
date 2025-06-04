// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cecostcategory

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/cecostcategory/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CeCostCategoryRuleRuleOutputReference interface {
	cdktf.ComplexObject
	And() CeCostCategoryRuleRuleAndList
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
	CostCategory() CeCostCategoryRuleRuleCostCategoryOutputReference
	CostCategoryInput() *CeCostCategoryRuleRuleCostCategory
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Dimension() CeCostCategoryRuleRuleDimensionOutputReference
	DimensionInput() *CeCostCategoryRuleRuleDimension
	// Experimental.
	Fqn() *string
	InternalValue() *CeCostCategoryRuleRule
	SetInternalValue(val *CeCostCategoryRuleRule)
	Not() CeCostCategoryRuleRuleNotOutputReference
	NotInput() *CeCostCategoryRuleRuleNot
	Or() CeCostCategoryRuleRuleOrList
	OrInput() interface{}
	Tags() CeCostCategoryRuleRuleTagsOutputReference
	TagsInput() *CeCostCategoryRuleRuleTags
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
	PutCostCategory(value *CeCostCategoryRuleRuleCostCategory)
	PutDimension(value *CeCostCategoryRuleRuleDimension)
	PutNot(value *CeCostCategoryRuleRuleNot)
	PutOr(value interface{})
	PutTags(value *CeCostCategoryRuleRuleTags)
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

// The jsii proxy struct for CeCostCategoryRuleRuleOutputReference
type jsiiProxy_CeCostCategoryRuleRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CeCostCategoryRuleRuleOutputReference) And() CeCostCategoryRuleRuleAndList {
	var returns CeCostCategoryRuleRuleAndList
	_jsii_.Get(
		j,
		"and",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeCostCategoryRuleRuleOutputReference) AndInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"andInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeCostCategoryRuleRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeCostCategoryRuleRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeCostCategoryRuleRuleOutputReference) CostCategory() CeCostCategoryRuleRuleCostCategoryOutputReference {
	var returns CeCostCategoryRuleRuleCostCategoryOutputReference
	_jsii_.Get(
		j,
		"costCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeCostCategoryRuleRuleOutputReference) CostCategoryInput() *CeCostCategoryRuleRuleCostCategory {
	var returns *CeCostCategoryRuleRuleCostCategory
	_jsii_.Get(
		j,
		"costCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeCostCategoryRuleRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeCostCategoryRuleRuleOutputReference) Dimension() CeCostCategoryRuleRuleDimensionOutputReference {
	var returns CeCostCategoryRuleRuleDimensionOutputReference
	_jsii_.Get(
		j,
		"dimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeCostCategoryRuleRuleOutputReference) DimensionInput() *CeCostCategoryRuleRuleDimension {
	var returns *CeCostCategoryRuleRuleDimension
	_jsii_.Get(
		j,
		"dimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeCostCategoryRuleRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeCostCategoryRuleRuleOutputReference) InternalValue() *CeCostCategoryRuleRule {
	var returns *CeCostCategoryRuleRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeCostCategoryRuleRuleOutputReference) Not() CeCostCategoryRuleRuleNotOutputReference {
	var returns CeCostCategoryRuleRuleNotOutputReference
	_jsii_.Get(
		j,
		"not",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeCostCategoryRuleRuleOutputReference) NotInput() *CeCostCategoryRuleRuleNot {
	var returns *CeCostCategoryRuleRuleNot
	_jsii_.Get(
		j,
		"notInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeCostCategoryRuleRuleOutputReference) Or() CeCostCategoryRuleRuleOrList {
	var returns CeCostCategoryRuleRuleOrList
	_jsii_.Get(
		j,
		"or",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeCostCategoryRuleRuleOutputReference) OrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeCostCategoryRuleRuleOutputReference) Tags() CeCostCategoryRuleRuleTagsOutputReference {
	var returns CeCostCategoryRuleRuleTagsOutputReference
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeCostCategoryRuleRuleOutputReference) TagsInput() *CeCostCategoryRuleRuleTags {
	var returns *CeCostCategoryRuleRuleTags
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeCostCategoryRuleRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CeCostCategoryRuleRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCeCostCategoryRuleRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CeCostCategoryRuleRuleOutputReference {
	_init_.Initialize()

	if err := validateNewCeCostCategoryRuleRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CeCostCategoryRuleRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ceCostCategory.CeCostCategoryRuleRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCeCostCategoryRuleRuleOutputReference_Override(c CeCostCategoryRuleRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ceCostCategory.CeCostCategoryRuleRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CeCostCategoryRuleRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CeCostCategoryRuleRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CeCostCategoryRuleRuleOutputReference)SetInternalValue(val *CeCostCategoryRuleRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CeCostCategoryRuleRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CeCostCategoryRuleRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) PutAnd(value interface{}) {
	if err := c.validatePutAndParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAnd",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) PutCostCategory(value *CeCostCategoryRuleRuleCostCategory) {
	if err := c.validatePutCostCategoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCostCategory",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) PutDimension(value *CeCostCategoryRuleRuleDimension) {
	if err := c.validatePutDimensionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDimension",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) PutNot(value *CeCostCategoryRuleRuleNot) {
	if err := c.validatePutNotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNot",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) PutOr(value interface{}) {
	if err := c.validatePutOrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOr",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) PutTags(value *CeCostCategoryRuleRuleTags) {
	if err := c.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTags",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) ResetAnd() {
	_jsii_.InvokeVoid(
		c,
		"resetAnd",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) ResetCostCategory() {
	_jsii_.InvokeVoid(
		c,
		"resetCostCategory",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) ResetDimension() {
	_jsii_.InvokeVoid(
		c,
		"resetDimension",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) ResetNot() {
	_jsii_.InvokeVoid(
		c,
		"resetNot",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) ResetOr() {
	_jsii_.InvokeVoid(
		c,
		"resetOr",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CeCostCategoryRuleRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

