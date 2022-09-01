package elb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/elb/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlbListenerDefaultActionOutputReference interface {
	cdktf.ComplexObject
	AuthenticateCognito() AlbListenerDefaultActionAuthenticateCognitoOutputReference
	AuthenticateCognitoInput() *AlbListenerDefaultActionAuthenticateCognito
	AuthenticateOidc() AlbListenerDefaultActionAuthenticateOidcOutputReference
	AuthenticateOidcInput() *AlbListenerDefaultActionAuthenticateOidc
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
	FixedResponse() AlbListenerDefaultActionFixedResponseOutputReference
	FixedResponseInput() *AlbListenerDefaultActionFixedResponse
	Forward() AlbListenerDefaultActionForwardOutputReference
	ForwardInput() *AlbListenerDefaultActionForward
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Order() *float64
	SetOrder(val *float64)
	OrderInput() *float64
	Redirect() AlbListenerDefaultActionRedirectOutputReference
	RedirectInput() *AlbListenerDefaultActionRedirect
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
	PutAuthenticateCognito(value *AlbListenerDefaultActionAuthenticateCognito)
	PutAuthenticateOidc(value *AlbListenerDefaultActionAuthenticateOidc)
	PutFixedResponse(value *AlbListenerDefaultActionFixedResponse)
	PutForward(value *AlbListenerDefaultActionForward)
	PutRedirect(value *AlbListenerDefaultActionRedirect)
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

// The jsii proxy struct for AlbListenerDefaultActionOutputReference
type jsiiProxy_AlbListenerDefaultActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference) AuthenticateCognito() AlbListenerDefaultActionAuthenticateCognitoOutputReference {
	var returns AlbListenerDefaultActionAuthenticateCognitoOutputReference
	_jsii_.Get(
		j,
		"authenticateCognito",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference) AuthenticateCognitoInput() *AlbListenerDefaultActionAuthenticateCognito {
	var returns *AlbListenerDefaultActionAuthenticateCognito
	_jsii_.Get(
		j,
		"authenticateCognitoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference) AuthenticateOidc() AlbListenerDefaultActionAuthenticateOidcOutputReference {
	var returns AlbListenerDefaultActionAuthenticateOidcOutputReference
	_jsii_.Get(
		j,
		"authenticateOidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference) AuthenticateOidcInput() *AlbListenerDefaultActionAuthenticateOidc {
	var returns *AlbListenerDefaultActionAuthenticateOidc
	_jsii_.Get(
		j,
		"authenticateOidcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference) FixedResponse() AlbListenerDefaultActionFixedResponseOutputReference {
	var returns AlbListenerDefaultActionFixedResponseOutputReference
	_jsii_.Get(
		j,
		"fixedResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference) FixedResponseInput() *AlbListenerDefaultActionFixedResponse {
	var returns *AlbListenerDefaultActionFixedResponse
	_jsii_.Get(
		j,
		"fixedResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference) Forward() AlbListenerDefaultActionForwardOutputReference {
	var returns AlbListenerDefaultActionForwardOutputReference
	_jsii_.Get(
		j,
		"forward",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference) ForwardInput() *AlbListenerDefaultActionForward {
	var returns *AlbListenerDefaultActionForward
	_jsii_.Get(
		j,
		"forwardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference) Order() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference) OrderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference) Redirect() AlbListenerDefaultActionRedirectOutputReference {
	var returns AlbListenerDefaultActionRedirectOutputReference
	_jsii_.Get(
		j,
		"redirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference) RedirectInput() *AlbListenerDefaultActionRedirect {
	var returns *AlbListenerDefaultActionRedirect
	_jsii_.Get(
		j,
		"redirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference) TargetGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference) TargetGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewAlbListenerDefaultActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AlbListenerDefaultActionOutputReference {
	_init_.Initialize()

	if err := validateNewAlbListenerDefaultActionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlbListenerDefaultActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.elb.AlbListenerDefaultActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAlbListenerDefaultActionOutputReference_Override(a AlbListenerDefaultActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.elb.AlbListenerDefaultActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference)SetOrder(val *float64) {
	if err := j.validateSetOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference)SetTargetGroupArn(val *string) {
	if err := j.validateSetTargetGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetGroupArn",
		val,
	)
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AlbListenerDefaultActionOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) PutAuthenticateCognito(value *AlbListenerDefaultActionAuthenticateCognito) {
	if err := a.validatePutAuthenticateCognitoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuthenticateCognito",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) PutAuthenticateOidc(value *AlbListenerDefaultActionAuthenticateOidc) {
	if err := a.validatePutAuthenticateOidcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuthenticateOidc",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) PutFixedResponse(value *AlbListenerDefaultActionFixedResponse) {
	if err := a.validatePutFixedResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFixedResponse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) PutForward(value *AlbListenerDefaultActionForward) {
	if err := a.validatePutForwardParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putForward",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) PutRedirect(value *AlbListenerDefaultActionRedirect) {
	if err := a.validatePutRedirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedirect",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) ResetAuthenticateCognito() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthenticateCognito",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) ResetAuthenticateOidc() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthenticateOidc",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) ResetFixedResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetFixedResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) ResetForward() {
	_jsii_.InvokeVoid(
		a,
		"resetForward",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) ResetOrder() {
	_jsii_.InvokeVoid(
		a,
		"resetOrder",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) ResetRedirect() {
	_jsii_.InvokeVoid(
		a,
		"resetRedirect",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) ResetTargetGroupArn() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetGroupArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AlbListenerDefaultActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

