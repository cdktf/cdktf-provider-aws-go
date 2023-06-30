package lblistenerrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/lblistenerrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbListenerRuleActionAuthenticateCognitoOutputReference interface {
	cdktf.ComplexObject
	AuthenticationRequestExtraParams() *map[string]*string
	SetAuthenticationRequestExtraParams(val *map[string]*string)
	AuthenticationRequestExtraParamsInput() *map[string]*string
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
	InternalValue() *LbListenerRuleActionAuthenticateCognito
	SetInternalValue(val *LbListenerRuleActionAuthenticateCognito)
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
	UserPoolArn() *string
	SetUserPoolArn(val *string)
	UserPoolArnInput() *string
	UserPoolClientId() *string
	SetUserPoolClientId(val *string)
	UserPoolClientIdInput() *string
	UserPoolDomain() *string
	SetUserPoolDomain(val *string)
	UserPoolDomainInput() *string
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

// The jsii proxy struct for LbListenerRuleActionAuthenticateCognitoOutputReference
type jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) AuthenticationRequestExtraParams() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"authenticationRequestExtraParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) AuthenticationRequestExtraParamsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"authenticationRequestExtraParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) InternalValue() *LbListenerRuleActionAuthenticateCognito {
	var returns *LbListenerRuleActionAuthenticateCognito
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) OnUnauthenticatedRequest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onUnauthenticatedRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) OnUnauthenticatedRequestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onUnauthenticatedRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) SessionCookieName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionCookieName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) SessionCookieNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionCookieNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) SessionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) SessionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) UserPoolArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) UserPoolArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) UserPoolClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) UserPoolClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) UserPoolDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) UserPoolDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolDomainInput",
		&returns,
	)
	return returns
}


func NewLbListenerRuleActionAuthenticateCognitoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LbListenerRuleActionAuthenticateCognitoOutputReference {
	_init_.Initialize()

	if err := validateNewLbListenerRuleActionAuthenticateCognitoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.lbListenerRule.LbListenerRuleActionAuthenticateCognitoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLbListenerRuleActionAuthenticateCognitoOutputReference_Override(l LbListenerRuleActionAuthenticateCognitoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lbListenerRule.LbListenerRuleActionAuthenticateCognitoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference)SetAuthenticationRequestExtraParams(val *map[string]*string) {
	if err := j.validateSetAuthenticationRequestExtraParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationRequestExtraParams",
		val,
	)
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference)SetInternalValue(val *LbListenerRuleActionAuthenticateCognito) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference)SetOnUnauthenticatedRequest(val *string) {
	if err := j.validateSetOnUnauthenticatedRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onUnauthenticatedRequest",
		val,
	)
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference)SetSessionCookieName(val *string) {
	if err := j.validateSetSessionCookieNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionCookieName",
		val,
	)
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference)SetSessionTimeout(val *float64) {
	if err := j.validateSetSessionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionTimeout",
		val,
	)
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference)SetUserPoolArn(val *string) {
	if err := j.validateSetUserPoolArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPoolArn",
		val,
	)
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference)SetUserPoolClientId(val *string) {
	if err := j.validateSetUserPoolClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPoolClientId",
		val,
	)
}

func (j *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference)SetUserPoolDomain(val *string) {
	if err := j.validateSetUserPoolDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPoolDomain",
		val,
	)
}

func (l *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) ResetAuthenticationRequestExtraParams() {
	_jsii_.InvokeVoid(
		l,
		"resetAuthenticationRequestExtraParams",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) ResetOnUnauthenticatedRequest() {
	_jsii_.InvokeVoid(
		l,
		"resetOnUnauthenticatedRequest",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		l,
		"resetScope",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) ResetSessionCookieName() {
	_jsii_.InvokeVoid(
		l,
		"resetSessionCookieName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) ResetSessionTimeout() {
	_jsii_.InvokeVoid(
		l,
		"resetSessionTimeout",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LbListenerRuleActionAuthenticateCognitoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

