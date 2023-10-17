// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/wafv2webacl/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference interface {
	cdktf.ComplexObject
	AllQueryArguments() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchAllQueryArgumentsOutputReference
	AllQueryArgumentsInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchAllQueryArguments
	Body() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchBodyOutputReference
	BodyInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchBody
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
	Cookies() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchCookiesOutputReference
	CookiesInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchCookies
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Headers() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchHeadersList
	HeadersInput() interface{}
	InternalValue() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatch
	SetInternalValue(val *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatch)
	JsonBody() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchJsonBodyOutputReference
	JsonBodyInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchJsonBody
	Method() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchMethodOutputReference
	MethodInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchMethod
	QueryString() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchQueryStringOutputReference
	QueryStringInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchQueryString
	SingleHeader() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchSingleHeaderOutputReference
	SingleHeaderInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchSingleHeader
	SingleQueryArgument() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchSingleQueryArgumentOutputReference
	SingleQueryArgumentInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchSingleQueryArgument
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UriPath() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchUriPathOutputReference
	UriPathInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchUriPath
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
	PutAllQueryArguments(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchAllQueryArguments)
	PutBody(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchBody)
	PutCookies(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchCookies)
	PutHeaders(value interface{})
	PutJsonBody(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchJsonBody)
	PutMethod(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchMethod)
	PutQueryString(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchQueryString)
	PutSingleHeader(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchSingleHeader)
	PutSingleQueryArgument(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchSingleQueryArgument)
	PutUriPath(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchUriPath)
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

// The jsii proxy struct for Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference
type jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) AllQueryArguments() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchAllQueryArgumentsOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchAllQueryArgumentsOutputReference
	_jsii_.Get(
		j,
		"allQueryArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) AllQueryArgumentsInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchAllQueryArguments {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchAllQueryArguments
	_jsii_.Get(
		j,
		"allQueryArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) Body() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchBodyOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchBodyOutputReference
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) BodyInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchBody {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchBody
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) Cookies() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchCookiesOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchCookiesOutputReference
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) CookiesInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchCookies {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchCookies
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) Headers() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchHeadersList {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) InternalValue() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatch {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) JsonBody() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchJsonBodyOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchJsonBodyOutputReference
	_jsii_.Get(
		j,
		"jsonBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) JsonBodyInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchJsonBody {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchJsonBody
	_jsii_.Get(
		j,
		"jsonBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) Method() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchMethodOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchMethodOutputReference
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) MethodInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchMethod {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchMethod
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) QueryString() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchQueryStringOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchQueryStringOutputReference
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) QueryStringInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchQueryString {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchQueryString
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) SingleHeader() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchSingleHeaderOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchSingleHeaderOutputReference
	_jsii_.Get(
		j,
		"singleHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) SingleHeaderInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchSingleHeader {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchSingleHeader
	_jsii_.Get(
		j,
		"singleHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) SingleQueryArgument() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchSingleQueryArgumentOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchSingleQueryArgumentOutputReference
	_jsii_.Get(
		j,
		"singleQueryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) SingleQueryArgumentInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchSingleQueryArgument {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchSingleQueryArgument
	_jsii_.Get(
		j,
		"singleQueryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) UriPath() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchUriPathOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchUriPathOutputReference
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) UriPathInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchUriPath {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchUriPath
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


func NewWafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference_Override(w Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference)SetInternalValue(val *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) PutAllQueryArguments(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchAllQueryArguments) {
	if err := w.validatePutAllQueryArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAllQueryArguments",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) PutBody(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchBody) {
	if err := w.validatePutBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) PutCookies(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchCookies) {
	if err := w.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCookies",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) PutHeaders(value interface{}) {
	if err := w.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putHeaders",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) PutJsonBody(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchJsonBody) {
	if err := w.validatePutJsonBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putJsonBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) PutMethod(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchMethod) {
	if err := w.validatePutMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMethod",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) PutQueryString(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchQueryString) {
	if err := w.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putQueryString",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) PutSingleHeader(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchSingleHeader) {
	if err := w.validatePutSingleHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleHeader",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) PutSingleQueryArgument(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchSingleQueryArgument) {
	if err := w.validatePutSingleQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleQueryArgument",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) PutUriPath(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchUriPath) {
	if err := w.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putUriPath",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) ResetAllQueryArguments() {
	_jsii_.InvokeVoid(
		w,
		"resetAllQueryArguments",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		w,
		"resetBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) ResetCookies() {
	_jsii_.InvokeVoid(
		w,
		"resetCookies",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		w,
		"resetHeaders",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) ResetJsonBody() {
	_jsii_.InvokeVoid(
		w,
		"resetJsonBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		w,
		"resetMethod",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		w,
		"resetQueryString",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) ResetSingleHeader() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleHeader",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) ResetSingleQueryArgument() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleQueryArgument",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		w,
		"resetUriPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementXssMatchStatementFieldToMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

