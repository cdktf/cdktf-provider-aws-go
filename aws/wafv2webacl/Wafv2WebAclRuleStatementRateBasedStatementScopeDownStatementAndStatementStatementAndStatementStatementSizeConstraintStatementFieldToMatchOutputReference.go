// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/wafv2webacl/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference interface {
	cdktf.ComplexObject
	AllQueryArguments() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchAllQueryArgumentsOutputReference
	AllQueryArgumentsInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments
	Body() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchBodyOutputReference
	BodyInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchBody
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
	Cookies() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchCookiesOutputReference
	CookiesInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchCookies
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Headers() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchHeadersList
	HeadersInput() interface{}
	InternalValue() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatch
	SetInternalValue(val *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatch)
	JsonBody() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchJsonBodyOutputReference
	JsonBodyInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchJsonBody
	Method() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchMethodOutputReference
	MethodInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchMethod
	QueryString() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchQueryStringOutputReference
	QueryStringInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchQueryString
	SingleHeader() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleHeaderOutputReference
	SingleHeaderInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleHeader
	SingleQueryArgument() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgumentOutputReference
	SingleQueryArgumentInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UriPath() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchUriPathOutputReference
	UriPathInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchUriPath
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
	PutAllQueryArguments(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments)
	PutBody(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchBody)
	PutCookies(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchCookies)
	PutHeaders(value interface{})
	PutJsonBody(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchJsonBody)
	PutMethod(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchMethod)
	PutQueryString(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchQueryString)
	PutSingleHeader(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleHeader)
	PutSingleQueryArgument(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument)
	PutUriPath(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchUriPath)
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

// The jsii proxy struct for Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference
type jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) AllQueryArguments() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchAllQueryArgumentsOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchAllQueryArgumentsOutputReference
	_jsii_.Get(
		j,
		"allQueryArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) AllQueryArgumentsInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments
	_jsii_.Get(
		j,
		"allQueryArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Body() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchBodyOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchBodyOutputReference
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) BodyInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchBody {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchBody
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Cookies() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchCookiesOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchCookiesOutputReference
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) CookiesInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchCookies {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchCookies
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Headers() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchHeadersList {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) InternalValue() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatch {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) JsonBody() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchJsonBodyOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchJsonBodyOutputReference
	_jsii_.Get(
		j,
		"jsonBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) JsonBodyInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchJsonBody {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchJsonBody
	_jsii_.Get(
		j,
		"jsonBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Method() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchMethodOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchMethodOutputReference
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) MethodInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchMethod {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchMethod
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) QueryString() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchQueryStringOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchQueryStringOutputReference
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) QueryStringInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchQueryString {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchQueryString
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) SingleHeader() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleHeaderOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleHeaderOutputReference
	_jsii_.Get(
		j,
		"singleHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) SingleHeaderInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleHeader {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleHeader
	_jsii_.Get(
		j,
		"singleHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) SingleQueryArgument() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgumentOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgumentOutputReference
	_jsii_.Get(
		j,
		"singleQueryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) SingleQueryArgumentInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument
	_jsii_.Get(
		j,
		"singleQueryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) UriPath() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchUriPathOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchUriPathOutputReference
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) UriPathInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchUriPath {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchUriPath
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


func NewWafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference_Override(w Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetInternalValue(val *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutAllQueryArguments(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments) {
	if err := w.validatePutAllQueryArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAllQueryArguments",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutBody(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchBody) {
	if err := w.validatePutBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutCookies(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchCookies) {
	if err := w.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCookies",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutHeaders(value interface{}) {
	if err := w.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putHeaders",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutJsonBody(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchJsonBody) {
	if err := w.validatePutJsonBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putJsonBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutMethod(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchMethod) {
	if err := w.validatePutMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMethod",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutQueryString(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchQueryString) {
	if err := w.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putQueryString",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutSingleHeader(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleHeader) {
	if err := w.validatePutSingleHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleHeader",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutSingleQueryArgument(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument) {
	if err := w.validatePutSingleQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleQueryArgument",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutUriPath(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchUriPath) {
	if err := w.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putUriPath",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetAllQueryArguments() {
	_jsii_.InvokeVoid(
		w,
		"resetAllQueryArguments",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		w,
		"resetBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetCookies() {
	_jsii_.InvokeVoid(
		w,
		"resetCookies",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		w,
		"resetHeaders",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetJsonBody() {
	_jsii_.InvokeVoid(
		w,
		"resetJsonBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		w,
		"resetMethod",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		w,
		"resetQueryString",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetSingleHeader() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleHeader",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetSingleQueryArgument() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleQueryArgument",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		w,
		"resetUriPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

