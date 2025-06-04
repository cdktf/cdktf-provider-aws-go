// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paymentcryptographykey

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v20/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v20/paymentcryptographykey/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference interface {
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
	Decrypt() interface{}
	SetDecrypt(val interface{})
	DecryptInput() interface{}
	DeriveKey() interface{}
	SetDeriveKey(val interface{})
	DeriveKeyInput() interface{}
	Encrypt() interface{}
	SetEncrypt(val interface{})
	EncryptInput() interface{}
	// Experimental.
	Fqn() *string
	Generate() interface{}
	SetGenerate(val interface{})
	GenerateInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NoRestrictions() interface{}
	SetNoRestrictions(val interface{})
	NoRestrictionsInput() interface{}
	Sign() interface{}
	SetSign(val interface{})
	SignInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Unwrap() interface{}
	SetUnwrap(val interface{})
	UnwrapInput() interface{}
	Verify() interface{}
	SetVerify(val interface{})
	VerifyInput() interface{}
	Wrap() interface{}
	SetWrap(val interface{})
	WrapInput() interface{}
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
	ResetDecrypt()
	ResetDeriveKey()
	ResetEncrypt()
	ResetGenerate()
	ResetNoRestrictions()
	ResetSign()
	ResetUnwrap()
	ResetVerify()
	ResetWrap()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference
type jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) Decrypt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"decrypt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) DecryptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"decryptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) DeriveKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deriveKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) DeriveKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deriveKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) Encrypt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encrypt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) EncryptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) Generate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) GenerateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) NoRestrictions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noRestrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) NoRestrictionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noRestrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) Sign() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) SignInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) Unwrap() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unwrap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) UnwrapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unwrapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) Verify() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) VerifyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) Wrap() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wrap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) WrapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wrapInput",
		&returns,
	)
	return returns
}


func NewPaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference {
	_init_.Initialize()

	if err := validateNewPaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.paymentcryptographyKey.PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference_Override(p PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.paymentcryptographyKey.PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference)SetDecrypt(val interface{}) {
	if err := j.validateSetDecryptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"decrypt",
		val,
	)
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference)SetDeriveKey(val interface{}) {
	if err := j.validateSetDeriveKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deriveKey",
		val,
	)
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference)SetEncrypt(val interface{}) {
	if err := j.validateSetEncryptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encrypt",
		val,
	)
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference)SetGenerate(val interface{}) {
	if err := j.validateSetGenerateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"generate",
		val,
	)
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference)SetNoRestrictions(val interface{}) {
	if err := j.validateSetNoRestrictionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noRestrictions",
		val,
	)
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference)SetSign(val interface{}) {
	if err := j.validateSetSignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sign",
		val,
	)
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference)SetUnwrap(val interface{}) {
	if err := j.validateSetUnwrapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unwrap",
		val,
	)
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference)SetVerify(val interface{}) {
	if err := j.validateSetVerifyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verify",
		val,
	)
}

func (j *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference)SetWrap(val interface{}) {
	if err := j.validateSetWrapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrap",
		val,
	)
}

func (p *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) ResetDecrypt() {
	_jsii_.InvokeVoid(
		p,
		"resetDecrypt",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) ResetDeriveKey() {
	_jsii_.InvokeVoid(
		p,
		"resetDeriveKey",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) ResetEncrypt() {
	_jsii_.InvokeVoid(
		p,
		"resetEncrypt",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) ResetGenerate() {
	_jsii_.InvokeVoid(
		p,
		"resetGenerate",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) ResetNoRestrictions() {
	_jsii_.InvokeVoid(
		p,
		"resetNoRestrictions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) ResetSign() {
	_jsii_.InvokeVoid(
		p,
		"resetSign",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) ResetUnwrap() {
	_jsii_.InvokeVoid(
		p,
		"resetUnwrap",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) ResetVerify() {
	_jsii_.InvokeVoid(
		p,
		"resetVerify",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) ResetWrap() {
	_jsii_.InvokeVoid(
		p,
		"resetWrap",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaymentcryptographyKeyKeyAttributesKeyModesOfUseOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

