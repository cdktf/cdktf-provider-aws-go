package wafv2webacl

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/wafv2webacl/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference interface {
	cdktf.ComplexObject
	AllQueryArguments() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArgumentsOutputReference
	AllQueryArgumentsInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArguments
	Body() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBodyOutputReference
	BodyInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBody
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
	Cookies() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference
	CookiesInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookies
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Headers() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchHeadersList
	HeadersInput() interface{}
	InternalValue() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatch
	SetInternalValue(val *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatch)
	JsonBody() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyOutputReference
	JsonBodyInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchJsonBody
	Method() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethodOutputReference
	MethodInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethod
	QueryString() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryStringOutputReference
	QueryStringInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryString
	SingleHeader() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeaderOutputReference
	SingleHeaderInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeader
	SingleQueryArgument() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgumentOutputReference
	SingleQueryArgumentInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgument
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UriPath() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPathOutputReference
	UriPathInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPath
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
	PutAllQueryArguments(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArguments)
	PutBody(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBody)
	PutCookies(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookies)
	PutHeaders(value interface{})
	PutJsonBody(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchJsonBody)
	PutMethod(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethod)
	PutQueryString(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryString)
	PutSingleHeader(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeader)
	PutSingleQueryArgument(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgument)
	PutUriPath(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPath)
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

// The jsii proxy struct for Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference
type jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) AllQueryArguments() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArgumentsOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArgumentsOutputReference
	_jsii_.Get(
		j,
		"allQueryArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) AllQueryArgumentsInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArguments {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArguments
	_jsii_.Get(
		j,
		"allQueryArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) Body() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBodyOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBodyOutputReference
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) BodyInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBody {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBody
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) Cookies() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesOutputReference
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) CookiesInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookies {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookies
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) Headers() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchHeadersList {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) InternalValue() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatch {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) JsonBody() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyOutputReference
	_jsii_.Get(
		j,
		"jsonBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) JsonBodyInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchJsonBody {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchJsonBody
	_jsii_.Get(
		j,
		"jsonBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) Method() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethodOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethodOutputReference
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) MethodInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethod {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethod
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) QueryString() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryStringOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryStringOutputReference
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) QueryStringInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryString {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryString
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) SingleHeader() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeaderOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeaderOutputReference
	_jsii_.Get(
		j,
		"singleHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) SingleHeaderInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeader {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeader
	_jsii_.Get(
		j,
		"singleHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) SingleQueryArgument() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgumentOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgumentOutputReference
	_jsii_.Get(
		j,
		"singleQueryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) SingleQueryArgumentInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgument {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgument
	_jsii_.Get(
		j,
		"singleQueryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) UriPath() Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPathOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPathOutputReference
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) UriPathInput() *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPath {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPath
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


func NewWafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference_Override(w Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference)SetInternalValue(val *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutAllQueryArguments(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArguments) {
	if err := w.validatePutAllQueryArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAllQueryArguments",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutBody(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchBody) {
	if err := w.validatePutBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutCookies(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookies) {
	if err := w.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCookies",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutHeaders(value interface{}) {
	if err := w.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putHeaders",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutJsonBody(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchJsonBody) {
	if err := w.validatePutJsonBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putJsonBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutMethod(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchMethod) {
	if err := w.validatePutMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMethod",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutQueryString(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchQueryString) {
	if err := w.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putQueryString",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutSingleHeader(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeader) {
	if err := w.validatePutSingleHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleHeader",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutSingleQueryArgument(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgument) {
	if err := w.validatePutSingleQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleQueryArgument",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) PutUriPath(value *Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchUriPath) {
	if err := w.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putUriPath",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetAllQueryArguments() {
	_jsii_.InvokeVoid(
		w,
		"resetAllQueryArguments",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		w,
		"resetBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetCookies() {
	_jsii_.InvokeVoid(
		w,
		"resetCookies",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		w,
		"resetHeaders",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetJsonBody() {
	_jsii_.InvokeVoid(
		w,
		"resetJsonBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		w,
		"resetMethod",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		w,
		"resetQueryString",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetSingleHeader() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleHeader",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetSingleQueryArgument() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleQueryArgument",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		w,
		"resetUriPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementAndStatementStatementRegexPatternSetReferenceStatementFieldToMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

