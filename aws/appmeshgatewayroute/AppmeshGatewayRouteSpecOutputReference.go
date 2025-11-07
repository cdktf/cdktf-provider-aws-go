// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshgatewayroute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/appmeshgatewayroute/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppmeshGatewayRouteSpecOutputReference interface {
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
	GrpcRoute() AppmeshGatewayRouteSpecGrpcRouteOutputReference
	GrpcRouteInput() *AppmeshGatewayRouteSpecGrpcRoute
	Http2Route() AppmeshGatewayRouteSpecHttp2RouteOutputReference
	Http2RouteInput() *AppmeshGatewayRouteSpecHttp2Route
	HttpRoute() AppmeshGatewayRouteSpecHttpRouteOutputReference
	HttpRouteInput() *AppmeshGatewayRouteSpecHttpRoute
	InternalValue() *AppmeshGatewayRouteSpec
	SetInternalValue(val *AppmeshGatewayRouteSpec)
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutGrpcRoute(value *AppmeshGatewayRouteSpecGrpcRoute)
	PutHttp2Route(value *AppmeshGatewayRouteSpecHttp2Route)
	PutHttpRoute(value *AppmeshGatewayRouteSpecHttpRoute)
	ResetGrpcRoute()
	ResetHttp2Route()
	ResetHttpRoute()
	ResetPriority()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppmeshGatewayRouteSpecOutputReference
type jsiiProxy_AppmeshGatewayRouteSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) GrpcRoute() AppmeshGatewayRouteSpecGrpcRouteOutputReference {
	var returns AppmeshGatewayRouteSpecGrpcRouteOutputReference
	_jsii_.Get(
		j,
		"grpcRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) GrpcRouteInput() *AppmeshGatewayRouteSpecGrpcRoute {
	var returns *AppmeshGatewayRouteSpecGrpcRoute
	_jsii_.Get(
		j,
		"grpcRouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) Http2Route() AppmeshGatewayRouteSpecHttp2RouteOutputReference {
	var returns AppmeshGatewayRouteSpecHttp2RouteOutputReference
	_jsii_.Get(
		j,
		"http2Route",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) Http2RouteInput() *AppmeshGatewayRouteSpecHttp2Route {
	var returns *AppmeshGatewayRouteSpecHttp2Route
	_jsii_.Get(
		j,
		"http2RouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) HttpRoute() AppmeshGatewayRouteSpecHttpRouteOutputReference {
	var returns AppmeshGatewayRouteSpecHttpRouteOutputReference
	_jsii_.Get(
		j,
		"httpRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) HttpRouteInput() *AppmeshGatewayRouteSpecHttpRoute {
	var returns *AppmeshGatewayRouteSpecHttpRoute
	_jsii_.Get(
		j,
		"httpRouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) InternalValue() *AppmeshGatewayRouteSpec {
	var returns *AppmeshGatewayRouteSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppmeshGatewayRouteSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppmeshGatewayRouteSpecOutputReference {
	_init_.Initialize()

	if err := validateNewAppmeshGatewayRouteSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppmeshGatewayRouteSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshGatewayRoute.AppmeshGatewayRouteSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppmeshGatewayRouteSpecOutputReference_Override(a AppmeshGatewayRouteSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshGatewayRoute.AppmeshGatewayRouteSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecOutputReference)SetInternalValue(val *AppmeshGatewayRouteSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppmeshGatewayRouteSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) PutGrpcRoute(value *AppmeshGatewayRouteSpecGrpcRoute) {
	if err := a.validatePutGrpcRouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGrpcRoute",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) PutHttp2Route(value *AppmeshGatewayRouteSpecHttp2Route) {
	if err := a.validatePutHttp2RouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttp2Route",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) PutHttpRoute(value *AppmeshGatewayRouteSpecHttpRoute) {
	if err := a.validatePutHttpRouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttpRoute",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) ResetGrpcRoute() {
	_jsii_.InvokeVoid(
		a,
		"resetGrpcRoute",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) ResetHttp2Route() {
	_jsii_.InvokeVoid(
		a,
		"resetHttp2Route",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) ResetHttpRoute() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpRoute",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		a,
		"resetPriority",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshGatewayRouteSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

