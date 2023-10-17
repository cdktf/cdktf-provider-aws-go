// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/wafv2webacl/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference interface {
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
	InternalValue() *Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookies
	SetInternalValue(val *Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookies)
	MatchPattern() Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList
	MatchPatternInput() interface{}
	MatchScope() *string
	SetMatchScope(val *string)
	MatchScopeInput() *string
	OversizeHandling() *string
	SetOversizeHandling(val *string)
	OversizeHandlingInput() *string
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
	PutMatchPattern(value interface{})
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference
type jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) InternalValue() *Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookies {
	var returns *Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookies
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) MatchPattern() Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList {
	var returns Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList
	_jsii_.Get(
		j,
		"matchPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) MatchPatternInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) MatchScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) MatchScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) OversizeHandling() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oversizeHandling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) OversizeHandlingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oversizeHandlingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference_Override(w Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference)SetInternalValue(val *Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookies) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference)SetMatchScope(val *string) {
	if err := j.validateSetMatchScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchScope",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference)SetOversizeHandling(val *string) {
	if err := j.validateSetOversizeHandlingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oversizeHandling",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) PutMatchPattern(value interface{}) {
	if err := w.validatePutMatchPatternParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMatchPattern",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

