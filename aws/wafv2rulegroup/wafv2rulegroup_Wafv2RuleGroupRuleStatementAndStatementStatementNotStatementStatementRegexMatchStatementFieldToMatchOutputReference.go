package wafv2rulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/wafv2rulegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference interface {
	cdktf.ComplexObject
	AllQueryArguments() Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchAllQueryArgumentsOutputReference
	AllQueryArgumentsInput() *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchAllQueryArguments
	Body() Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchBodyOutputReference
	BodyInput() *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchBody
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
	Cookies() Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchCookiesOutputReference
	CookiesInput() *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchCookies
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Headers() Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchHeadersList
	HeadersInput() interface{}
	InternalValue() *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatch
	SetInternalValue(val *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatch)
	JsonBody() Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchJsonBodyOutputReference
	JsonBodyInput() *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchJsonBody
	Method() Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchMethodOutputReference
	MethodInput() *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchMethod
	QueryString() Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchQueryStringOutputReference
	QueryStringInput() *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchQueryString
	SingleHeader() Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchSingleHeaderOutputReference
	SingleHeaderInput() *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchSingleHeader
	SingleQueryArgument() Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchSingleQueryArgumentOutputReference
	SingleQueryArgumentInput() *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchSingleQueryArgument
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UriPath() Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchUriPathOutputReference
	UriPathInput() *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchUriPath
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
	PutAllQueryArguments(value *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchAllQueryArguments)
	PutBody(value *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchBody)
	PutCookies(value *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchCookies)
	PutHeaders(value interface{})
	PutJsonBody(value *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchJsonBody)
	PutMethod(value *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchMethod)
	PutQueryString(value *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchQueryString)
	PutSingleHeader(value *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchSingleHeader)
	PutSingleQueryArgument(value *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchSingleQueryArgument)
	PutUriPath(value *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchUriPath)
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

// The jsii proxy struct for Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference
type jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) AllQueryArguments() Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchAllQueryArgumentsOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchAllQueryArgumentsOutputReference
	_jsii_.Get(
		j,
		"allQueryArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) AllQueryArgumentsInput() *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchAllQueryArguments {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchAllQueryArguments
	_jsii_.Get(
		j,
		"allQueryArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) Body() Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchBodyOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchBodyOutputReference
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) BodyInput() *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchBody {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchBody
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) Cookies() Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchCookiesOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchCookiesOutputReference
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) CookiesInput() *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchCookies {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchCookies
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) Headers() Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchHeadersList {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) InternalValue() *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatch {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) JsonBody() Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchJsonBodyOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchJsonBodyOutputReference
	_jsii_.Get(
		j,
		"jsonBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) JsonBodyInput() *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchJsonBody {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchJsonBody
	_jsii_.Get(
		j,
		"jsonBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) Method() Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchMethodOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchMethodOutputReference
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) MethodInput() *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchMethod {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchMethod
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) QueryString() Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchQueryStringOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchQueryStringOutputReference
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) QueryStringInput() *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchQueryString {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchQueryString
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) SingleHeader() Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchSingleHeaderOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchSingleHeaderOutputReference
	_jsii_.Get(
		j,
		"singleHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) SingleHeaderInput() *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchSingleHeader {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchSingleHeader
	_jsii_.Get(
		j,
		"singleHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) SingleQueryArgument() Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchSingleQueryArgumentOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchSingleQueryArgumentOutputReference
	_jsii_.Get(
		j,
		"singleQueryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) SingleQueryArgumentInput() *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchSingleQueryArgument {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchSingleQueryArgument
	_jsii_.Get(
		j,
		"singleQueryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) UriPath() Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchUriPathOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchUriPathOutputReference
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) UriPathInput() *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchUriPath {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchUriPath
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


func NewWafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference_Override(w Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference)SetInternalValue(val *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) PutAllQueryArguments(value *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchAllQueryArguments) {
	if err := w.validatePutAllQueryArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAllQueryArguments",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) PutBody(value *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchBody) {
	if err := w.validatePutBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) PutCookies(value *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchCookies) {
	if err := w.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCookies",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) PutHeaders(value interface{}) {
	if err := w.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putHeaders",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) PutJsonBody(value *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchJsonBody) {
	if err := w.validatePutJsonBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putJsonBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) PutMethod(value *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchMethod) {
	if err := w.validatePutMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMethod",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) PutQueryString(value *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchQueryString) {
	if err := w.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putQueryString",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) PutSingleHeader(value *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchSingleHeader) {
	if err := w.validatePutSingleHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleHeader",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) PutSingleQueryArgument(value *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchSingleQueryArgument) {
	if err := w.validatePutSingleQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleQueryArgument",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) PutUriPath(value *Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchUriPath) {
	if err := w.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putUriPath",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) ResetAllQueryArguments() {
	_jsii_.InvokeVoid(
		w,
		"resetAllQueryArguments",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		w,
		"resetBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) ResetCookies() {
	_jsii_.InvokeVoid(
		w,
		"resetCookies",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		w,
		"resetHeaders",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) ResetJsonBody() {
	_jsii_.InvokeVoid(
		w,
		"resetJsonBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		w,
		"resetMethod",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		w,
		"resetQueryString",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) ResetSingleHeader() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleHeader",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) ResetSingleQueryArgument() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleQueryArgument",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		w,
		"resetUriPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementNotStatementStatementRegexMatchStatementFieldToMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

