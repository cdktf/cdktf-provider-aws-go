package wafv2rulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/wafv2rulegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList interface {
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
	Get(index *float64) Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList
type jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewWafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList {
	_init_.Initialize()

	if err := validateNewWafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewWafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList_Override(w Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		w,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList) Get(index *float64) Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternOutputReference {
	if err := w.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternOutputReference

	_jsii_.Invoke(
		w,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

