// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdafunctionurl

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/lambdafunctionurl/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LambdaFunctionUrlCorsOutputReference interface {
	cdktf.ComplexObject
	AllowCredentials() interface{}
	SetAllowCredentials(val interface{})
	AllowCredentialsInput() interface{}
	AllowHeaders() *[]*string
	SetAllowHeaders(val *[]*string)
	AllowHeadersInput() *[]*string
	AllowMethods() *[]*string
	SetAllowMethods(val *[]*string)
	AllowMethodsInput() *[]*string
	AllowOrigins() *[]*string
	SetAllowOrigins(val *[]*string)
	AllowOriginsInput() *[]*string
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
	ExposeHeaders() *[]*string
	SetExposeHeaders(val *[]*string)
	ExposeHeadersInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *LambdaFunctionUrlCors
	SetInternalValue(val *LambdaFunctionUrlCors)
	MaxAge() *float64
	SetMaxAge(val *float64)
	MaxAgeInput() *float64
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
	ResetAllowCredentials()
	ResetAllowHeaders()
	ResetAllowMethods()
	ResetAllowOrigins()
	ResetExposeHeaders()
	ResetMaxAge()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LambdaFunctionUrlCorsOutputReference
type jsiiProxy_LambdaFunctionUrlCorsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference) AllowCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference) AllowCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference) AllowHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference) AllowHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference) AllowMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference) AllowMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference) AllowOrigins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference) AllowOriginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference) ExposeHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exposeHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference) ExposeHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exposeHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference) InternalValue() *LambdaFunctionUrlCors {
	var returns *LambdaFunctionUrlCors
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference) MaxAge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference) MaxAgeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLambdaFunctionUrlCorsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LambdaFunctionUrlCorsOutputReference {
	_init_.Initialize()

	if err := validateNewLambdaFunctionUrlCorsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaFunctionUrlCorsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.lambdaFunctionUrl.LambdaFunctionUrlCorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLambdaFunctionUrlCorsOutputReference_Override(l LambdaFunctionUrlCorsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lambdaFunctionUrl.LambdaFunctionUrlCorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference)SetAllowCredentials(val interface{}) {
	if err := j.validateSetAllowCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowCredentials",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference)SetAllowHeaders(val *[]*string) {
	if err := j.validateSetAllowHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowHeaders",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference)SetAllowMethods(val *[]*string) {
	if err := j.validateSetAllowMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowMethods",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference)SetAllowOrigins(val *[]*string) {
	if err := j.validateSetAllowOriginsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowOrigins",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference)SetExposeHeaders(val *[]*string) {
	if err := j.validateSetExposeHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exposeHeaders",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference)SetInternalValue(val *LambdaFunctionUrlCors) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference)SetMaxAge(val *float64) {
	if err := j.validateSetMaxAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAge",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaFunctionUrlCorsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LambdaFunctionUrlCorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionUrlCorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LambdaFunctionUrlCorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LambdaFunctionUrlCorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LambdaFunctionUrlCorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LambdaFunctionUrlCorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LambdaFunctionUrlCorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LambdaFunctionUrlCorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LambdaFunctionUrlCorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LambdaFunctionUrlCorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LambdaFunctionUrlCorsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionUrlCorsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LambdaFunctionUrlCorsOutputReference) ResetAllowCredentials() {
	_jsii_.InvokeVoid(
		l,
		"resetAllowCredentials",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunctionUrlCorsOutputReference) ResetAllowHeaders() {
	_jsii_.InvokeVoid(
		l,
		"resetAllowHeaders",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunctionUrlCorsOutputReference) ResetAllowMethods() {
	_jsii_.InvokeVoid(
		l,
		"resetAllowMethods",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunctionUrlCorsOutputReference) ResetAllowOrigins() {
	_jsii_.InvokeVoid(
		l,
		"resetAllowOrigins",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunctionUrlCorsOutputReference) ResetExposeHeaders() {
	_jsii_.InvokeVoid(
		l,
		"resetExposeHeaders",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunctionUrlCorsOutputReference) ResetMaxAge() {
	_jsii_.InvokeVoid(
		l,
		"resetMaxAge",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunctionUrlCorsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LambdaFunctionUrlCorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

