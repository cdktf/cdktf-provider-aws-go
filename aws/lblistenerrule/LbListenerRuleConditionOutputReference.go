// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lblistenerrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/lblistenerrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbListenerRuleConditionOutputReference interface {
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
	HostHeader() LbListenerRuleConditionHostHeaderOutputReference
	HostHeaderInput() *LbListenerRuleConditionHostHeader
	HttpHeader() LbListenerRuleConditionHttpHeaderOutputReference
	HttpHeaderInput() *LbListenerRuleConditionHttpHeader
	HttpRequestMethod() LbListenerRuleConditionHttpRequestMethodOutputReference
	HttpRequestMethodInput() *LbListenerRuleConditionHttpRequestMethod
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PathPattern() LbListenerRuleConditionPathPatternOutputReference
	PathPatternInput() *LbListenerRuleConditionPathPattern
	QueryString() LbListenerRuleConditionQueryStringList
	QueryStringInput() interface{}
	SourceIp() LbListenerRuleConditionSourceIpOutputReference
	SourceIpInput() *LbListenerRuleConditionSourceIp
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
	PutHostHeader(value *LbListenerRuleConditionHostHeader)
	PutHttpHeader(value *LbListenerRuleConditionHttpHeader)
	PutHttpRequestMethod(value *LbListenerRuleConditionHttpRequestMethod)
	PutPathPattern(value *LbListenerRuleConditionPathPattern)
	PutQueryString(value interface{})
	PutSourceIp(value *LbListenerRuleConditionSourceIp)
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

// The jsii proxy struct for LbListenerRuleConditionOutputReference
type jsiiProxy_LbListenerRuleConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) HostHeader() LbListenerRuleConditionHostHeaderOutputReference {
	var returns LbListenerRuleConditionHostHeaderOutputReference
	_jsii_.Get(
		j,
		"hostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) HostHeaderInput() *LbListenerRuleConditionHostHeader {
	var returns *LbListenerRuleConditionHostHeader
	_jsii_.Get(
		j,
		"hostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) HttpHeader() LbListenerRuleConditionHttpHeaderOutputReference {
	var returns LbListenerRuleConditionHttpHeaderOutputReference
	_jsii_.Get(
		j,
		"httpHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) HttpHeaderInput() *LbListenerRuleConditionHttpHeader {
	var returns *LbListenerRuleConditionHttpHeader
	_jsii_.Get(
		j,
		"httpHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) HttpRequestMethod() LbListenerRuleConditionHttpRequestMethodOutputReference {
	var returns LbListenerRuleConditionHttpRequestMethodOutputReference
	_jsii_.Get(
		j,
		"httpRequestMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) HttpRequestMethodInput() *LbListenerRuleConditionHttpRequestMethod {
	var returns *LbListenerRuleConditionHttpRequestMethod
	_jsii_.Get(
		j,
		"httpRequestMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) PathPattern() LbListenerRuleConditionPathPatternOutputReference {
	var returns LbListenerRuleConditionPathPatternOutputReference
	_jsii_.Get(
		j,
		"pathPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) PathPatternInput() *LbListenerRuleConditionPathPattern {
	var returns *LbListenerRuleConditionPathPattern
	_jsii_.Get(
		j,
		"pathPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) QueryString() LbListenerRuleConditionQueryStringList {
	var returns LbListenerRuleConditionQueryStringList
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) QueryStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) SourceIp() LbListenerRuleConditionSourceIpOutputReference {
	var returns LbListenerRuleConditionSourceIpOutputReference
	_jsii_.Get(
		j,
		"sourceIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) SourceIpInput() *LbListenerRuleConditionSourceIp {
	var returns *LbListenerRuleConditionSourceIp
	_jsii_.Get(
		j,
		"sourceIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLbListenerRuleConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LbListenerRuleConditionOutputReference {
	_init_.Initialize()

	if err := validateNewLbListenerRuleConditionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LbListenerRuleConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.lbListenerRule.LbListenerRuleConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLbListenerRuleConditionOutputReference_Override(l LbListenerRuleConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lbListenerRule.LbListenerRuleConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LbListenerRuleConditionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) PutHostHeader(value *LbListenerRuleConditionHostHeader) {
	if err := l.validatePutHostHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putHostHeader",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) PutHttpHeader(value *LbListenerRuleConditionHttpHeader) {
	if err := l.validatePutHttpHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putHttpHeader",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) PutHttpRequestMethod(value *LbListenerRuleConditionHttpRequestMethod) {
	if err := l.validatePutHttpRequestMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putHttpRequestMethod",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) PutPathPattern(value *LbListenerRuleConditionPathPattern) {
	if err := l.validatePutPathPatternParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putPathPattern",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) PutQueryString(value interface{}) {
	if err := l.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putQueryString",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) PutSourceIp(value *LbListenerRuleConditionSourceIp) {
	if err := l.validatePutSourceIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSourceIp",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) ResetHostHeader() {
	_jsii_.InvokeVoid(
		l,
		"resetHostHeader",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) ResetHttpHeader() {
	_jsii_.InvokeVoid(
		l,
		"resetHttpHeader",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) ResetHttpRequestMethod() {
	_jsii_.InvokeVoid(
		l,
		"resetHttpRequestMethod",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) ResetPathPattern() {
	_jsii_.InvokeVoid(
		l,
		"resetPathPattern",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		l,
		"resetQueryString",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) ResetSourceIp() {
	_jsii_.InvokeVoid(
		l,
		"resetSourceIp",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

