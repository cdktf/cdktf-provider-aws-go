package wafv2rulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/wafv2rulegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference interface {
	cdktf.ComplexObject
	AllQueryArguments() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchAllQueryArgumentsOutputReference
	AllQueryArgumentsInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchAllQueryArguments
	Body() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchBodyOutputReference
	BodyInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchBody
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
	Cookies() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesOutputReference
	CookiesInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchCookies
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Headers() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchHeadersList
	HeadersInput() interface{}
	InternalValue() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatch
	SetInternalValue(val *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatch)
	JsonBody() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyOutputReference
	JsonBodyInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBody
	Method() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchMethodOutputReference
	MethodInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchMethod
	QueryString() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchQueryStringOutputReference
	QueryStringInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchQueryString
	SingleHeader() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleHeaderOutputReference
	SingleHeaderInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleHeader
	SingleQueryArgument() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleQueryArgumentOutputReference
	SingleQueryArgumentInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleQueryArgument
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UriPath() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriPathOutputReference
	UriPathInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriPath
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
	PutAllQueryArguments(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchAllQueryArguments)
	PutBody(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchBody)
	PutCookies(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchCookies)
	PutHeaders(value interface{})
	PutJsonBody(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBody)
	PutMethod(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchMethod)
	PutQueryString(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchQueryString)
	PutSingleHeader(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleHeader)
	PutSingleQueryArgument(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleQueryArgument)
	PutUriPath(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriPath)
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

// The jsii proxy struct for Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference
type jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) AllQueryArguments() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchAllQueryArgumentsOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchAllQueryArgumentsOutputReference
	_jsii_.Get(
		j,
		"allQueryArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) AllQueryArgumentsInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchAllQueryArguments {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchAllQueryArguments
	_jsii_.Get(
		j,
		"allQueryArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) Body() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchBodyOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchBodyOutputReference
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) BodyInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchBody {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchBody
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) Cookies() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesOutputReference
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) CookiesInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchCookies {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchCookies
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) Headers() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchHeadersList {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) InternalValue() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatch {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) JsonBody() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyOutputReference
	_jsii_.Get(
		j,
		"jsonBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) JsonBodyInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBody {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBody
	_jsii_.Get(
		j,
		"jsonBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) Method() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchMethodOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchMethodOutputReference
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) MethodInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchMethod {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchMethod
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) QueryString() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchQueryStringOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchQueryStringOutputReference
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) QueryStringInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchQueryString {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchQueryString
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) SingleHeader() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleHeaderOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleHeaderOutputReference
	_jsii_.Get(
		j,
		"singleHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) SingleHeaderInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleHeader {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleHeader
	_jsii_.Get(
		j,
		"singleHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) SingleQueryArgument() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleQueryArgumentOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleQueryArgumentOutputReference
	_jsii_.Get(
		j,
		"singleQueryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) SingleQueryArgumentInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleQueryArgument {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleQueryArgument
	_jsii_.Get(
		j,
		"singleQueryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) UriPath() Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriPathOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriPathOutputReference
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) UriPathInput() *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriPath {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriPath
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


func NewWafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference_Override(w Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference)SetInternalValue(val *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) PutAllQueryArguments(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchAllQueryArguments) {
	if err := w.validatePutAllQueryArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAllQueryArguments",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) PutBody(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchBody) {
	if err := w.validatePutBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) PutCookies(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchCookies) {
	if err := w.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCookies",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) PutHeaders(value interface{}) {
	if err := w.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putHeaders",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) PutJsonBody(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBody) {
	if err := w.validatePutJsonBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putJsonBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) PutMethod(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchMethod) {
	if err := w.validatePutMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMethod",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) PutQueryString(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchQueryString) {
	if err := w.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putQueryString",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) PutSingleHeader(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleHeader) {
	if err := w.validatePutSingleHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleHeader",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) PutSingleQueryArgument(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchSingleQueryArgument) {
	if err := w.validatePutSingleQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleQueryArgument",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) PutUriPath(value *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriPath) {
	if err := w.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putUriPath",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ResetAllQueryArguments() {
	_jsii_.InvokeVoid(
		w,
		"resetAllQueryArguments",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		w,
		"resetBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ResetCookies() {
	_jsii_.InvokeVoid(
		w,
		"resetCookies",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		w,
		"resetHeaders",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ResetJsonBody() {
	_jsii_.InvokeVoid(
		w,
		"resetJsonBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		w,
		"resetMethod",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		w,
		"resetQueryString",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ResetSingleHeader() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleHeader",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ResetSingleQueryArgument() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleQueryArgument",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		w,
		"resetUriPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

