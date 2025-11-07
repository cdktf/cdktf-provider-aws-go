// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package verifiedaccesstrustprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/verifiedaccesstrustprovider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference interface {
	cdktf.ComplexObject
	AuthorizationEndpoint() *string
	SetAuthorizationEndpoint(val *string)
	AuthorizationEndpointInput() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
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
	InternalValue() *VerifiedaccessTrustProviderNativeApplicationOidcOptions
	SetInternalValue(val *VerifiedaccessTrustProviderNativeApplicationOidcOptions)
	Issuer() *string
	SetIssuer(val *string)
	IssuerInput() *string
	PublicSigningKeyEndpoint() *string
	SetPublicSigningKeyEndpoint(val *string)
	PublicSigningKeyEndpointInput() *string
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TokenEndpoint() *string
	SetTokenEndpoint(val *string)
	TokenEndpointInput() *string
	UserInfoEndpoint() *string
	SetUserInfoEndpoint(val *string)
	UserInfoEndpointInput() *string
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
	ResetAuthorizationEndpoint()
	ResetClientId()
	ResetIssuer()
	ResetPublicSigningKeyEndpoint()
	ResetScope()
	ResetTokenEndpoint()
	ResetUserInfoEndpoint()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference
type jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) AuthorizationEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) AuthorizationEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) InternalValue() *VerifiedaccessTrustProviderNativeApplicationOidcOptions {
	var returns *VerifiedaccessTrustProviderNativeApplicationOidcOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) PublicSigningKeyEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicSigningKeyEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) PublicSigningKeyEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicSigningKeyEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) TokenEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) TokenEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) UserInfoEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInfoEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) UserInfoEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInfoEndpointInput",
		&returns,
	)
	return returns
}


func NewVerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewVerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.verifiedaccessTrustProvider.VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference_Override(v VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.verifiedaccessTrustProvider.VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference)SetAuthorizationEndpoint(val *string) {
	if err := j.validateSetAuthorizationEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationEndpoint",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference)SetInternalValue(val *VerifiedaccessTrustProviderNativeApplicationOidcOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference)SetPublicSigningKeyEndpoint(val *string) {
	if err := j.validateSetPublicSigningKeyEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicSigningKeyEndpoint",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference)SetTokenEndpoint(val *string) {
	if err := j.validateSetTokenEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenEndpoint",
		val,
	)
}

func (j *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference)SetUserInfoEndpoint(val *string) {
	if err := j.validateSetUserInfoEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userInfoEndpoint",
		val,
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) ResetAuthorizationEndpoint() {
	_jsii_.InvokeVoid(
		v,
		"resetAuthorizationEndpoint",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		v,
		"resetClientId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) ResetIssuer() {
	_jsii_.InvokeVoid(
		v,
		"resetIssuer",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) ResetPublicSigningKeyEndpoint() {
	_jsii_.InvokeVoid(
		v,
		"resetPublicSigningKeyEndpoint",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		v,
		"resetScope",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) ResetTokenEndpoint() {
	_jsii_.InvokeVoid(
		v,
		"resetTokenEndpoint",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) ResetUserInfoEndpoint() {
	_jsii_.InvokeVoid(
		v,
		"resetUserInfoEndpoint",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedaccessTrustProviderNativeApplicationOidcOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

