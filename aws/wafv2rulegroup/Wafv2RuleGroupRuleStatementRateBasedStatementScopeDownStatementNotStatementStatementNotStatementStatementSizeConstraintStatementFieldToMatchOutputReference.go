// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2rulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/wafv2rulegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference interface {
	cdktf.ComplexObject
	AllQueryArguments() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchAllQueryArgumentsOutputReference
	AllQueryArgumentsInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments
	Body() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchBodyOutputReference
	BodyInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchBody
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
	Cookies() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchCookiesOutputReference
	CookiesInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchCookies
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Headers() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersList
	HeadersInput() interface{}
	InternalValue() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatch
	SetInternalValue(val *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatch)
	JsonBody() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchJsonBodyOutputReference
	JsonBodyInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchJsonBody
	Method() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchMethodOutputReference
	MethodInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchMethod
	QueryString() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchQueryStringOutputReference
	QueryStringInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchQueryString
	SingleHeader() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchSingleHeaderOutputReference
	SingleHeaderInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchSingleHeader
	SingleQueryArgument() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgumentOutputReference
	SingleQueryArgumentInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UriPath() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchUriPathOutputReference
	UriPathInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchUriPath
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
	PutAllQueryArguments(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments)
	PutBody(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchBody)
	PutCookies(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchCookies)
	PutHeaders(value interface{})
	PutJsonBody(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchJsonBody)
	PutMethod(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchMethod)
	PutQueryString(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchQueryString)
	PutSingleHeader(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchSingleHeader)
	PutSingleQueryArgument(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument)
	PutUriPath(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchUriPath)
	ResetAllQueryArguments()
	ResetBody()
	ResetCookies()
	ResetHeaders()
	ResetJsonBody()
	ResetMethod()
	ResetQueryString()
	ResetSingleHeader()
	ResetSingleQueryArgument()
	ResetUriPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference
type jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) AllQueryArguments() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchAllQueryArgumentsOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchAllQueryArgumentsOutputReference
	_jsii_.Get(
		j,
		"allQueryArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) AllQueryArgumentsInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments
	_jsii_.Get(
		j,
		"allQueryArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Body() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchBodyOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchBodyOutputReference
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) BodyInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchBody {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchBody
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Cookies() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchCookiesOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchCookiesOutputReference
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) CookiesInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchCookies {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchCookies
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Headers() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersList {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) InternalValue() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatch {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) JsonBody() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchJsonBodyOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchJsonBodyOutputReference
	_jsii_.Get(
		j,
		"jsonBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) JsonBodyInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchJsonBody {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchJsonBody
	_jsii_.Get(
		j,
		"jsonBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Method() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchMethodOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchMethodOutputReference
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) MethodInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchMethod {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchMethod
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) QueryString() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchQueryStringOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchQueryStringOutputReference
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) QueryStringInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchQueryString {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchQueryString
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) SingleHeader() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchSingleHeaderOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchSingleHeaderOutputReference
	_jsii_.Get(
		j,
		"singleHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) SingleHeaderInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchSingleHeader {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchSingleHeader
	_jsii_.Get(
		j,
		"singleHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) SingleQueryArgument() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgumentOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgumentOutputReference
	_jsii_.Get(
		j,
		"singleQueryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) SingleQueryArgumentInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument
	_jsii_.Get(
		j,
		"singleQueryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) UriPath() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchUriPathOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchUriPathOutputReference
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) UriPathInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchUriPath {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchUriPath
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


func NewWafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference_Override(w Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetInternalValue(val *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutAllQueryArguments(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments) {
	if err := w.validatePutAllQueryArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAllQueryArguments",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutBody(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchBody) {
	if err := w.validatePutBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutCookies(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchCookies) {
	if err := w.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCookies",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutHeaders(value interface{}) {
	if err := w.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putHeaders",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutJsonBody(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchJsonBody) {
	if err := w.validatePutJsonBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putJsonBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutMethod(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchMethod) {
	if err := w.validatePutMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMethod",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutQueryString(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchQueryString) {
	if err := w.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putQueryString",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutSingleHeader(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchSingleHeader) {
	if err := w.validatePutSingleHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleHeader",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutSingleQueryArgument(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument) {
	if err := w.validatePutSingleQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleQueryArgument",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutUriPath(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchUriPath) {
	if err := w.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putUriPath",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetAllQueryArguments() {
	_jsii_.InvokeVoid(
		w,
		"resetAllQueryArguments",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		w,
		"resetBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetCookies() {
	_jsii_.InvokeVoid(
		w,
		"resetCookies",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		w,
		"resetHeaders",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetJsonBody() {
	_jsii_.InvokeVoid(
		w,
		"resetJsonBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		w,
		"resetMethod",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		w,
		"resetQueryString",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetSingleHeader() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleHeader",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetSingleQueryArgument() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleQueryArgument",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		w,
		"resetUriPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

