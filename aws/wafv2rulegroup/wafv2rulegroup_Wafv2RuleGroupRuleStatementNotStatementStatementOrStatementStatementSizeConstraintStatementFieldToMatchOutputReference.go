package wafv2rulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/wafv2rulegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference interface {
	cdktf.ComplexObject
	AllQueryArguments() Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchAllQueryArgumentsOutputReference
	AllQueryArgumentsInput() *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments
	Body() Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchBodyOutputReference
	BodyInput() *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchBody
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
	Cookies() Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchCookiesOutputReference
	CookiesInput() *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchCookies
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Headers() Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList
	HeadersInput() interface{}
	InternalValue() *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatch
	SetInternalValue(val *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatch)
	JsonBody() Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchJsonBodyOutputReference
	JsonBodyInput() *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchJsonBody
	Method() Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchMethodOutputReference
	MethodInput() *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchMethod
	QueryString() Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchQueryStringOutputReference
	QueryStringInput() *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchQueryString
	SingleHeader() Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleHeaderOutputReference
	SingleHeaderInput() *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleHeader
	SingleQueryArgument() Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgumentOutputReference
	SingleQueryArgumentInput() *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UriPath() Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchUriPathOutputReference
	UriPathInput() *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchUriPath
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
	PutAllQueryArguments(value *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments)
	PutBody(value *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchBody)
	PutCookies(value *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchCookies)
	PutHeaders(value interface{})
	PutJsonBody(value *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchJsonBody)
	PutMethod(value *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchMethod)
	PutQueryString(value *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchQueryString)
	PutSingleHeader(value *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleHeader)
	PutSingleQueryArgument(value *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument)
	PutUriPath(value *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchUriPath)
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

// The jsii proxy struct for Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference
type jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) AllQueryArguments() Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchAllQueryArgumentsOutputReference {
	var returns Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchAllQueryArgumentsOutputReference
	_jsii_.Get(
		j,
		"allQueryArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) AllQueryArgumentsInput() *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments {
	var returns *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments
	_jsii_.Get(
		j,
		"allQueryArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Body() Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchBodyOutputReference {
	var returns Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchBodyOutputReference
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) BodyInput() *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchBody {
	var returns *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchBody
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Cookies() Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchCookiesOutputReference {
	var returns Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchCookiesOutputReference
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) CookiesInput() *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchCookies {
	var returns *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchCookies
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Headers() Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList {
	var returns Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) InternalValue() *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatch {
	var returns *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) JsonBody() Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchJsonBodyOutputReference {
	var returns Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchJsonBodyOutputReference
	_jsii_.Get(
		j,
		"jsonBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) JsonBodyInput() *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchJsonBody {
	var returns *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchJsonBody
	_jsii_.Get(
		j,
		"jsonBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Method() Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchMethodOutputReference {
	var returns Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchMethodOutputReference
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) MethodInput() *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchMethod {
	var returns *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchMethod
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) QueryString() Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchQueryStringOutputReference {
	var returns Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchQueryStringOutputReference
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) QueryStringInput() *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchQueryString {
	var returns *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchQueryString
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) SingleHeader() Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleHeaderOutputReference {
	var returns Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleHeaderOutputReference
	_jsii_.Get(
		j,
		"singleHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) SingleHeaderInput() *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleHeader {
	var returns *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleHeader
	_jsii_.Get(
		j,
		"singleHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) SingleQueryArgument() Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgumentOutputReference {
	var returns Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgumentOutputReference
	_jsii_.Get(
		j,
		"singleQueryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) SingleQueryArgumentInput() *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument {
	var returns *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument
	_jsii_.Get(
		j,
		"singleQueryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) UriPath() Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchUriPathOutputReference {
	var returns Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchUriPathOutputReference
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) UriPathInput() *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchUriPath {
	var returns *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchUriPath
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


func NewWafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference_Override(w Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetInternalValue(val *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutAllQueryArguments(value *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments) {
	if err := w.validatePutAllQueryArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAllQueryArguments",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutBody(value *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchBody) {
	if err := w.validatePutBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutCookies(value *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchCookies) {
	if err := w.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCookies",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutHeaders(value interface{}) {
	if err := w.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putHeaders",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutJsonBody(value *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchJsonBody) {
	if err := w.validatePutJsonBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putJsonBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutMethod(value *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchMethod) {
	if err := w.validatePutMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMethod",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutQueryString(value *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchQueryString) {
	if err := w.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putQueryString",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutSingleHeader(value *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleHeader) {
	if err := w.validatePutSingleHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleHeader",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutSingleQueryArgument(value *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument) {
	if err := w.validatePutSingleQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleQueryArgument",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutUriPath(value *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchUriPath) {
	if err := w.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putUriPath",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetAllQueryArguments() {
	_jsii_.InvokeVoid(
		w,
		"resetAllQueryArguments",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		w,
		"resetBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetCookies() {
	_jsii_.InvokeVoid(
		w,
		"resetCookies",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		w,
		"resetHeaders",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetJsonBody() {
	_jsii_.InvokeVoid(
		w,
		"resetJsonBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		w,
		"resetMethod",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		w,
		"resetQueryString",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetSingleHeader() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleHeader",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetSingleQueryArgument() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleQueryArgument",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		w,
		"resetUriPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

