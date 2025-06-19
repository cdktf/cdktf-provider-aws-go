// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package macie2findingsfilter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/macie2findingsfilter/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Macie2FindingsFilterFindingCriteriaCriterionOutputReference interface {
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
	Eq() *[]*string
	SetEq(val *[]*string)
	EqExactMatch() *[]*string
	SetEqExactMatch(val *[]*string)
	EqExactMatchInput() *[]*string
	EqInput() *[]*string
	Field() *string
	SetField(val *string)
	FieldInput() *string
	// Experimental.
	Fqn() *string
	Gt() *string
	SetGt(val *string)
	Gte() *string
	SetGte(val *string)
	GteInput() *string
	GtInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Lt() *string
	SetLt(val *string)
	Lte() *string
	SetLte(val *string)
	LteInput() *string
	LtInput() *string
	Neq() *[]*string
	SetNeq(val *[]*string)
	NeqInput() *[]*string
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
	ResetEq()
	ResetEqExactMatch()
	ResetGt()
	ResetGte()
	ResetLt()
	ResetLte()
	ResetNeq()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Macie2FindingsFilterFindingCriteriaCriterionOutputReference
type jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) Eq() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) EqExactMatch() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eqExactMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) EqExactMatchInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eqExactMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) EqInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eqInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) Field() *string {
	var returns *string
	_jsii_.Get(
		j,
		"field",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) FieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) Gt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) Gte() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) GteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) GtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) Lt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) Lte() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) LteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) LtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ltInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) Neq() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"neq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) NeqInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"neqInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMacie2FindingsFilterFindingCriteriaCriterionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Macie2FindingsFilterFindingCriteriaCriterionOutputReference {
	_init_.Initialize()

	if err := validateNewMacie2FindingsFilterFindingCriteriaCriterionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.macie2FindingsFilter.Macie2FindingsFilterFindingCriteriaCriterionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMacie2FindingsFilterFindingCriteriaCriterionOutputReference_Override(m Macie2FindingsFilterFindingCriteriaCriterionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.macie2FindingsFilter.Macie2FindingsFilterFindingCriteriaCriterionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference)SetEq(val *[]*string) {
	if err := j.validateSetEqParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eq",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference)SetEqExactMatch(val *[]*string) {
	if err := j.validateSetEqExactMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eqExactMatch",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference)SetField(val *string) {
	if err := j.validateSetFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"field",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference)SetGt(val *string) {
	if err := j.validateSetGtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gt",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference)SetGte(val *string) {
	if err := j.validateSetGteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gte",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference)SetLt(val *string) {
	if err := j.validateSetLtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lt",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference)SetLte(val *string) {
	if err := j.validateSetLteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lte",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference)SetNeq(val *[]*string) {
	if err := j.validateSetNeqParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"neq",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) ResetEq() {
	_jsii_.InvokeVoid(
		m,
		"resetEq",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) ResetEqExactMatch() {
	_jsii_.InvokeVoid(
		m,
		"resetEqExactMatch",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) ResetGt() {
	_jsii_.InvokeVoid(
		m,
		"resetGt",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) ResetGte() {
	_jsii_.InvokeVoid(
		m,
		"resetGte",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) ResetLt() {
	_jsii_.InvokeVoid(
		m,
		"resetLt",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) ResetLte() {
	_jsii_.InvokeVoid(
		m,
		"resetLte",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) ResetNeq() {
	_jsii_.InvokeVoid(
		m,
		"resetNeq",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Macie2FindingsFilterFindingCriteriaCriterionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

