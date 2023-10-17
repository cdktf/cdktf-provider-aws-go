// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2rulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/wafv2rulegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList
type jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewWafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList {
	_init_.Initialize()

	if err := validateNewWafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewWafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList_Override(w Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		w,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList) Get(index *float64) Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersOutputReference {
	if err := w.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersOutputReference

	_jsii_.Invoke(
		w,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

