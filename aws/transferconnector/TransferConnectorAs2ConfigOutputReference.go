// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transferconnector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/transferconnector/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TransferConnectorAs2ConfigOutputReference interface {
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
	Compression() *string
	SetCompression(val *string)
	CompressionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EncryptionAlgorithm() *string
	SetEncryptionAlgorithm(val *string)
	EncryptionAlgorithmInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *TransferConnectorAs2Config
	SetInternalValue(val *TransferConnectorAs2Config)
	LocalProfileId() *string
	SetLocalProfileId(val *string)
	LocalProfileIdInput() *string
	MdnResponse() *string
	SetMdnResponse(val *string)
	MdnResponseInput() *string
	MdnSigningAlgorithm() *string
	SetMdnSigningAlgorithm(val *string)
	MdnSigningAlgorithmInput() *string
	MessageSubject() *string
	SetMessageSubject(val *string)
	MessageSubjectInput() *string
	PartnerProfileId() *string
	SetPartnerProfileId(val *string)
	PartnerProfileIdInput() *string
	SigningAlgorithm() *string
	SetSigningAlgorithm(val *string)
	SigningAlgorithmInput() *string
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
	ResetMdnSigningAlgorithm()
	ResetMessageSubject()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TransferConnectorAs2ConfigOutputReference
type jsiiProxy_TransferConnectorAs2ConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference) Compression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference) CompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference) EncryptionAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference) EncryptionAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference) InternalValue() *TransferConnectorAs2Config {
	var returns *TransferConnectorAs2Config
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference) LocalProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference) LocalProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference) MdnResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mdnResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference) MdnResponseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mdnResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference) MdnSigningAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mdnSigningAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference) MdnSigningAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mdnSigningAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference) MessageSubject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageSubject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference) MessageSubjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageSubjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference) PartnerProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference) PartnerProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference) SigningAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference) SigningAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTransferConnectorAs2ConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TransferConnectorAs2ConfigOutputReference {
	_init_.Initialize()

	if err := validateNewTransferConnectorAs2ConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TransferConnectorAs2ConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.transferConnector.TransferConnectorAs2ConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTransferConnectorAs2ConfigOutputReference_Override(t TransferConnectorAs2ConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.transferConnector.TransferConnectorAs2ConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference)SetCompression(val *string) {
	if err := j.validateSetCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compression",
		val,
	)
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference)SetEncryptionAlgorithm(val *string) {
	if err := j.validateSetEncryptionAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionAlgorithm",
		val,
	)
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference)SetInternalValue(val *TransferConnectorAs2Config) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference)SetLocalProfileId(val *string) {
	if err := j.validateSetLocalProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localProfileId",
		val,
	)
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference)SetMdnResponse(val *string) {
	if err := j.validateSetMdnResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mdnResponse",
		val,
	)
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference)SetMdnSigningAlgorithm(val *string) {
	if err := j.validateSetMdnSigningAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mdnSigningAlgorithm",
		val,
	)
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference)SetMessageSubject(val *string) {
	if err := j.validateSetMessageSubjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageSubject",
		val,
	)
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference)SetPartnerProfileId(val *string) {
	if err := j.validateSetPartnerProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partnerProfileId",
		val,
	)
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference)SetSigningAlgorithm(val *string) {
	if err := j.validateSetSigningAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signingAlgorithm",
		val,
	)
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TransferConnectorAs2ConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TransferConnectorAs2ConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferConnectorAs2ConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferConnectorAs2ConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferConnectorAs2ConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferConnectorAs2ConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferConnectorAs2ConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferConnectorAs2ConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferConnectorAs2ConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferConnectorAs2ConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferConnectorAs2ConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferConnectorAs2ConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferConnectorAs2ConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferConnectorAs2ConfigOutputReference) ResetMdnSigningAlgorithm() {
	_jsii_.InvokeVoid(
		t,
		"resetMdnSigningAlgorithm",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferConnectorAs2ConfigOutputReference) ResetMessageSubject() {
	_jsii_.InvokeVoid(
		t,
		"resetMessageSubject",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TransferConnectorAs2ConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TransferConnectorAs2ConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

