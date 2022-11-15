package medialivechannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v10/medialivechannel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference interface {
	cdktf.ComplexObject
	AttenuationControl() *string
	SetAttenuationControl(val *string)
	AttenuationControlInput() *string
	Bitrate() *float64
	SetBitrate(val *float64)
	BitrateInput() *float64
	BitstreamMode() *string
	SetBitstreamMode(val *string)
	BitstreamModeInput() *string
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
	DcFilter() *string
	SetDcFilter(val *string)
	DcFilterInput() *string
	Dialnorm() *float64
	SetDialnorm(val *float64)
	DialnormInput() *float64
	DrcLine() *string
	SetDrcLine(val *string)
	DrcLineInput() *string
	DrcRf() *string
	SetDrcRf(val *string)
	DrcRfInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3Settings
	SetInternalValue(val *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3Settings)
	LfeControl() *string
	SetLfeControl(val *string)
	LfeControlInput() *string
	LfeFilter() *string
	SetLfeFilter(val *string)
	LfeFilterInput() *string
	LoRoCenterMixLevel() *float64
	SetLoRoCenterMixLevel(val *float64)
	LoRoCenterMixLevelInput() *float64
	LoRoSurroundMixLevel() *float64
	SetLoRoSurroundMixLevel(val *float64)
	LoRoSurroundMixLevelInput() *float64
	LtRtCenterMixLevel() *float64
	SetLtRtCenterMixLevel(val *float64)
	LtRtCenterMixLevelInput() *float64
	LtRtSurroundMixLevel() *float64
	SetLtRtSurroundMixLevel(val *float64)
	LtRtSurroundMixLevelInput() *float64
	MetadataControl() *string
	SetMetadataControl(val *string)
	MetadataControlInput() *string
	PassthroughControl() *string
	SetPassthroughControl(val *string)
	PassthroughControlInput() *string
	PhaseControl() *string
	SetPhaseControl(val *string)
	PhaseControlInput() *string
	StereoDownmix() *string
	SetStereoDownmix(val *string)
	StereoDownmixInput() *string
	SurroundExMode() *string
	SetSurroundExMode(val *string)
	SurroundExModeInput() *string
	SurroundMode() *string
	SetSurroundMode(val *string)
	SurroundModeInput() *string
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
	ResetAttenuationControl()
	ResetBitrate()
	ResetBitstreamMode()
	ResetCodingMode()
	ResetDcFilter()
	ResetDialnorm()
	ResetDrcLine()
	ResetDrcRf()
	ResetLfeControl()
	ResetLfeFilter()
	ResetLoRoCenterMixLevel()
	ResetLoRoSurroundMixLevel()
	ResetLtRtCenterMixLevel()
	ResetLtRtSurroundMixLevel()
	ResetMetadataControl()
	ResetPassthroughControl()
	ResetPhaseControl()
	ResetStereoDownmix()
	ResetSurroundExMode()
	ResetSurroundMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference
type jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) AttenuationControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attenuationControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) AttenuationControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attenuationControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) Bitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) BitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) BitstreamMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitstreamMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) BitstreamModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitstreamModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) CodingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) CodingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) DcFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dcFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) DcFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dcFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) Dialnorm() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dialnorm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) DialnormInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dialnormInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) DrcLine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drcLine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) DrcLineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drcLineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) DrcRf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drcRf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) DrcRfInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drcRfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) InternalValue() *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3Settings {
	var returns *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3Settings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) LfeControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lfeControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) LfeControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lfeControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) LfeFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lfeFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) LfeFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lfeFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) LoRoCenterMixLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loRoCenterMixLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) LoRoCenterMixLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loRoCenterMixLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) LoRoSurroundMixLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loRoSurroundMixLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) LoRoSurroundMixLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"loRoSurroundMixLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) LtRtCenterMixLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ltRtCenterMixLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) LtRtCenterMixLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ltRtCenterMixLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) LtRtSurroundMixLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ltRtSurroundMixLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) LtRtSurroundMixLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ltRtSurroundMixLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) MetadataControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) MetadataControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) PassthroughControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passthroughControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) PassthroughControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passthroughControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) PhaseControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phaseControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) PhaseControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phaseControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) StereoDownmix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stereoDownmix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) StereoDownmixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stereoDownmixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) SurroundExMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"surroundExMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) SurroundExModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"surroundExModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) SurroundMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"surroundMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) SurroundModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"surroundModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference_Override(m MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetAttenuationControl(val *string) {
	if err := j.validateSetAttenuationControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attenuationControl",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetBitrate(val *float64) {
	if err := j.validateSetBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitrate",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetBitstreamMode(val *string) {
	if err := j.validateSetBitstreamModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitstreamMode",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetCodingMode(val *string) {
	if err := j.validateSetCodingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codingMode",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetDcFilter(val *string) {
	if err := j.validateSetDcFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dcFilter",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetDialnorm(val *float64) {
	if err := j.validateSetDialnormParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dialnorm",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetDrcLine(val *string) {
	if err := j.validateSetDrcLineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drcLine",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetDrcRf(val *string) {
	if err := j.validateSetDrcRfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drcRf",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetInternalValue(val *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3Settings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetLfeControl(val *string) {
	if err := j.validateSetLfeControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lfeControl",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetLfeFilter(val *string) {
	if err := j.validateSetLfeFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lfeFilter",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetLoRoCenterMixLevel(val *float64) {
	if err := j.validateSetLoRoCenterMixLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loRoCenterMixLevel",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetLoRoSurroundMixLevel(val *float64) {
	if err := j.validateSetLoRoSurroundMixLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loRoSurroundMixLevel",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetLtRtCenterMixLevel(val *float64) {
	if err := j.validateSetLtRtCenterMixLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ltRtCenterMixLevel",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetLtRtSurroundMixLevel(val *float64) {
	if err := j.validateSetLtRtSurroundMixLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ltRtSurroundMixLevel",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetMetadataControl(val *string) {
	if err := j.validateSetMetadataControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataControl",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetPassthroughControl(val *string) {
	if err := j.validateSetPassthroughControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passthroughControl",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetPhaseControl(val *string) {
	if err := j.validateSetPhaseControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phaseControl",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetStereoDownmix(val *string) {
	if err := j.validateSetStereoDownmixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stereoDownmix",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetSurroundExMode(val *string) {
	if err := j.validateSetSurroundExModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"surroundExMode",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetSurroundMode(val *string) {
	if err := j.validateSetSurroundModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"surroundMode",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) ResetAttenuationControl() {
	_jsii_.InvokeVoid(
		m,
		"resetAttenuationControl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) ResetBitrate() {
	_jsii_.InvokeVoid(
		m,
		"resetBitrate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) ResetBitstreamMode() {
	_jsii_.InvokeVoid(
		m,
		"resetBitstreamMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) ResetCodingMode() {
	_jsii_.InvokeVoid(
		m,
		"resetCodingMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) ResetDcFilter() {
	_jsii_.InvokeVoid(
		m,
		"resetDcFilter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) ResetDialnorm() {
	_jsii_.InvokeVoid(
		m,
		"resetDialnorm",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) ResetDrcLine() {
	_jsii_.InvokeVoid(
		m,
		"resetDrcLine",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) ResetDrcRf() {
	_jsii_.InvokeVoid(
		m,
		"resetDrcRf",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) ResetLfeControl() {
	_jsii_.InvokeVoid(
		m,
		"resetLfeControl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) ResetLfeFilter() {
	_jsii_.InvokeVoid(
		m,
		"resetLfeFilter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) ResetLoRoCenterMixLevel() {
	_jsii_.InvokeVoid(
		m,
		"resetLoRoCenterMixLevel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) ResetLoRoSurroundMixLevel() {
	_jsii_.InvokeVoid(
		m,
		"resetLoRoSurroundMixLevel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) ResetLtRtCenterMixLevel() {
	_jsii_.InvokeVoid(
		m,
		"resetLtRtCenterMixLevel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) ResetLtRtSurroundMixLevel() {
	_jsii_.InvokeVoid(
		m,
		"resetLtRtSurroundMixLevel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) ResetMetadataControl() {
	_jsii_.InvokeVoid(
		m,
		"resetMetadataControl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) ResetPassthroughControl() {
	_jsii_.InvokeVoid(
		m,
		"resetPassthroughControl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) ResetPhaseControl() {
	_jsii_.InvokeVoid(
		m,
		"resetPhaseControl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) ResetStereoDownmix() {
	_jsii_.InvokeVoid(
		m,
		"resetStereoDownmix",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) ResetSurroundExMode() {
	_jsii_.InvokeVoid(
		m,
		"resetSurroundExMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) ResetSurroundMode() {
	_jsii_.InvokeVoid(
		m,
		"resetSurroundMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3SettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

