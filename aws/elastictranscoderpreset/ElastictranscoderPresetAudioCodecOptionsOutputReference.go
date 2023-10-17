// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastictranscoderpreset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/elastictranscoderpreset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastictranscoderPresetAudioCodecOptionsOutputReference interface {
	cdktf.ComplexObject
	BitDepth() *string
	SetBitDepth(val *string)
	BitDepthInput() *string
	BitOrder() *string
	SetBitOrder(val *string)
	BitOrderInput() *string
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
	InternalValue() *ElastictranscoderPresetAudioCodecOptions
	SetInternalValue(val *ElastictranscoderPresetAudioCodecOptions)
	Profile() *string
	SetProfile(val *string)
	ProfileInput() *string
	Signed() *string
	SetSigned(val *string)
	SignedInput() *string
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
	ResetBitDepth()
	ResetBitOrder()
	ResetProfile()
	ResetSigned()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElastictranscoderPresetAudioCodecOptionsOutputReference
type jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) BitDepth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) BitDepthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitDepthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) BitOrder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) BitOrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) InternalValue() *ElastictranscoderPresetAudioCodecOptions {
	var returns *ElastictranscoderPresetAudioCodecOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) ProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) Signed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) SignedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElastictranscoderPresetAudioCodecOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ElastictranscoderPresetAudioCodecOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewElastictranscoderPresetAudioCodecOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.elastictranscoderPreset.ElastictranscoderPresetAudioCodecOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElastictranscoderPresetAudioCodecOptionsOutputReference_Override(e ElastictranscoderPresetAudioCodecOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.elastictranscoderPreset.ElastictranscoderPresetAudioCodecOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference)SetBitDepth(val *string) {
	if err := j.validateSetBitDepthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitDepth",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference)SetBitOrder(val *string) {
	if err := j.validateSetBitOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitOrder",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference)SetInternalValue(val *ElastictranscoderPresetAudioCodecOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference)SetProfile(val *string) {
	if err := j.validateSetProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profile",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference)SetSigned(val *string) {
	if err := j.validateSetSignedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signed",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) ResetBitDepth() {
	_jsii_.InvokeVoid(
		e,
		"resetBitDepth",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) ResetBitOrder() {
	_jsii_.InvokeVoid(
		e,
		"resetBitOrder",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) ResetProfile() {
	_jsii_.InvokeVoid(
		e,
		"resetProfile",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) ResetSigned() {
	_jsii_.InvokeVoid(
		e,
		"resetSigned",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

