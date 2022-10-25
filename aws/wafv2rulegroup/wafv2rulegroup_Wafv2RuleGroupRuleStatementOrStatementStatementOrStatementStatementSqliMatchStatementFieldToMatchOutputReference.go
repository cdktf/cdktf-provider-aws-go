package wafv2rulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/wafv2rulegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference interface {
	cdktf.ComplexObject
	AllQueryArguments() Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchAllQueryArgumentsOutputReference
	AllQueryArgumentsInput() *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchAllQueryArguments
	Body() Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchBodyOutputReference
	BodyInput() *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchBody
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
	Cookies() Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchCookiesOutputReference
	CookiesInput() *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchCookies
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Headers() Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchHeadersList
	HeadersInput() interface{}
	InternalValue() *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatch
	SetInternalValue(val *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatch)
	JsonBody() Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchJsonBodyOutputReference
	JsonBodyInput() *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchJsonBody
	Method() Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchMethodOutputReference
	MethodInput() *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchMethod
	QueryString() Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchQueryStringOutputReference
	QueryStringInput() *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchQueryString
	SingleHeader() Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchSingleHeaderOutputReference
	SingleHeaderInput() *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchSingleHeader
	SingleQueryArgument() Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchSingleQueryArgumentOutputReference
	SingleQueryArgumentInput() *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchSingleQueryArgument
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UriPath() Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchUriPathOutputReference
	UriPathInput() *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchUriPath
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
	PutAllQueryArguments(value *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchAllQueryArguments)
	PutBody(value *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchBody)
	PutCookies(value *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchCookies)
	PutHeaders(value interface{})
	PutJsonBody(value *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchJsonBody)
	PutMethod(value *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchMethod)
	PutQueryString(value *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchQueryString)
	PutSingleHeader(value *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchSingleHeader)
	PutSingleQueryArgument(value *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchSingleQueryArgument)
	PutUriPath(value *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchUriPath)
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

// The jsii proxy struct for Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference
type jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) AllQueryArguments() Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchAllQueryArgumentsOutputReference {
	var returns Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchAllQueryArgumentsOutputReference
	_jsii_.Get(
		j,
		"allQueryArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) AllQueryArgumentsInput() *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchAllQueryArguments {
	var returns *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchAllQueryArguments
	_jsii_.Get(
		j,
		"allQueryArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) Body() Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchBodyOutputReference {
	var returns Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchBodyOutputReference
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) BodyInput() *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchBody {
	var returns *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchBody
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) Cookies() Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchCookiesOutputReference {
	var returns Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchCookiesOutputReference
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) CookiesInput() *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchCookies {
	var returns *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchCookies
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) Headers() Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchHeadersList {
	var returns Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) InternalValue() *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatch {
	var returns *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) JsonBody() Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchJsonBodyOutputReference {
	var returns Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchJsonBodyOutputReference
	_jsii_.Get(
		j,
		"jsonBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) JsonBodyInput() *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchJsonBody {
	var returns *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchJsonBody
	_jsii_.Get(
		j,
		"jsonBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) Method() Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchMethodOutputReference {
	var returns Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchMethodOutputReference
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) MethodInput() *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchMethod {
	var returns *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchMethod
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) QueryString() Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchQueryStringOutputReference {
	var returns Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchQueryStringOutputReference
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) QueryStringInput() *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchQueryString {
	var returns *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchQueryString
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) SingleHeader() Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchSingleHeaderOutputReference {
	var returns Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchSingleHeaderOutputReference
	_jsii_.Get(
		j,
		"singleHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) SingleHeaderInput() *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchSingleHeader {
	var returns *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchSingleHeader
	_jsii_.Get(
		j,
		"singleHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) SingleQueryArgument() Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchSingleQueryArgumentOutputReference {
	var returns Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchSingleQueryArgumentOutputReference
	_jsii_.Get(
		j,
		"singleQueryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) SingleQueryArgumentInput() *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchSingleQueryArgument {
	var returns *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchSingleQueryArgument
	_jsii_.Get(
		j,
		"singleQueryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) UriPath() Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchUriPathOutputReference {
	var returns Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchUriPathOutputReference
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) UriPathInput() *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchUriPath {
	var returns *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchUriPath
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


func NewWafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference_Override(w Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference)SetInternalValue(val *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) PutAllQueryArguments(value *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchAllQueryArguments) {
	if err := w.validatePutAllQueryArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAllQueryArguments",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) PutBody(value *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchBody) {
	if err := w.validatePutBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) PutCookies(value *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchCookies) {
	if err := w.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCookies",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) PutHeaders(value interface{}) {
	if err := w.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putHeaders",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) PutJsonBody(value *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchJsonBody) {
	if err := w.validatePutJsonBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putJsonBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) PutMethod(value *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchMethod) {
	if err := w.validatePutMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMethod",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) PutQueryString(value *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchQueryString) {
	if err := w.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putQueryString",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) PutSingleHeader(value *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchSingleHeader) {
	if err := w.validatePutSingleHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleHeader",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) PutSingleQueryArgument(value *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchSingleQueryArgument) {
	if err := w.validatePutSingleQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleQueryArgument",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) PutUriPath(value *Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchUriPath) {
	if err := w.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putUriPath",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) ResetAllQueryArguments() {
	_jsii_.InvokeVoid(
		w,
		"resetAllQueryArguments",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		w,
		"resetBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) ResetCookies() {
	_jsii_.InvokeVoid(
		w,
		"resetCookies",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		w,
		"resetHeaders",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) ResetJsonBody() {
	_jsii_.InvokeVoid(
		w,
		"resetJsonBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		w,
		"resetMethod",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		w,
		"resetQueryString",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) ResetSingleHeader() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleHeader",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) ResetSingleQueryArgument() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleQueryArgument",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		w,
		"resetUriPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementOrStatementStatementOrStatementStatementSqliMatchStatementFieldToMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

