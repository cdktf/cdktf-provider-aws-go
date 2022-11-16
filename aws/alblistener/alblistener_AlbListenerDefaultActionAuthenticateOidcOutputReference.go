package alblistener

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/alblistener/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbListenerDefaultActionAuthenticateOidcOutputReference interface {
	cdktf.ComplexObject
	AuthenticationRequestExtraParams() *map[string]*string
	SetAuthenticationRequestExtraParams(val *map[string]*string)
	AuthenticationRequestExtraParamsInput() *map[string]*string
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
	InternalValue() *AlbListenerDefaultActionAuthenticateOidc
	SetInternalValue(val *AlbListenerDefaultActionAuthenticateOidc)
	Issuer() *string
	SetIssuer(val *string)
	IssuerInput() *string
	OnUnauthenticatedRequest() *string
	SetOnUnauthenticatedRequest(val *string)
	OnUnauthenticatedRequestInput() *string
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
	SessionCookieName() *string
	SetSessionCookieName(val *string)
	SessionCookieNameInput() *string
	SessionTimeout() *float64
	SetSessionTimeout(val *float64)
	SessionTimeoutInput() *float64
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAuthenticationRequestExtraParams()
	ResetOnUnauthenticatedRequest()
	ResetScope()
	ResetSessionCookieName()
	ResetSessionTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlbListenerDefaultActionAuthenticateOidcOutputReference
type jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) AuthenticationRequestExtraParams() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"authenticationRequestExtraParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) AuthenticationRequestExtraParamsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"authenticationRequestExtraParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) AuthorizationEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) AuthorizationEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) InternalValue() *AlbListenerDefaultActionAuthenticateOidc {
	var returns *AlbListenerDefaultActionAuthenticateOidc
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) OnUnauthenticatedRequest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onUnauthenticatedRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) OnUnauthenticatedRequestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onUnauthenticatedRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) SessionCookieName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionCookieName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) SessionCookieNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionCookieNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) SessionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) SessionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) TokenEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) TokenEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) UserInfoEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInfoEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) UserInfoEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInfoEndpointInput",
		&returns,
	)
	return returns
}


func NewAlbListenerDefaultActionAuthenticateOidcOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlbListenerDefaultActionAuthenticateOidcOutputReference {
	_init_.Initialize()

	if err := validateNewAlbListenerDefaultActionAuthenticateOidcOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.albListener.AlbListenerDefaultActionAuthenticateOidcOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlbListenerDefaultActionAuthenticateOidcOutputReference_Override(a AlbListenerDefaultActionAuthenticateOidcOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.albListener.AlbListenerDefaultActionAuthenticateOidcOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference)SetAuthenticationRequestExtraParams(val *map[string]*string) {
	if err := j.validateSetAuthenticationRequestExtraParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationRequestExtraParams",
		val,
	)
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference)SetAuthorizationEndpoint(val *string) {
	if err := j.validateSetAuthorizationEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationEndpoint",
		val,
	)
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference)SetInternalValue(val *AlbListenerDefaultActionAuthenticateOidc) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference)SetOnUnauthenticatedRequest(val *string) {
	if err := j.validateSetOnUnauthenticatedRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onUnauthenticatedRequest",
		val,
	)
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference)SetSessionCookieName(val *string) {
	if err := j.validateSetSessionCookieNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionCookieName",
		val,
	)
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference)SetSessionTimeout(val *float64) {
	if err := j.validateSetSessionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionTimeout",
		val,
	)
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference)SetTokenEndpoint(val *string) {
	if err := j.validateSetTokenEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenEndpoint",
		val,
	)
}

func (j *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference)SetUserInfoEndpoint(val *string) {
	if err := j.validateSetUserInfoEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userInfoEndpoint",
		val,
	)
}

func (a *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) ResetAuthenticationRequestExtraParams() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthenticationRequestExtraParams",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) ResetOnUnauthenticatedRequest() {
	_jsii_.InvokeVoid(
		a,
		"resetOnUnauthenticatedRequest",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		a,
		"resetScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) ResetSessionCookieName() {
	_jsii_.InvokeVoid(
		a,
		"resetSessionCookieName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) ResetSessionTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetSessionTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AlbListenerDefaultActionAuthenticateOidcOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

