// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2rulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/wafv2rulegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference interface {
	cdktf.ComplexObject
	All() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternAllOutputReference
	AllInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternAll
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
	ExcludedCookies() *[]*string
	SetExcludedCookies(val *[]*string)
	ExcludedCookiesInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludedCookies() *[]*string
	SetIncludedCookies(val *[]*string)
	IncludedCookiesInput() *[]*string
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
	PutAll(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternAll)
	ResetAll()
	ResetExcludedCookies()
	ResetIncludedCookies()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference
type jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) All() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternAllOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternAllOutputReference
	_jsii_.Get(
		j,
		"all",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) AllInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternAll {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternAll
	_jsii_.Get(
		j,
		"allInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) ExcludedCookies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedCookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) ExcludedCookiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedCookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) IncludedCookies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedCookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) IncludedCookiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedCookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference_Override(w Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference)SetExcludedCookies(val *[]*string) {
	if err := j.validateSetExcludedCookiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedCookies",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference)SetIncludedCookies(val *[]*string) {
	if err := j.validateSetIncludedCookiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedCookies",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) PutAll(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternAll) {
	if err := w.validatePutAllParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAll",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) ResetAll() {
	_jsii_.InvokeVoid(
		w,
		"resetAll",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) ResetExcludedCookies() {
	_jsii_.InvokeVoid(
		w,
		"resetExcludedCookies",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) ResetIncludedCookies() {
	_jsii_.InvokeVoid(
		w,
		"resetIncludedCookies",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatementStatementXssMatchStatementFieldToMatchCookiesMatchPatternOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

