// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2rulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/wafv2rulegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference interface {
	cdktf.ComplexObject
	AllQueryArguments() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArgumentsOutputReference
	AllQueryArgumentsInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArguments
	Body() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBodyOutputReference
	BodyInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBody
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
	Cookies() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference
	CookiesInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookies
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Headers() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchHeadersList
	HeadersInput() interface{}
	InternalValue() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatch
	SetInternalValue(val *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatch)
	JsonBody() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyOutputReference
	JsonBodyInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchJsonBody
	Method() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethodOutputReference
	MethodInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethod
	QueryString() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryStringOutputReference
	QueryStringInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryString
	SingleHeader() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeaderOutputReference
	SingleHeaderInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeader
	SingleQueryArgument() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgumentOutputReference
	SingleQueryArgumentInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgument
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UriPath() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPathOutputReference
	UriPathInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPath
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
	PutAllQueryArguments(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArguments)
	PutBody(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBody)
	PutCookies(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookies)
	PutHeaders(value interface{})
	PutJsonBody(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchJsonBody)
	PutMethod(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethod)
	PutQueryString(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryString)
	PutSingleHeader(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeader)
	PutSingleQueryArgument(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgument)
	PutUriPath(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPath)
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

// The jsii proxy struct for Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference
type jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) AllQueryArguments() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArgumentsOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArgumentsOutputReference
	_jsii_.Get(
		j,
		"allQueryArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) AllQueryArgumentsInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArguments {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArguments
	_jsii_.Get(
		j,
		"allQueryArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) Body() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBodyOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBodyOutputReference
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) BodyInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBody {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBody
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) Cookies() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) CookiesInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookies {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookies
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) Headers() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchHeadersList {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) InternalValue() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatch {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) JsonBody() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyOutputReference
	_jsii_.Get(
		j,
		"jsonBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) JsonBodyInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchJsonBody {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchJsonBody
	_jsii_.Get(
		j,
		"jsonBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) Method() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethodOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethodOutputReference
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) MethodInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethod {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethod
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) QueryString() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryStringOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryStringOutputReference
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) QueryStringInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryString {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryString
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) SingleHeader() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeaderOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeaderOutputReference
	_jsii_.Get(
		j,
		"singleHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) SingleHeaderInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeader {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeader
	_jsii_.Get(
		j,
		"singleHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) SingleQueryArgument() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgumentOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgumentOutputReference
	_jsii_.Get(
		j,
		"singleQueryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) SingleQueryArgumentInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgument {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgument
	_jsii_.Get(
		j,
		"singleQueryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) UriPath() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPathOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPathOutputReference
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) UriPathInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPath {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPath
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


func NewWafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference_Override(w Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference)SetInternalValue(val *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutAllQueryArguments(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArguments) {
	if err := w.validatePutAllQueryArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAllQueryArguments",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutBody(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBody) {
	if err := w.validatePutBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutCookies(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookies) {
	if err := w.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCookies",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutHeaders(value interface{}) {
	if err := w.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putHeaders",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutJsonBody(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchJsonBody) {
	if err := w.validatePutJsonBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putJsonBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutMethod(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethod) {
	if err := w.validatePutMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMethod",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutQueryString(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryString) {
	if err := w.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putQueryString",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutSingleHeader(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeader) {
	if err := w.validatePutSingleHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleHeader",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutSingleQueryArgument(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgument) {
	if err := w.validatePutSingleQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleQueryArgument",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutUriPath(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPath) {
	if err := w.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putUriPath",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetAllQueryArguments() {
	_jsii_.InvokeVoid(
		w,
		"resetAllQueryArguments",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		w,
		"resetBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetCookies() {
	_jsii_.InvokeVoid(
		w,
		"resetCookies",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		w,
		"resetHeaders",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetJsonBody() {
	_jsii_.InvokeVoid(
		w,
		"resetJsonBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		w,
		"resetMethod",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		w,
		"resetQueryString",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetSingleHeader() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleHeader",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetSingleQueryArgument() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleQueryArgument",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		w,
		"resetUriPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

