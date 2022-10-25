package wafv2rulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/wafv2rulegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference interface {
	cdktf.ComplexObject
	AllQueryArguments() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchAllQueryArgumentsOutputReference
	AllQueryArgumentsInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments
	Body() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchBodyOutputReference
	BodyInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchBody
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
	Cookies() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchCookiesOutputReference
	CookiesInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchCookies
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Headers() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchHeadersList
	HeadersInput() interface{}
	InternalValue() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatch
	SetInternalValue(val *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatch)
	JsonBody() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchJsonBodyOutputReference
	JsonBodyInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchJsonBody
	Method() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchMethodOutputReference
	MethodInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchMethod
	QueryString() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchQueryStringOutputReference
	QueryStringInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchQueryString
	SingleHeader() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleHeaderOutputReference
	SingleHeaderInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleHeader
	SingleQueryArgument() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgumentOutputReference
	SingleQueryArgumentInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UriPath() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchUriPathOutputReference
	UriPathInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchUriPath
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
	PutAllQueryArguments(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments)
	PutBody(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchBody)
	PutCookies(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchCookies)
	PutHeaders(value interface{})
	PutJsonBody(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchJsonBody)
	PutMethod(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchMethod)
	PutQueryString(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchQueryString)
	PutSingleHeader(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleHeader)
	PutSingleQueryArgument(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument)
	PutUriPath(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchUriPath)
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

// The jsii proxy struct for Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference
type jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) AllQueryArguments() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchAllQueryArgumentsOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchAllQueryArgumentsOutputReference
	_jsii_.Get(
		j,
		"allQueryArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) AllQueryArgumentsInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments
	_jsii_.Get(
		j,
		"allQueryArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Body() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchBodyOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchBodyOutputReference
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) BodyInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchBody {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchBody
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Cookies() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchCookiesOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchCookiesOutputReference
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) CookiesInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchCookies {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchCookies
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Headers() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchHeadersList {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) InternalValue() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatch {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) JsonBody() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchJsonBodyOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchJsonBodyOutputReference
	_jsii_.Get(
		j,
		"jsonBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) JsonBodyInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchJsonBody {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchJsonBody
	_jsii_.Get(
		j,
		"jsonBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Method() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchMethodOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchMethodOutputReference
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) MethodInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchMethod {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchMethod
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) QueryString() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchQueryStringOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchQueryStringOutputReference
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) QueryStringInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchQueryString {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchQueryString
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) SingleHeader() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleHeaderOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleHeaderOutputReference
	_jsii_.Get(
		j,
		"singleHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) SingleHeaderInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleHeader {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleHeader
	_jsii_.Get(
		j,
		"singleHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) SingleQueryArgument() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgumentOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgumentOutputReference
	_jsii_.Get(
		j,
		"singleQueryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) SingleQueryArgumentInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument
	_jsii_.Get(
		j,
		"singleQueryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) UriPath() Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchUriPathOutputReference {
	var returns Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchUriPathOutputReference
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) UriPathInput() *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchUriPath {
	var returns *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchUriPath
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


func NewWafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference_Override(w Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetInternalValue(val *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutAllQueryArguments(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments) {
	if err := w.validatePutAllQueryArgumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAllQueryArguments",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutBody(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchBody) {
	if err := w.validatePutBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutCookies(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchCookies) {
	if err := w.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCookies",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutHeaders(value interface{}) {
	if err := w.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putHeaders",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutJsonBody(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchJsonBody) {
	if err := w.validatePutJsonBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putJsonBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutMethod(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchMethod) {
	if err := w.validatePutMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMethod",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutQueryString(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchQueryString) {
	if err := w.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putQueryString",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutSingleHeader(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleHeader) {
	if err := w.validatePutSingleHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleHeader",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutSingleQueryArgument(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument) {
	if err := w.validatePutSingleQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleQueryArgument",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) PutUriPath(value *Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchUriPath) {
	if err := w.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putUriPath",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetAllQueryArguments() {
	_jsii_.InvokeVoid(
		w,
		"resetAllQueryArguments",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		w,
		"resetBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetCookies() {
	_jsii_.InvokeVoid(
		w,
		"resetCookies",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		w,
		"resetHeaders",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetJsonBody() {
	_jsii_.InvokeVoid(
		w,
		"resetJsonBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		w,
		"resetMethod",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		w,
		"resetQueryString",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetSingleHeader() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleHeader",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetSingleQueryArgument() {
	_jsii_.InvokeVoid(
		w,
		"resetSingleQueryArgument",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		w,
		"resetUriPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementAndStatementStatementAndStatementStatementSizeConstraintStatementFieldToMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

