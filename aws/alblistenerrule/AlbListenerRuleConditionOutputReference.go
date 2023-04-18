package alblistenerrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v14/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v14/alblistenerrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbListenerRuleConditionOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	HostHeader() AlbListenerRuleConditionHostHeaderOutputReference
	HostHeaderInput() *AlbListenerRuleConditionHostHeader
	HttpHeader() AlbListenerRuleConditionHttpHeaderOutputReference
	HttpHeaderInput() *AlbListenerRuleConditionHttpHeader
	HttpRequestMethod() AlbListenerRuleConditionHttpRequestMethodOutputReference
	HttpRequestMethodInput() *AlbListenerRuleConditionHttpRequestMethod
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PathPattern() AlbListenerRuleConditionPathPatternOutputReference
	PathPatternInput() *AlbListenerRuleConditionPathPattern
	QueryString() AlbListenerRuleConditionQueryStringList
	QueryStringInput() interface{}
	SourceIp() AlbListenerRuleConditionSourceIpOutputReference
	SourceIpInput() *AlbListenerRuleConditionSourceIp
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutHostHeader(value *AlbListenerRuleConditionHostHeader)
	PutHttpHeader(value *AlbListenerRuleConditionHttpHeader)
	PutHttpRequestMethod(value *AlbListenerRuleConditionHttpRequestMethod)
	PutPathPattern(value *AlbListenerRuleConditionPathPattern)
	PutQueryString(value interface{})
	PutSourceIp(value *AlbListenerRuleConditionSourceIp)
	ResetHostHeader()
	ResetHttpHeader()
	ResetHttpRequestMethod()
	ResetPathPattern()
	ResetQueryString()
	ResetSourceIp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlbListenerRuleConditionOutputReference
type jsiiProxy_AlbListenerRuleConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbListenerRuleConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleConditionOutputReference) HostHeader() AlbListenerRuleConditionHostHeaderOutputReference {
	var returns AlbListenerRuleConditionHostHeaderOutputReference
	_jsii_.Get(
		j,
		"hostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleConditionOutputReference) HostHeaderInput() *AlbListenerRuleConditionHostHeader {
	var returns *AlbListenerRuleConditionHostHeader
	_jsii_.Get(
		j,
		"hostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleConditionOutputReference) HttpHeader() AlbListenerRuleConditionHttpHeaderOutputReference {
	var returns AlbListenerRuleConditionHttpHeaderOutputReference
	_jsii_.Get(
		j,
		"httpHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleConditionOutputReference) HttpHeaderInput() *AlbListenerRuleConditionHttpHeader {
	var returns *AlbListenerRuleConditionHttpHeader
	_jsii_.Get(
		j,
		"httpHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleConditionOutputReference) HttpRequestMethod() AlbListenerRuleConditionHttpRequestMethodOutputReference {
	var returns AlbListenerRuleConditionHttpRequestMethodOutputReference
	_jsii_.Get(
		j,
		"httpRequestMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleConditionOutputReference) HttpRequestMethodInput() *AlbListenerRuleConditionHttpRequestMethod {
	var returns *AlbListenerRuleConditionHttpRequestMethod
	_jsii_.Get(
		j,
		"httpRequestMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleConditionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleConditionOutputReference) PathPattern() AlbListenerRuleConditionPathPatternOutputReference {
	var returns AlbListenerRuleConditionPathPatternOutputReference
	_jsii_.Get(
		j,
		"pathPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleConditionOutputReference) PathPatternInput() *AlbListenerRuleConditionPathPattern {
	var returns *AlbListenerRuleConditionPathPattern
	_jsii_.Get(
		j,
		"pathPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleConditionOutputReference) QueryString() AlbListenerRuleConditionQueryStringList {
	var returns AlbListenerRuleConditionQueryStringList
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleConditionOutputReference) QueryStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleConditionOutputReference) SourceIp() AlbListenerRuleConditionSourceIpOutputReference {
	var returns AlbListenerRuleConditionSourceIpOutputReference
	_jsii_.Get(
		j,
		"sourceIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleConditionOutputReference) SourceIpInput() *AlbListenerRuleConditionSourceIp {
	var returns *AlbListenerRuleConditionSourceIp
	_jsii_.Get(
		j,
		"sourceIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlbListenerRuleConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AlbListenerRuleConditionOutputReference {
	_init_.Initialize()

	if err := validateNewAlbListenerRuleConditionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlbListenerRuleConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.albListenerRule.AlbListenerRuleConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAlbListenerRuleConditionOutputReference_Override(a AlbListenerRuleConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.albListenerRule.AlbListenerRuleConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AlbListenerRuleConditionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbListenerRuleConditionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbListenerRuleConditionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbListenerRuleConditionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbListenerRuleConditionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) PutHostHeader(value *AlbListenerRuleConditionHostHeader) {
	if err := a.validatePutHostHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHostHeader",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) PutHttpHeader(value *AlbListenerRuleConditionHttpHeader) {
	if err := a.validatePutHttpHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttpHeader",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) PutHttpRequestMethod(value *AlbListenerRuleConditionHttpRequestMethod) {
	if err := a.validatePutHttpRequestMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttpRequestMethod",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) PutPathPattern(value *AlbListenerRuleConditionPathPattern) {
	if err := a.validatePutPathPatternParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPathPattern",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) PutQueryString(value interface{}) {
	if err := a.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putQueryString",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) PutSourceIp(value *AlbListenerRuleConditionSourceIp) {
	if err := a.validatePutSourceIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSourceIp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) ResetHostHeader() {
	_jsii_.InvokeVoid(
		a,
		"resetHostHeader",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) ResetHttpHeader() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpHeader",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) ResetHttpRequestMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpRequestMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) ResetPathPattern() {
	_jsii_.InvokeVoid(
		a,
		"resetPathPattern",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		a,
		"resetQueryString",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) ResetSourceIp() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceIp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListenerRuleConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

