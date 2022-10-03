package alblistenerrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v10/alblistenerrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbListenerRuleActionOutputReference interface {
	cdktf.ComplexObject
	AuthenticateCognito() AlbListenerRuleActionAuthenticateCognitoOutputReference
	AuthenticateCognitoInput() *AlbListenerRuleActionAuthenticateCognito
	AuthenticateOidc() AlbListenerRuleActionAuthenticateOidcOutputReference
	AuthenticateOidcInput() *AlbListenerRuleActionAuthenticateOidc
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
	FixedResponse() AlbListenerRuleActionFixedResponseOutputReference
	FixedResponseInput() *AlbListenerRuleActionFixedResponse
	Forward() AlbListenerRuleActionForwardOutputReference
	ForwardInput() *AlbListenerRuleActionForward
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Order() *float64
	SetOrder(val *float64)
	OrderInput() *float64
	Redirect() AlbListenerRuleActionRedirectOutputReference
	RedirectInput() *AlbListenerRuleActionRedirect
	TargetGroupArn() *string
	SetTargetGroupArn(val *string)
	TargetGroupArnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutAuthenticateCognito(value *AlbListenerRuleActionAuthenticateCognito)
	PutAuthenticateOidc(value *AlbListenerRuleActionAuthenticateOidc)
	PutFixedResponse(value *AlbListenerRuleActionFixedResponse)
	PutForward(value *AlbListenerRuleActionForward)
	PutRedirect(value *AlbListenerRuleActionRedirect)
	ResetAuthenticateCognito()
	ResetAuthenticateOidc()
	ResetFixedResponse()
	ResetForward()
	ResetOrder()
	ResetRedirect()
	ResetTargetGroupArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlbListenerRuleActionOutputReference
type jsiiProxy_AlbListenerRuleActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference) AuthenticateCognito() AlbListenerRuleActionAuthenticateCognitoOutputReference {
	var returns AlbListenerRuleActionAuthenticateCognitoOutputReference
	_jsii_.Get(
		j,
		"authenticateCognito",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference) AuthenticateCognitoInput() *AlbListenerRuleActionAuthenticateCognito {
	var returns *AlbListenerRuleActionAuthenticateCognito
	_jsii_.Get(
		j,
		"authenticateCognitoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference) AuthenticateOidc() AlbListenerRuleActionAuthenticateOidcOutputReference {
	var returns AlbListenerRuleActionAuthenticateOidcOutputReference
	_jsii_.Get(
		j,
		"authenticateOidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference) AuthenticateOidcInput() *AlbListenerRuleActionAuthenticateOidc {
	var returns *AlbListenerRuleActionAuthenticateOidc
	_jsii_.Get(
		j,
		"authenticateOidcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference) FixedResponse() AlbListenerRuleActionFixedResponseOutputReference {
	var returns AlbListenerRuleActionFixedResponseOutputReference
	_jsii_.Get(
		j,
		"fixedResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference) FixedResponseInput() *AlbListenerRuleActionFixedResponse {
	var returns *AlbListenerRuleActionFixedResponse
	_jsii_.Get(
		j,
		"fixedResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference) Forward() AlbListenerRuleActionForwardOutputReference {
	var returns AlbListenerRuleActionForwardOutputReference
	_jsii_.Get(
		j,
		"forward",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference) ForwardInput() *AlbListenerRuleActionForward {
	var returns *AlbListenerRuleActionForward
	_jsii_.Get(
		j,
		"forwardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference) Order() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference) OrderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference) Redirect() AlbListenerRuleActionRedirectOutputReference {
	var returns AlbListenerRuleActionRedirectOutputReference
	_jsii_.Get(
		j,
		"redirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference) RedirectInput() *AlbListenerRuleActionRedirect {
	var returns *AlbListenerRuleActionRedirect
	_jsii_.Get(
		j,
		"redirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference) TargetGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference) TargetGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewAlbListenerRuleActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AlbListenerRuleActionOutputReference {
	_init_.Initialize()

	if err := validateNewAlbListenerRuleActionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlbListenerRuleActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.albListenerRule.AlbListenerRuleActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAlbListenerRuleActionOutputReference_Override(a AlbListenerRuleActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.albListenerRule.AlbListenerRuleActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference)SetOrder(val *float64) {
	if err := j.validateSetOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference)SetTargetGroupArn(val *string) {
	if err := j.validateSetTargetGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetGroupArn",
		val,
	)
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AlbListenerRuleActionOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) PutAuthenticateCognito(value *AlbListenerRuleActionAuthenticateCognito) {
	if err := a.validatePutAuthenticateCognitoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuthenticateCognito",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) PutAuthenticateOidc(value *AlbListenerRuleActionAuthenticateOidc) {
	if err := a.validatePutAuthenticateOidcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuthenticateOidc",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) PutFixedResponse(value *AlbListenerRuleActionFixedResponse) {
	if err := a.validatePutFixedResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFixedResponse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) PutForward(value *AlbListenerRuleActionForward) {
	if err := a.validatePutForwardParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putForward",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) PutRedirect(value *AlbListenerRuleActionRedirect) {
	if err := a.validatePutRedirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedirect",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) ResetAuthenticateCognito() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthenticateCognito",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) ResetAuthenticateOidc() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthenticateOidc",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) ResetFixedResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetFixedResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) ResetForward() {
	_jsii_.InvokeVoid(
		a,
		"resetForward",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) ResetOrder() {
	_jsii_.InvokeVoid(
		a,
		"resetOrder",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) ResetRedirect() {
	_jsii_.InvokeVoid(
		a,
		"resetRedirect",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) ResetTargetGroupArn() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetGroupArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AlbListenerRuleActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

