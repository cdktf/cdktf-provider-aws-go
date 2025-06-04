// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshgatewayroute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/appmeshgatewayroute/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference interface {
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
	Header() AppmeshGatewayRouteSpecHttp2RouteMatchHeaderList
	HeaderInput() interface{}
	Hostname() AppmeshGatewayRouteSpecHttp2RouteMatchHostnameOutputReference
	HostnameInput() *AppmeshGatewayRouteSpecHttp2RouteMatchHostname
	InternalValue() *AppmeshGatewayRouteSpecHttp2RouteMatch
	SetInternalValue(val *AppmeshGatewayRouteSpecHttp2RouteMatch)
	Path() AppmeshGatewayRouteSpecHttp2RouteMatchPathOutputReference
	PathInput() *AppmeshGatewayRouteSpecHttp2RouteMatchPath
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
	QueryParameter() AppmeshGatewayRouteSpecHttp2RouteMatchQueryParameterList
	QueryParameterInput() interface{}
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
	PutHeader(value interface{})
	PutHostname(value *AppmeshGatewayRouteSpecHttp2RouteMatchHostname)
	PutPath(value *AppmeshGatewayRouteSpecHttp2RouteMatchPath)
	PutQueryParameter(value interface{})
	ResetHeader()
	ResetHostname()
	ResetPath()
	ResetPort()
	ResetPrefix()
	ResetQueryParameter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference
type jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) Header() AppmeshGatewayRouteSpecHttp2RouteMatchHeaderList {
	var returns AppmeshGatewayRouteSpecHttp2RouteMatchHeaderList
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) HeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) Hostname() AppmeshGatewayRouteSpecHttp2RouteMatchHostnameOutputReference {
	var returns AppmeshGatewayRouteSpecHttp2RouteMatchHostnameOutputReference
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) HostnameInput() *AppmeshGatewayRouteSpecHttp2RouteMatchHostname {
	var returns *AppmeshGatewayRouteSpecHttp2RouteMatchHostname
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) InternalValue() *AppmeshGatewayRouteSpecHttp2RouteMatch {
	var returns *AppmeshGatewayRouteSpecHttp2RouteMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) Path() AppmeshGatewayRouteSpecHttp2RouteMatchPathOutputReference {
	var returns AppmeshGatewayRouteSpecHttp2RouteMatchPathOutputReference
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) PathInput() *AppmeshGatewayRouteSpecHttp2RouteMatchPath {
	var returns *AppmeshGatewayRouteSpecHttp2RouteMatchPath
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) QueryParameter() AppmeshGatewayRouteSpecHttp2RouteMatchQueryParameterList {
	var returns AppmeshGatewayRouteSpecHttp2RouteMatchQueryParameterList
	_jsii_.Get(
		j,
		"queryParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) QueryParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryParameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppmeshGatewayRouteSpecHttp2RouteMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference {
	_init_.Initialize()

	if err := validateNewAppmeshGatewayRouteSpecHttp2RouteMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshGatewayRoute.AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppmeshGatewayRouteSpecHttp2RouteMatchOutputReference_Override(a AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshGatewayRoute.AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference)SetInternalValue(val *AppmeshGatewayRouteSpecHttp2RouteMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) PutHeader(value interface{}) {
	if err := a.validatePutHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHeader",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) PutHostname(value *AppmeshGatewayRouteSpecHttp2RouteMatchHostname) {
	if err := a.validatePutHostnameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHostname",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) PutPath(value *AppmeshGatewayRouteSpecHttp2RouteMatchPath) {
	if err := a.validatePutPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPath",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) PutQueryParameter(value interface{}) {
	if err := a.validatePutQueryParameterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putQueryParameter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		a,
		"resetHeader",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		a,
		"resetHostname",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		a,
		"resetPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		a,
		"resetPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) ResetQueryParameter() {
	_jsii_.InvokeVoid(
		a,
		"resetQueryParameter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

