// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package guarddutyfilter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/guarddutyfilter/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GuarddutyFilterFindingCriteriaCriterionOutputReference interface {
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
	EqualTo() *[]*string
	SetEqualTo(val *[]*string)
	EqualToInput() *[]*string
	Field() *string
	SetField(val *string)
	FieldInput() *string
	// Experimental.
	Fqn() *string
	GreaterThan() *string
	SetGreaterThan(val *string)
	GreaterThanInput() *string
	GreaterThanOrEqual() *string
	SetGreaterThanOrEqual(val *string)
	GreaterThanOrEqualInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LessThan() *string
	SetLessThan(val *string)
	LessThanInput() *string
	LessThanOrEqual() *string
	SetLessThanOrEqual(val *string)
	LessThanOrEqualInput() *string
	NotEquals() *[]*string
	SetNotEquals(val *[]*string)
	NotEqualsInput() *[]*string
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
	ResetEqualTo()
	ResetGreaterThan()
	ResetGreaterThanOrEqual()
	ResetLessThan()
	ResetLessThanOrEqual()
	ResetNotEquals()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GuarddutyFilterFindingCriteriaCriterionOutputReference
type jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) EqualTo() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"equalTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) EqualToInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"equalToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) Field() *string {
	var returns *string
	_jsii_.Get(
		j,
		"field",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) FieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GreaterThan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"greaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GreaterThanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"greaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GreaterThanOrEqual() *string {
	var returns *string
	_jsii_.Get(
		j,
		"greaterThanOrEqual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GreaterThanOrEqualInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"greaterThanOrEqualInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) LessThan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lessThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) LessThanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lessThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) LessThanOrEqual() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lessThanOrEqual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) LessThanOrEqualInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lessThanOrEqualInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) NotEquals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) NotEqualsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGuarddutyFilterFindingCriteriaCriterionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GuarddutyFilterFindingCriteriaCriterionOutputReference {
	_init_.Initialize()

	if err := validateNewGuarddutyFilterFindingCriteriaCriterionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.guarddutyFilter.GuarddutyFilterFindingCriteriaCriterionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGuarddutyFilterFindingCriteriaCriterionOutputReference_Override(g GuarddutyFilterFindingCriteriaCriterionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.guarddutyFilter.GuarddutyFilterFindingCriteriaCriterionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetEqualTo(val *[]*string) {
	if err := j.validateSetEqualToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"equalTo",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetField(val *string) {
	if err := j.validateSetFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"field",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetGreaterThan(val *string) {
	if err := j.validateSetGreaterThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"greaterThan",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetGreaterThanOrEqual(val *string) {
	if err := j.validateSetGreaterThanOrEqualParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"greaterThanOrEqual",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetLessThan(val *string) {
	if err := j.validateSetLessThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lessThan",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetLessThanOrEqual(val *string) {
	if err := j.validateSetLessThanOrEqualParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lessThanOrEqual",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetNotEquals(val *[]*string) {
	if err := j.validateSetNotEqualsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notEquals",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ResetEqualTo() {
	_jsii_.InvokeVoid(
		g,
		"resetEqualTo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ResetGreaterThan() {
	_jsii_.InvokeVoid(
		g,
		"resetGreaterThan",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ResetGreaterThanOrEqual() {
	_jsii_.InvokeVoid(
		g,
		"resetGreaterThanOrEqual",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ResetLessThan() {
	_jsii_.InvokeVoid(
		g,
		"resetLessThan",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ResetLessThanOrEqual() {
	_jsii_.InvokeVoid(
		g,
		"resetLessThanOrEqual",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ResetNotEquals() {
	_jsii_.InvokeVoid(
		g,
		"resetNotEquals",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuarddutyFilterFindingCriteriaCriterionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

