package wafv2webacl

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/wafv2webacl/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference interface {
	cdktf.ComplexObject
	AllQueryArguments() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchAllQueryArgumentsOutputReference
	AllQueryArgumentsInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchAllQueryArguments
	Body() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchBodyOutputReference
	BodyInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchBody
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
	Cookies() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesOutputReference
	CookiesInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchCookies
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Headers() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchHeadersList
	HeadersInput() interface{}
	InternalValue() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatch
	SetInternalValue(val *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatch)
	JsonBody() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyOutputReference
	JsonBodyInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBody
	Method() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchMethodOutputReference
	MethodInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchMethod
	QueryString() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchQueryStringOutputReference
	QueryStringInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchQueryString
	SingleHeader() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleHeaderOutputReference
	SingleHeaderInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleHeader
	SingleQueryArgument() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleQueryArgumentOutputReference
	SingleQueryArgumentInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleQueryArgument
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UriPath() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriPathOutputReference
	UriPathInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriPath
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
	PutAllQueryArguments(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchAllQueryArguments)
	PutBody(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchBody)
	PutCookies(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchCookies)
	PutHeaders(value interface{})
	PutJsonBody(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBody)
	PutMethod(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchMethod)
	PutQueryString(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchQueryString)
	PutSingleHeader(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleHeader)
	PutSingleQueryArgument(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleQueryArgument)
	PutUriPath(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriPath)
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

// The jsii proxy struct for Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference
type jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) AllQueryArguments() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchAllQueryArgumentsOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchAllQueryArgumentsOutputReference
	_jsii_.Get(
		j,
		"allQueryArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) AllQueryArgumentsInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchAllQueryArguments {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchAllQueryArguments
	_jsii_.Get(
		j,
		"allQueryArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) Body() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchBodyOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchBodyOutputReference
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) BodyInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchBody {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchBody
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) Cookies() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesOutputReference
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) CookiesInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchCookies {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchCookies
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) Headers() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchHeadersList {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) InternalValue() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatch {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) JsonBody() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyOutputReference
	_jsii_.Get(
		j,
		"jsonBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) JsonBodyInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBody {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBody
	_jsii_.Get(
		j,
		"jsonBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) Method() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchMethodOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchMethodOutputReference
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) MethodInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchMethod {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchMethod
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) QueryString() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchQueryStringOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchQueryStringOutputReference
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) QueryStringInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchQueryString {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchQueryString
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) SingleHeader() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleHeaderOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleHeaderOutputReference
	_jsii_.Get(
		j,
		"singleHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) SingleHeaderInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleHeader {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleHeader
	_jsii_.Get(
		j,
		"singleHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) SingleQueryArgument() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleQueryArgumentOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleQueryArgumentOutputReference
	_jsii_.Get(
		j,
		"singleQueryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) SingleQueryArgumentInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleQueryArgument {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleQueryArgument
	_jsii_.Get(
		j,
		"singleQueryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) UriPath() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriPathOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriPathOutputReference
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) UriPathInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriPath {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriPath
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


func NewWafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference_Override(w Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference)SetInternalValue(val *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) PutAllQueryArguments(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchAllQueryArguments) {
	if err := w.validatePutAllQueryArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAllQueryArguments",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) PutBody(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchBody) {
	if err := w.validatePutBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) PutCookies(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchCookies) {
	if err := w.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCookies",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) PutHeaders(value interface{}) {
	if err := w.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putHeaders",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) PutJsonBody(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBody) {
	if err := w.validatePutJsonBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putJsonBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) PutMethod(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchMethod) {
	if err := w.validatePutMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMethod",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) PutQueryString(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchQueryString) {
	if err := w.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putQueryString",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) PutSingleHeader(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleHeader) {
	if err := w.validatePutSingleHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleHeader",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) PutSingleQueryArgument(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleQueryArgument) {
	if err := w.validatePutSingleQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleQueryArgument",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) PutUriPath(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriPath) {
	if err := w.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putUriPath",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ResetAllQueryArguments() {
	_jsii_.InvokeVoid(
		w,
		"resetAllQueryArguments",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		w,
		"resetBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ResetCookies() {
	_jsii_.InvokeVoid(
		w,
		"resetCookies",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		w,
		"resetHeaders",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ResetJsonBody() {
	_jsii_.InvokeVoid(
		w,
		"resetJsonBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		w,
		"resetMethod",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		w,
		"resetQueryString",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ResetSingleHeader() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleHeader",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ResetSingleQueryArgument() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleQueryArgument",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		w,
		"resetUriPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

