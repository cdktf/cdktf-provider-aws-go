// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lblistener

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/lblistener/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbListenerMutualAuthenticationOutputReference interface {
	cdktf.ComplexObject
	AdvertiseTrustStoreCaNames() *string
	SetAdvertiseTrustStoreCaNames(val *string)
	AdvertiseTrustStoreCaNamesInput() *string
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
	IgnoreClientCertificateExpiry() interface{}
	SetIgnoreClientCertificateExpiry(val interface{})
	IgnoreClientCertificateExpiryInput() interface{}
	InternalValue() *LbListenerMutualAuthentication
	SetInternalValue(val *LbListenerMutualAuthentication)
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrustStoreArn() *string
	SetTrustStoreArn(val *string)
	TrustStoreArnInput() *string
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
	ResetAdvertiseTrustStoreCaNames()
	ResetIgnoreClientCertificateExpiry()
	ResetTrustStoreArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LbListenerMutualAuthenticationOutputReference
type jsiiProxy_LbListenerMutualAuthenticationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LbListenerMutualAuthenticationOutputReference) AdvertiseTrustStoreCaNames() *string {
	var returns *string
	_jsii_.Get(
		j,
		"advertiseTrustStoreCaNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerMutualAuthenticationOutputReference) AdvertiseTrustStoreCaNamesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"advertiseTrustStoreCaNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerMutualAuthenticationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerMutualAuthenticationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerMutualAuthenticationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerMutualAuthenticationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerMutualAuthenticationOutputReference) IgnoreClientCertificateExpiry() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreClientCertificateExpiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerMutualAuthenticationOutputReference) IgnoreClientCertificateExpiryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreClientCertificateExpiryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerMutualAuthenticationOutputReference) InternalValue() *LbListenerMutualAuthentication {
	var returns *LbListenerMutualAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerMutualAuthenticationOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerMutualAuthenticationOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerMutualAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerMutualAuthenticationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerMutualAuthenticationOutputReference) TrustStoreArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStoreArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbListenerMutualAuthenticationOutputReference) TrustStoreArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStoreArnInput",
		&returns,
	)
	return returns
}


func NewLbListenerMutualAuthenticationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LbListenerMutualAuthenticationOutputReference {
	_init_.Initialize()

	if err := validateNewLbListenerMutualAuthenticationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LbListenerMutualAuthenticationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.lbListener.LbListenerMutualAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLbListenerMutualAuthenticationOutputReference_Override(l LbListenerMutualAuthenticationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lbListener.LbListenerMutualAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LbListenerMutualAuthenticationOutputReference)SetAdvertiseTrustStoreCaNames(val *string) {
	if err := j.validateSetAdvertiseTrustStoreCaNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advertiseTrustStoreCaNames",
		val,
	)
}

func (j *jsiiProxy_LbListenerMutualAuthenticationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LbListenerMutualAuthenticationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LbListenerMutualAuthenticationOutputReference)SetIgnoreClientCertificateExpiry(val interface{}) {
	if err := j.validateSetIgnoreClientCertificateExpiryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreClientCertificateExpiry",
		val,
	)
}

func (j *jsiiProxy_LbListenerMutualAuthenticationOutputReference)SetInternalValue(val *LbListenerMutualAuthentication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LbListenerMutualAuthenticationOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_LbListenerMutualAuthenticationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LbListenerMutualAuthenticationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LbListenerMutualAuthenticationOutputReference)SetTrustStoreArn(val *string) {
	if err := j.validateSetTrustStoreArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustStoreArn",
		val,
	)
}

func (l *jsiiProxy_LbListenerMutualAuthenticationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerMutualAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LbListenerMutualAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LbListenerMutualAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LbListenerMutualAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LbListenerMutualAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LbListenerMutualAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LbListenerMutualAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LbListenerMutualAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LbListenerMutualAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LbListenerMutualAuthenticationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerMutualAuthenticationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerMutualAuthenticationOutputReference) ResetAdvertiseTrustStoreCaNames() {
	_jsii_.InvokeVoid(
		l,
		"resetAdvertiseTrustStoreCaNames",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerMutualAuthenticationOutputReference) ResetIgnoreClientCertificateExpiry() {
	_jsii_.InvokeVoid(
		l,
		"resetIgnoreClientCertificateExpiry",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerMutualAuthenticationOutputReference) ResetTrustStoreArn() {
	_jsii_.InvokeVoid(
		l,
		"resetTrustStoreArn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbListenerMutualAuthenticationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbListenerMutualAuthenticationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

