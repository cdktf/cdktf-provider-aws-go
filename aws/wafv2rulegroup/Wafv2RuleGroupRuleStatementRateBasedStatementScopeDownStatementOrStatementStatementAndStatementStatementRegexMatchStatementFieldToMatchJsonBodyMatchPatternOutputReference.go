// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2rulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/wafv2rulegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference interface {
	cdktf.ComplexObject
	All() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternAllOutputReference
	AllInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternAll
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
	IncludedPaths() *[]*string
	SetIncludedPaths(val *[]*string)
	IncludedPathsInput() *[]*string
	InternalValue() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPattern
	SetInternalValue(val *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPattern)
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
	PutAll(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternAll)
	ResetAll()
	ResetIncludedPaths()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference
type jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) All() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternAllOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternAllOutputReference
	_jsii_.Get(
		j,
		"all",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) AllInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternAll {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternAll
	_jsii_.Get(
		j,
		"allInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) IncludedPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) IncludedPathsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) InternalValue() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPattern {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPattern
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference_Override(w Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference)SetIncludedPaths(val *[]*string) {
	if err := j.validateSetIncludedPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedPaths",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference)SetInternalValue(val *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPattern) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) PutAll(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternAll) {
	if err := w.validatePutAllParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAll",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) ResetAll() {
	_jsii_.InvokeVoid(
		w,
		"resetAll",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) ResetIncludedPaths() {
	_jsii_.InvokeVoid(
		w,
		"resetIncludedPaths",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

