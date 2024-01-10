// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/wafv2webacl/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference interface {
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
	Cookie() Wafv2WebAclRuleStatementRateBasedStatementCustomKeyCookieOutputReference
	CookieInput() *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyCookie
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ForwardedIp() Wafv2WebAclRuleStatementRateBasedStatementCustomKeyForwardedIpOutputReference
	ForwardedIpInput() *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyForwardedIp
	// Experimental.
	Fqn() *string
	Header() Wafv2WebAclRuleStatementRateBasedStatementCustomKeyHeaderOutputReference
	HeaderInput() *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyHeader
	HttpMethod() Wafv2WebAclRuleStatementRateBasedStatementCustomKeyHttpMethodOutputReference
	HttpMethodInput() *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyHttpMethod
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ip() Wafv2WebAclRuleStatementRateBasedStatementCustomKeyIpOutputReference
	IpInput() *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyIp
	LabelNamespace() Wafv2WebAclRuleStatementRateBasedStatementCustomKeyLabelNamespaceOutputReference
	LabelNamespaceInput() *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyLabelNamespace
	QueryArgument() Wafv2WebAclRuleStatementRateBasedStatementCustomKeyQueryArgumentOutputReference
	QueryArgumentInput() *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyQueryArgument
	QueryString() Wafv2WebAclRuleStatementRateBasedStatementCustomKeyQueryStringOutputReference
	QueryStringInput() *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyQueryString
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UriPath() Wafv2WebAclRuleStatementRateBasedStatementCustomKeyUriPathOutputReference
	UriPathInput() *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyUriPath
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
	PutCookie(value *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyCookie)
	PutForwardedIp(value *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyForwardedIp)
	PutHeader(value *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyHeader)
	PutHttpMethod(value *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyHttpMethod)
	PutIp(value *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyIp)
	PutLabelNamespace(value *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyLabelNamespace)
	PutQueryArgument(value *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyQueryArgument)
	PutQueryString(value *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyQueryString)
	PutUriPath(value *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyUriPath)
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

// The jsii proxy struct for Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference
type jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) Cookie() Wafv2WebAclRuleStatementRateBasedStatementCustomKeyCookieOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementCustomKeyCookieOutputReference
	_jsii_.Get(
		j,
		"cookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) CookieInput() *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyCookie {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyCookie
	_jsii_.Get(
		j,
		"cookieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) ForwardedIp() Wafv2WebAclRuleStatementRateBasedStatementCustomKeyForwardedIpOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementCustomKeyForwardedIpOutputReference
	_jsii_.Get(
		j,
		"forwardedIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) ForwardedIpInput() *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyForwardedIp {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyForwardedIp
	_jsii_.Get(
		j,
		"forwardedIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) Header() Wafv2WebAclRuleStatementRateBasedStatementCustomKeyHeaderOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementCustomKeyHeaderOutputReference
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) HeaderInput() *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyHeader {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyHeader
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) HttpMethod() Wafv2WebAclRuleStatementRateBasedStatementCustomKeyHttpMethodOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementCustomKeyHttpMethodOutputReference
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) HttpMethodInput() *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyHttpMethod {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyHttpMethod
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) Ip() Wafv2WebAclRuleStatementRateBasedStatementCustomKeyIpOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementCustomKeyIpOutputReference
	_jsii_.Get(
		j,
		"ip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) IpInput() *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyIp {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyIp
	_jsii_.Get(
		j,
		"ipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) LabelNamespace() Wafv2WebAclRuleStatementRateBasedStatementCustomKeyLabelNamespaceOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementCustomKeyLabelNamespaceOutputReference
	_jsii_.Get(
		j,
		"labelNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) LabelNamespaceInput() *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyLabelNamespace {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyLabelNamespace
	_jsii_.Get(
		j,
		"labelNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) QueryArgument() Wafv2WebAclRuleStatementRateBasedStatementCustomKeyQueryArgumentOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementCustomKeyQueryArgumentOutputReference
	_jsii_.Get(
		j,
		"queryArgument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) QueryArgumentInput() *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyQueryArgument {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyQueryArgument
	_jsii_.Get(
		j,
		"queryArgumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) QueryString() Wafv2WebAclRuleStatementRateBasedStatementCustomKeyQueryStringOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementCustomKeyQueryStringOutputReference
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) QueryStringInput() *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyQueryString {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyQueryString
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) UriPath() Wafv2WebAclRuleStatementRateBasedStatementCustomKeyUriPathOutputReference {
	var returns Wafv2WebAclRuleStatementRateBasedStatementCustomKeyUriPathOutputReference
	_jsii_.Get(
		j,
		"uriPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) UriPathInput() *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyUriPath {
	var returns *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyUriPath
	_jsii_.Get(
		j,
		"uriPathInput",
		&returns,
	)
	return returns
}


func NewWafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference_Override(w Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) PutCookie(value *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyCookie) {
	if err := w.validatePutCookieParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCookie",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) PutForwardedIp(value *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyForwardedIp) {
	if err := w.validatePutForwardedIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putForwardedIp",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) PutHeader(value *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyHeader) {
	if err := w.validatePutHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putHeader",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) PutHttpMethod(value *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyHttpMethod) {
	if err := w.validatePutHttpMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putHttpMethod",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) PutIp(value *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyIp) {
	if err := w.validatePutIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putIp",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) PutLabelNamespace(value *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyLabelNamespace) {
	if err := w.validatePutLabelNamespaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putLabelNamespace",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) PutQueryArgument(value *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyQueryArgument) {
	if err := w.validatePutQueryArgumentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putQueryArgument",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) PutQueryString(value *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyQueryString) {
	if err := w.validatePutQueryStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putQueryString",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) PutUriPath(value *Wafv2WebAclRuleStatementRateBasedStatementCustomKeyUriPath) {
	if err := w.validatePutUriPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putUriPath",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) ResetCookie() {
	_jsii_.InvokeVoid(
		w,
		"resetCookie",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) ResetForwardedIp() {
	_jsii_.InvokeVoid(
		w,
		"resetForwardedIp",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		w,
		"resetHeader",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) ResetHttpMethod() {
	_jsii_.InvokeVoid(
		w,
		"resetHttpMethod",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) ResetIp() {
	_jsii_.InvokeVoid(
		w,
		"resetIp",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) ResetLabelNamespace() {
	_jsii_.InvokeVoid(
		w,
		"resetLabelNamespace",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) ResetQueryArgument() {
	_jsii_.InvokeVoid(
		w,
		"resetQueryArgument",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) ResetQueryString() {
	_jsii_.InvokeVoid(
		w,
		"resetQueryString",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) ResetUriPath() {
	_jsii_.InvokeVoid(
		w,
		"resetUriPath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleStatementRateBasedStatementCustomKeyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

