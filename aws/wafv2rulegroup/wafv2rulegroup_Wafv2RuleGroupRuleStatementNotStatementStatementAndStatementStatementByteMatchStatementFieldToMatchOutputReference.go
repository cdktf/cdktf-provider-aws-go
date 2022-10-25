package wafv2rulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/wafv2rulegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference interface {
	cdktf.ComplexObject
	AllQueryArguments() Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchAllQueryArgumentsOutputReference
	AllQueryArgumentsInput() *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchAllQueryArguments
	Body() Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchBodyOutputReference
	BodyInput() *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchBody
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
	Cookies() Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchCookiesOutputReference
	CookiesInput() *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchCookies
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Headers() Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchHeadersList
	HeadersInput() interface{}
	InternalValue() *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatch
	SetInternalValue(val *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatch)
	JsonBody() Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchJsonBodyOutputReference
	JsonBodyInput() *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchJsonBody
	Method() Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchMethodOutputReference
	MethodInput() *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchMethod
	QueryString() Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchQueryStringOutputReference
	QueryStringInput() *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchQueryString
	SingleHeader() Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchSingleHeaderOutputReference
	SingleHeaderInput() *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchSingleHeader
	SingleQueryArgument() Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchSingleQueryArgumentOutputReference
	SingleQueryArgumentInput() *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchSingleQueryArgument
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UriPath() Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchUriPathOutputReference
	UriPathInput() *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchUriPath
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
	PutAllQueryArguments(value *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchAllQueryArguments)
	PutBody(value *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchBody)
	PutCookies(value *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchCookies)
	PutHeaders(value interface{})
	PutJsonBody(value *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchJsonBody)
	PutMethod(value *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchMethod)
	PutQueryString(value *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchQueryString)
	PutSingleHeader(value *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchSingleHeader)
	PutSingleQueryArgument(value *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchSingleQueryArgument)
	PutUriPath(value *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchUriPath)
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

// The jsii proxy struct for Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference
type jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) AllQueryArguments() Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchAllQueryArgumentsOutputReference {
	var returns Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchAllQueryArgumentsOutputReference
	_jsii_.Get(
		j,
		"allQueryArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) AllQueryArgumentsInput() *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchAllQueryArguments {
	var returns *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchAllQueryArguments
	_jsii_.Get(
		j,
		"allQueryArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) Body() Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchBodyOutputReference {
	var returns Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchBodyOutputReference
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) BodyInput() *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchBody {
	var returns *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchBody
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) Cookies() Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchCookiesOutputReference {
	var returns Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchCookiesOutputReference
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) CookiesInput() *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchCookies {
	var returns *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchCookies
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) Headers() Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchHeadersList {
	var returns Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) InternalValue() *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatch {
	var returns *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) JsonBody() Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchJsonBodyOutputReference {
	var returns Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchJsonBodyOutputReference
	_jsii_.Get(
		j,
		"jsonBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) JsonBodyInput() *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchJsonBody {
	var returns *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchJsonBody
	_jsii_.Get(
		j,
		"jsonBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) Method() Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchMethodOutputReference {
	var returns Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchMethodOutputReference
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) MethodInput() *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchMethod {
	var returns *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchMethod
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) QueryString() Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchQueryStringOutputReference {
	var returns Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchQueryStringOutputReference
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) QueryStringInput() *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchQueryString {
	var returns *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchQueryString
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) SingleHeader() Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchSingleHeaderOutputReference {
	var returns Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchSingleHeaderOutputReference
	_jsii_.Get(
		j,
		"singleHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) SingleHeaderInput() *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchSingleHeader {
	var returns *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchSingleHeader
	_jsii_.Get(
		j,
		"singleHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) SingleQueryArgument() Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchSingleQueryArgumentOutputReference {
	var returns Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchSingleQueryArgumentOutputReference
	_jsii_.Get(
		j,
		"singleQueryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) SingleQueryArgumentInput() *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchSingleQueryArgument {
	var returns *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchSingleQueryArgument
	_jsii_.Get(
		j,
		"singleQueryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) UriPath() Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchUriPathOutputReference {
	var returns Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchUriPathOutputReference
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) UriPathInput() *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchUriPath {
	var returns *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchUriPath
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


func NewWafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference_Override(w Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference)SetInternalValue(val *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) PutAllQueryArguments(value *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchAllQueryArguments) {
	if err := w.validatePutAllQueryArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAllQueryArguments",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) PutBody(value *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchBody) {
	if err := w.validatePutBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) PutCookies(value *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchCookies) {
	if err := w.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCookies",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) PutHeaders(value interface{}) {
	if err := w.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putHeaders",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) PutJsonBody(value *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchJsonBody) {
	if err := w.validatePutJsonBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putJsonBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) PutMethod(value *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchMethod) {
	if err := w.validatePutMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMethod",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) PutQueryString(value *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchQueryString) {
	if err := w.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putQueryString",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) PutSingleHeader(value *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchSingleHeader) {
	if err := w.validatePutSingleHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleHeader",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) PutSingleQueryArgument(value *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchSingleQueryArgument) {
	if err := w.validatePutSingleQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleQueryArgument",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) PutUriPath(value *Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchUriPath) {
	if err := w.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putUriPath",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) ResetAllQueryArguments() {
	_jsii_.InvokeVoid(
		w,
		"resetAllQueryArguments",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		w,
		"resetBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) ResetCookies() {
	_jsii_.InvokeVoid(
		w,
		"resetCookies",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		w,
		"resetHeaders",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) ResetJsonBody() {
	_jsii_.InvokeVoid(
		w,
		"resetJsonBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		w,
		"resetMethod",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		w,
		"resetQueryString",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) ResetSingleHeader() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleHeader",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) ResetSingleQueryArgument() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleQueryArgument",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		w,
		"resetUriPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementAndStatementStatementByteMatchStatementFieldToMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

