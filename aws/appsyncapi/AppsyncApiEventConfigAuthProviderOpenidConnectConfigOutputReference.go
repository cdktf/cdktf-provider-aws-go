// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncapi

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/appsyncapi/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference interface {
	cdktf.ComplexObject
	AuthTtl() *float64
	SetAuthTtl(val *float64)
	AuthTtlInput() *float64
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
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
	IatTtl() *float64
	SetIatTtl(val *float64)
	IatTtlInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Issuer() *string
	SetIssuer(val *string)
	IssuerInput() *string
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
	ResetAuthTtl()
	ResetClientId()
	ResetIatTtl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference
type jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) AuthTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) AuthTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) IatTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iatTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) IatTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iatTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference {
	_init_.Initialize()

	if err := validateNewAppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appsyncApi.AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference_Override(a AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appsyncApi.AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference)SetAuthTtl(val *float64) {
	if err := j.validateSetAuthTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authTtl",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference)SetIatTtl(val *float64) {
	if err := j.validateSetIatTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iatTtl",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) ResetAuthTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) ResetIatTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetIatTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppsyncApiEventConfigAuthProviderOpenidConnectConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

