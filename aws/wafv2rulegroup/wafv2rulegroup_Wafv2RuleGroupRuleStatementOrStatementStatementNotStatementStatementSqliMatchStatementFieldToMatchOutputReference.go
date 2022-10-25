package wafv2rulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/wafv2rulegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference interface {
	cdktf.ComplexObject
	AllQueryArguments() Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchAllQueryArgumentsOutputReference
	AllQueryArgumentsInput() *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchAllQueryArguments
	Body() Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchBodyOutputReference
	BodyInput() *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchBody
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
	Cookies() Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchCookiesOutputReference
	CookiesInput() *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchCookies
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Headers() Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchHeadersList
	HeadersInput() interface{}
	InternalValue() *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatch
	SetInternalValue(val *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatch)
	JsonBody() Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchJsonBodyOutputReference
	JsonBodyInput() *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchJsonBody
	Method() Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchMethodOutputReference
	MethodInput() *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchMethod
	QueryString() Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchQueryStringOutputReference
	QueryStringInput() *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchQueryString
	SingleHeader() Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchSingleHeaderOutputReference
	SingleHeaderInput() *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchSingleHeader
	SingleQueryArgument() Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchSingleQueryArgumentOutputReference
	SingleQueryArgumentInput() *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchSingleQueryArgument
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UriPath() Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchUriPathOutputReference
	UriPathInput() *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchUriPath
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
	PutAllQueryArguments(value *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchAllQueryArguments)
	PutBody(value *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchBody)
	PutCookies(value *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchCookies)
	PutHeaders(value interface{})
	PutJsonBody(value *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchJsonBody)
	PutMethod(value *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchMethod)
	PutQueryString(value *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchQueryString)
	PutSingleHeader(value *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchSingleHeader)
	PutSingleQueryArgument(value *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchSingleQueryArgument)
	PutUriPath(value *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchUriPath)
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

// The jsii proxy struct for Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference
type jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) AllQueryArguments() Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchAllQueryArgumentsOutputReference {
	var returns Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchAllQueryArgumentsOutputReference
	_jsii_.Get(
		j,
		"allQueryArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) AllQueryArgumentsInput() *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchAllQueryArguments {
	var returns *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchAllQueryArguments
	_jsii_.Get(
		j,
		"allQueryArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) Body() Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchBodyOutputReference {
	var returns Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchBodyOutputReference
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) BodyInput() *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchBody {
	var returns *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchBody
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) Cookies() Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchCookiesOutputReference {
	var returns Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchCookiesOutputReference
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) CookiesInput() *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchCookies {
	var returns *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchCookies
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) Headers() Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchHeadersList {
	var returns Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) InternalValue() *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatch {
	var returns *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) JsonBody() Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchJsonBodyOutputReference {
	var returns Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchJsonBodyOutputReference
	_jsii_.Get(
		j,
		"jsonBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) JsonBodyInput() *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchJsonBody {
	var returns *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchJsonBody
	_jsii_.Get(
		j,
		"jsonBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) Method() Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchMethodOutputReference {
	var returns Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchMethodOutputReference
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) MethodInput() *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchMethod {
	var returns *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchMethod
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) QueryString() Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchQueryStringOutputReference {
	var returns Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchQueryStringOutputReference
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) QueryStringInput() *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchQueryString {
	var returns *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchQueryString
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) SingleHeader() Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchSingleHeaderOutputReference {
	var returns Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchSingleHeaderOutputReference
	_jsii_.Get(
		j,
		"singleHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) SingleHeaderInput() *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchSingleHeader {
	var returns *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchSingleHeader
	_jsii_.Get(
		j,
		"singleHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) SingleQueryArgument() Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchSingleQueryArgumentOutputReference {
	var returns Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchSingleQueryArgumentOutputReference
	_jsii_.Get(
		j,
		"singleQueryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) SingleQueryArgumentInput() *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchSingleQueryArgument {
	var returns *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchSingleQueryArgument
	_jsii_.Get(
		j,
		"singleQueryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) UriPath() Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchUriPathOutputReference {
	var returns Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchUriPathOutputReference
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) UriPathInput() *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchUriPath {
	var returns *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchUriPath
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


func NewWafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference_Override(w Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference)SetInternalValue(val *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) PutAllQueryArguments(value *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchAllQueryArguments) {
	if err := w.validatePutAllQueryArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAllQueryArguments",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) PutBody(value *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchBody) {
	if err := w.validatePutBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) PutCookies(value *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchCookies) {
	if err := w.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCookies",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) PutHeaders(value interface{}) {
	if err := w.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putHeaders",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) PutJsonBody(value *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchJsonBody) {
	if err := w.validatePutJsonBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putJsonBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) PutMethod(value *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchMethod) {
	if err := w.validatePutMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMethod",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) PutQueryString(value *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchQueryString) {
	if err := w.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putQueryString",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) PutSingleHeader(value *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchSingleHeader) {
	if err := w.validatePutSingleHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleHeader",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) PutSingleQueryArgument(value *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchSingleQueryArgument) {
	if err := w.validatePutSingleQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleQueryArgument",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) PutUriPath(value *Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchUriPath) {
	if err := w.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putUriPath",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) ResetAllQueryArguments() {
	_jsii_.InvokeVoid(
		w,
		"resetAllQueryArguments",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		w,
		"resetBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) ResetCookies() {
	_jsii_.InvokeVoid(
		w,
		"resetCookies",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		w,
		"resetHeaders",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) ResetJsonBody() {
	_jsii_.InvokeVoid(
		w,
		"resetJsonBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		w,
		"resetMethod",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		w,
		"resetQueryString",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) ResetSingleHeader() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleHeader",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) ResetSingleQueryArgument() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleQueryArgument",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		w,
		"resetUriPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementNotStatementStatementSqliMatchStatementFieldToMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

