package medialivechannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v15/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v15/medialivechannel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference interface {
	cdktf.ComplexObject
	AacSettings() MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference
	AacSettingsInput() *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettings
	Ac3Settings() MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAc3SettingsOutputReference
	Ac3SettingsInput() *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAc3Settings
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
	Eac3AtmosSettings() MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference
	Eac3AtmosSettingsInput() *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettings
	Eac3Settings() MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference
	Eac3SettingsInput() *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3Settings
	// Experimental.
	Fqn() *string
	InternalValue() *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettings
	SetInternalValue(val *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettings)
	Mp2Settings() MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsMp2SettingsOutputReference
	Mp2SettingsInput() *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsMp2Settings
	PassThroughSettings() MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsPassThroughSettingsOutputReference
	PassThroughSettingsInput() *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsPassThroughSettings
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WavSettings() MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsWavSettingsOutputReference
	WavSettingsInput() *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsWavSettings
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
	PutAacSettings(value *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettings)
	PutAc3Settings(value *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAc3Settings)
	PutEac3AtmosSettings(value *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettings)
	PutEac3Settings(value *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3Settings)
	PutMp2Settings(value *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsMp2Settings)
	PutPassThroughSettings(value *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsPassThroughSettings)
	PutWavSettings(value *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsWavSettings)
	ResetAacSettings()
	ResetAc3Settings()
	ResetEac3AtmosSettings()
	ResetEac3Settings()
	ResetMp2Settings()
	ResetPassThroughSettings()
	ResetWavSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference
type jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) AacSettings() MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettingsOutputReference
	_jsii_.Get(
		j,
		"aacSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) AacSettingsInput() *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettings {
	var returns *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettings
	_jsii_.Get(
		j,
		"aacSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) Ac3Settings() MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAc3SettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAc3SettingsOutputReference
	_jsii_.Get(
		j,
		"ac3Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) Ac3SettingsInput() *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAc3Settings {
	var returns *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAc3Settings
	_jsii_.Get(
		j,
		"ac3SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) Eac3AtmosSettings() MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettingsOutputReference
	_jsii_.Get(
		j,
		"eac3AtmosSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) Eac3AtmosSettingsInput() *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettings {
	var returns *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettings
	_jsii_.Get(
		j,
		"eac3AtmosSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) Eac3Settings() MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference
	_jsii_.Get(
		j,
		"eac3Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) Eac3SettingsInput() *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3Settings {
	var returns *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3Settings
	_jsii_.Get(
		j,
		"eac3SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) InternalValue() *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettings {
	var returns *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) Mp2Settings() MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsMp2SettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsMp2SettingsOutputReference
	_jsii_.Get(
		j,
		"mp2Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) Mp2SettingsInput() *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsMp2Settings {
	var returns *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsMp2Settings
	_jsii_.Get(
		j,
		"mp2SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) PassThroughSettings() MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsPassThroughSettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsPassThroughSettingsOutputReference
	_jsii_.Get(
		j,
		"passThroughSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) PassThroughSettingsInput() *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsPassThroughSettings {
	var returns *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsPassThroughSettings
	_jsii_.Get(
		j,
		"passThroughSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) WavSettings() MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsWavSettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsWavSettingsOutputReference
	_jsii_.Get(
		j,
		"wavSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) WavSettingsInput() *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsWavSettings {
	var returns *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsWavSettings
	_jsii_.Get(
		j,
		"wavSettingsInput",
		&returns,
	)
	return returns
}


func NewMedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference_Override(m MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference)SetInternalValue(val *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) PutAacSettings(value *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettings) {
	if err := m.validatePutAacSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAacSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) PutAc3Settings(value *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAc3Settings) {
	if err := m.validatePutAc3SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAc3Settings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) PutEac3AtmosSettings(value *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettings) {
	if err := m.validatePutEac3AtmosSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEac3AtmosSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) PutEac3Settings(value *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3Settings) {
	if err := m.validatePutEac3SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEac3Settings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) PutMp2Settings(value *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsMp2Settings) {
	if err := m.validatePutMp2SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMp2Settings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) PutPassThroughSettings(value *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsPassThroughSettings) {
	if err := m.validatePutPassThroughSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPassThroughSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) PutWavSettings(value *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsWavSettings) {
	if err := m.validatePutWavSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putWavSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) ResetAacSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetAacSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) ResetAc3Settings() {
	_jsii_.InvokeVoid(
		m,
		"resetAc3Settings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) ResetEac3AtmosSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetEac3AtmosSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) ResetEac3Settings() {
	_jsii_.InvokeVoid(
		m,
		"resetEac3Settings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) ResetMp2Settings() {
	_jsii_.InvokeVoid(
		m,
		"resetMp2Settings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) ResetPassThroughSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetPassThroughSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) ResetWavSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetWavSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

