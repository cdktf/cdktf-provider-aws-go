package medialivechannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/medialivechannel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference interface {
	cdktf.ComplexObject
	Bitrate() *float64
	SetBitrate(val *float64)
	BitrateInput() *float64
	CodingMode() *string
	SetCodingMode(val *string)
	CodingModeInput() *string
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
	InputType() *string
	SetInputType(val *string)
	InputTypeInput() *string
	InternalValue() *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettings
	SetInternalValue(val *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettings)
	Profile() *string
	SetProfile(val *string)
	ProfileInput() *string
	RateControlMode() *string
	SetRateControlMode(val *string)
	RateControlModeInput() *string
	RawFormat() *string
	SetRawFormat(val *string)
	RawFormatInput() *string
	SampleRate() *float64
	SetSampleRate(val *float64)
	SampleRateInput() *float64
	Spec() *string
	SetSpec(val *string)
	SpecInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VbrQuality() *string
	SetVbrQuality(val *string)
	VbrQualityInput() *string
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
	ResetBitrate()
	ResetCodingMode()
	ResetInputType()
	ResetProfile()
	ResetRateControlMode()
	ResetRawFormat()
	ResetSampleRate()
	ResetSpec()
	ResetVbrQuality()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference
type jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) Bitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) BitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) CodingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) CodingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) InputType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) InputTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) InternalValue() *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettings {
	var returns *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) ProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) RateControlMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateControlMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) RateControlModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateControlModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) RawFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rawFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) RawFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rawFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) SampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) SampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) Spec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) SpecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"specInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) VbrQuality() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vbrQuality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) VbrQualityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vbrQualityInput",
		&returns,
	)
	return returns
}


func NewMedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference_Override(m MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference)SetBitrate(val *float64) {
	if err := j.validateSetBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitrate",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference)SetCodingMode(val *string) {
	if err := j.validateSetCodingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codingMode",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference)SetInputType(val *string) {
	if err := j.validateSetInputTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputType",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference)SetInternalValue(val *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference)SetProfile(val *string) {
	if err := j.validateSetProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profile",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference)SetRateControlMode(val *string) {
	if err := j.validateSetRateControlModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rateControlMode",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference)SetRawFormat(val *string) {
	if err := j.validateSetRawFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rawFormat",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference)SetSampleRate(val *float64) {
	if err := j.validateSetSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleRate",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference)SetSpec(val *string) {
	if err := j.validateSetSpecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spec",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference)SetVbrQuality(val *string) {
	if err := j.validateSetVbrQualityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vbrQuality",
		val,
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) ResetBitrate() {
	_jsii_.InvokeVoid(
		m,
		"resetBitrate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) ResetCodingMode() {
	_jsii_.InvokeVoid(
		m,
		"resetCodingMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) ResetInputType() {
	_jsii_.InvokeVoid(
		m,
		"resetInputType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) ResetProfile() {
	_jsii_.InvokeVoid(
		m,
		"resetProfile",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) ResetRateControlMode() {
	_jsii_.InvokeVoid(
		m,
		"resetRateControlMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) ResetRawFormat() {
	_jsii_.InvokeVoid(
		m,
		"resetRawFormat",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) ResetSampleRate() {
	_jsii_.InvokeVoid(
		m,
		"resetSampleRate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) ResetSpec() {
	_jsii_.InvokeVoid(
		m,
		"resetSpec",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) ResetVbrQuality() {
	_jsii_.InvokeVoid(
		m,
		"resetVbrQuality",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

