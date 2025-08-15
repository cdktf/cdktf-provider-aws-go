// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncapi

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/appsyncapi/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppsyncApiEventConfigOutputReference interface {
	cdktf.ComplexObject
	AuthProvider() AppsyncApiEventConfigAuthProviderList
	AuthProviderInput() interface{}
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
	ConnectionAuthMode() AppsyncApiEventConfigConnectionAuthModeList
	ConnectionAuthModeInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultPublishAuthMode() AppsyncApiEventConfigDefaultPublishAuthModeList
	DefaultPublishAuthModeInput() interface{}
	DefaultSubscribeAuthMode() AppsyncApiEventConfigDefaultSubscribeAuthModeList
	DefaultSubscribeAuthModeInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogConfig() AppsyncApiEventConfigLogConfigList
	LogConfigInput() interface{}
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
	PutAuthProvider(value interface{})
	PutConnectionAuthMode(value interface{})
	PutDefaultPublishAuthMode(value interface{})
	PutDefaultSubscribeAuthMode(value interface{})
	PutLogConfig(value interface{})
	ResetAuthProvider()
	ResetConnectionAuthMode()
	ResetDefaultPublishAuthMode()
	ResetDefaultSubscribeAuthMode()
	ResetLogConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppsyncApiEventConfigOutputReference
type jsiiProxy_AppsyncApiEventConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) AuthProvider() AppsyncApiEventConfigAuthProviderList {
	var returns AppsyncApiEventConfigAuthProviderList
	_jsii_.Get(
		j,
		"authProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) AuthProviderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) ConnectionAuthMode() AppsyncApiEventConfigConnectionAuthModeList {
	var returns AppsyncApiEventConfigConnectionAuthModeList
	_jsii_.Get(
		j,
		"connectionAuthMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) ConnectionAuthModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionAuthModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) DefaultPublishAuthMode() AppsyncApiEventConfigDefaultPublishAuthModeList {
	var returns AppsyncApiEventConfigDefaultPublishAuthModeList
	_jsii_.Get(
		j,
		"defaultPublishAuthMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) DefaultPublishAuthModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultPublishAuthModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) DefaultSubscribeAuthMode() AppsyncApiEventConfigDefaultSubscribeAuthModeList {
	var returns AppsyncApiEventConfigDefaultSubscribeAuthModeList
	_jsii_.Get(
		j,
		"defaultSubscribeAuthMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) DefaultSubscribeAuthModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultSubscribeAuthModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) LogConfig() AppsyncApiEventConfigLogConfigList {
	var returns AppsyncApiEventConfigLogConfigList
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) LogConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppsyncApiEventConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AppsyncApiEventConfigOutputReference {
	_init_.Initialize()

	if err := validateNewAppsyncApiEventConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppsyncApiEventConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appsyncApi.AppsyncApiEventConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAppsyncApiEventConfigOutputReference_Override(a AppsyncApiEventConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appsyncApi.AppsyncApiEventConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiEventConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) PutAuthProvider(value interface{}) {
	if err := a.validatePutAuthProviderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuthProvider",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) PutConnectionAuthMode(value interface{}) {
	if err := a.validatePutConnectionAuthModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConnectionAuthMode",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) PutDefaultPublishAuthMode(value interface{}) {
	if err := a.validatePutDefaultPublishAuthModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDefaultPublishAuthMode",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) PutDefaultSubscribeAuthMode(value interface{}) {
	if err := a.validatePutDefaultSubscribeAuthModeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDefaultSubscribeAuthMode",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) PutLogConfig(value interface{}) {
	if err := a.validatePutLogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) ResetAuthProvider() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthProvider",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) ResetConnectionAuthMode() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionAuthMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) ResetDefaultPublishAuthMode() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultPublishAuthMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) ResetDefaultSubscribeAuthMode() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultSubscribeAuthMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) ResetLogConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLogConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppsyncApiEventConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

