package wafv2rulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/wafv2rulegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference interface {
	cdktf.ComplexObject
	AllQueryArguments() Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchAllQueryArgumentsOutputReference
	AllQueryArgumentsInput() *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments
	Body() Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchBodyOutputReference
	BodyInput() *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchBody
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
	Cookies() Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchCookiesOutputReference
	CookiesInput() *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchCookies
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Headers() Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList
	HeadersInput() interface{}
	InternalValue() *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatch
	SetInternalValue(val *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatch)
	JsonBody() Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchJsonBodyOutputReference
	JsonBodyInput() *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchJsonBody
	Method() Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchMethodOutputReference
	MethodInput() *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchMethod
	QueryString() Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchQueryStringOutputReference
	QueryStringInput() *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchQueryString
	SingleHeader() Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleHeaderOutputReference
	SingleHeaderInput() *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleHeader
	SingleQueryArgument() Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgumentOutputReference
	SingleQueryArgumentInput() *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UriPath() Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchUriPathOutputReference
	UriPathInput() *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchUriPath
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
	PutAllQueryArguments(value *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments)
	PutBody(value *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchBody)
	PutCookies(value *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchCookies)
	PutHeaders(value interface{})
	PutJsonBody(value *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchJsonBody)
	PutMethod(value *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchMethod)
	PutQueryString(value *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchQueryString)
	PutSingleHeader(value *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleHeader)
	PutSingleQueryArgument(value *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument)
	PutUriPath(value *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchUriPath)
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

// The jsii proxy struct for Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference
type jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) AllQueryArguments() Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchAllQueryArgumentsOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchAllQueryArgumentsOutputReference
	_jsii_.Get(
		j,
		"allQueryArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) AllQueryArgumentsInput() *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments
	_jsii_.Get(
		j,
		"allQueryArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Body() Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchBodyOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchBodyOutputReference
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) BodyInput() *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchBody {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchBody
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Cookies() Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchCookiesOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchCookiesOutputReference
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) CookiesInput() *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchCookies {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchCookies
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Headers() Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) InternalValue() *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatch {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) JsonBody() Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchJsonBodyOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchJsonBodyOutputReference
	_jsii_.Get(
		j,
		"jsonBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) JsonBodyInput() *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchJsonBody {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchJsonBody
	_jsii_.Get(
		j,
		"jsonBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Method() Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchMethodOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchMethodOutputReference
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) MethodInput() *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchMethod {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchMethod
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) QueryString() Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchQueryStringOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchQueryStringOutputReference
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) QueryStringInput() *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchQueryString {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchQueryString
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) SingleHeader() Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleHeaderOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleHeaderOutputReference
	_jsii_.Get(
		j,
		"singleHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) SingleHeaderInput() *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleHeader {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleHeader
	_jsii_.Get(
		j,
		"singleHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) SingleQueryArgument() Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgumentOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgumentOutputReference
	_jsii_.Get(
		j,
		"singleQueryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) SingleQueryArgumentInput() *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument
	_jsii_.Get(
		j,
		"singleQueryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) UriPath() Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchUriPathOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchUriPathOutputReference
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) UriPathInput() *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchUriPath {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchUriPath
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


func NewWafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference_Override(w Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetInternalValue(val *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutAllQueryArguments(value *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments) {
	if err := w.validatePutAllQueryArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAllQueryArguments",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutBody(value *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchBody) {
	if err := w.validatePutBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutCookies(value *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchCookies) {
	if err := w.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCookies",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutHeaders(value interface{}) {
	if err := w.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putHeaders",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutJsonBody(value *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchJsonBody) {
	if err := w.validatePutJsonBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putJsonBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutMethod(value *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchMethod) {
	if err := w.validatePutMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMethod",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutQueryString(value *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchQueryString) {
	if err := w.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putQueryString",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutSingleHeader(value *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleHeader) {
	if err := w.validatePutSingleHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleHeader",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutSingleQueryArgument(value *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument) {
	if err := w.validatePutSingleQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleQueryArgument",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutUriPath(value *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchUriPath) {
	if err := w.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putUriPath",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetAllQueryArguments() {
	_jsii_.InvokeVoid(
		w,
		"resetAllQueryArguments",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		w,
		"resetBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetCookies() {
	_jsii_.InvokeVoid(
		w,
		"resetCookies",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		w,
		"resetHeaders",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetJsonBody() {
	_jsii_.InvokeVoid(
		w,
		"resetJsonBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		w,
		"resetMethod",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		w,
		"resetQueryString",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetSingleHeader() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleHeader",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetSingleQueryArgument() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleQueryArgument",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		w,
		"resetUriPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

