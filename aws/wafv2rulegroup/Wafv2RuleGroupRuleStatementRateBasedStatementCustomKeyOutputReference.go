// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2rulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v17/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v17/wafv2rulegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference interface {
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
	Cookie() Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyCookieOutputReference
	CookieInput() *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyCookie
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ForwardedIp() Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyForwardedIpOutputReference
	ForwardedIpInput() *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyForwardedIp
	// Experimental.
	Fqn() *string
	Header() Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyHeaderOutputReference
	HeaderInput() *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyHeader
	HttpMethod() Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyHttpMethodOutputReference
	HttpMethodInput() *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyHttpMethod
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ip() Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyIpOutputReference
	IpInput() *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyIp
	LabelNamespace() Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyLabelNamespaceOutputReference
	LabelNamespaceInput() *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyLabelNamespace
	QueryArgument() Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyQueryArgumentOutputReference
	QueryArgumentInput() *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyQueryArgument
	QueryString() Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyQueryStringOutputReference
	QueryStringInput() *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyQueryString
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UriPath() Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyUriPathOutputReference
	UriPathInput() *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyUriPath
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
	PutCookie(value *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyCookie)
	PutForwardedIp(value *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyForwardedIp)
	PutHeader(value *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyHeader)
	PutHttpMethod(value *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyHttpMethod)
	PutIp(value *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyIp)
	PutLabelNamespace(value *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyLabelNamespace)
	PutQueryArgument(value *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyQueryArgument)
	PutQueryString(value *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyQueryString)
	PutUriPath(value *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyUriPath)
	ResetCookie()
	ResetForwardedIp()
	ResetHeader()
	ResetHttpMethod()
	ResetIp()
	ResetLabelNamespace()
	ResetQueryArgument()
	ResetQueryString()
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

// The jsii proxy struct for Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference
type jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) Cookie() Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyCookieOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyCookieOutputReference
	_jsii_.Get(
		j,
		"cookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) CookieInput() *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyCookie {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyCookie
	_jsii_.Get(
		j,
		"cookieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) ForwardedIp() Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyForwardedIpOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyForwardedIpOutputReference
	_jsii_.Get(
		j,
		"forwardedIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) ForwardedIpInput() *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyForwardedIp {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyForwardedIp
	_jsii_.Get(
		j,
		"forwardedIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) Header() Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyHeaderOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyHeaderOutputReference
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) HeaderInput() *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyHeader {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyHeader
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) HttpMethod() Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyHttpMethodOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyHttpMethodOutputReference
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) HttpMethodInput() *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyHttpMethod {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyHttpMethod
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) Ip() Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyIpOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyIpOutputReference
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) IpInput() *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyIp {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyIp
	_jsii_.Get(
		j,
		"ipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) LabelNamespace() Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyLabelNamespaceOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyLabelNamespaceOutputReference
	_jsii_.Get(
		j,
		"labelNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) LabelNamespaceInput() *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyLabelNamespace {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyLabelNamespace
	_jsii_.Get(
		j,
		"labelNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) QueryArgument() Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyQueryArgumentOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyQueryArgumentOutputReference
	_jsii_.Get(
		j,
		"queryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) QueryArgumentInput() *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyQueryArgument {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyQueryArgument
	_jsii_.Get(
		j,
		"queryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) QueryString() Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyQueryStringOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyQueryStringOutputReference
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) QueryStringInput() *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyQueryString {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyQueryString
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) UriPath() Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyUriPathOutputReference {
	var returns Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyUriPathOutputReference
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) UriPathInput() *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyUriPath {
	var returns *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyUriPath
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


func NewWafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference_Override(w Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) PutCookie(value *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyCookie) {
	if err := w.validatePutCookieParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCookie",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) PutForwardedIp(value *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyForwardedIp) {
	if err := w.validatePutForwardedIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putForwardedIp",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) PutHeader(value *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyHeader) {
	if err := w.validatePutHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putHeader",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) PutHttpMethod(value *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyHttpMethod) {
	if err := w.validatePutHttpMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putHttpMethod",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) PutIp(value *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyIp) {
	if err := w.validatePutIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putIp",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) PutLabelNamespace(value *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyLabelNamespace) {
	if err := w.validatePutLabelNamespaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putLabelNamespace",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) PutQueryArgument(value *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyQueryArgument) {
	if err := w.validatePutQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putQueryArgument",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) PutQueryString(value *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyQueryString) {
	if err := w.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putQueryString",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) PutUriPath(value *Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyUriPath) {
	if err := w.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putUriPath",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) ResetCookie() {
	_jsii_.InvokeVoid(
		w,
		"resetCookie",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) ResetForwardedIp() {
	_jsii_.InvokeVoid(
		w,
		"resetForwardedIp",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		w,
		"resetHeader",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) ResetHttpMethod() {
	_jsii_.InvokeVoid(
		w,
		"resetHttpMethod",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) ResetIp() {
	_jsii_.InvokeVoid(
		w,
		"resetIp",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) ResetLabelNamespace() {
	_jsii_.InvokeVoid(
		w,
		"resetLabelNamespace",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) ResetQueryArgument() {
	_jsii_.InvokeVoid(
		w,
		"resetQueryArgument",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		w,
		"resetQueryString",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		w,
		"resetUriPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2RuleGroupRuleStatementRateBasedStatementCustomKeyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

