package medialivechannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v15/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v15/medialivechannel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference interface {
	cdktf.ComplexObject
	AdaptiveQuantization() *string
	SetAdaptiveQuantization(val *string)
	AdaptiveQuantizationInput() *string
	AfdSignaling() *string
	SetAfdSignaling(val *string)
	AfdSignalingInput() *string
	AlternativeTransferFunction() *string
	SetAlternativeTransferFunction(val *string)
	AlternativeTransferFunctionInput() *string
	Bitrate() *float64
	SetBitrate(val *float64)
	BitrateInput() *float64
	BufSize() *float64
	SetBufSize(val *float64)
	BufSizeInput() *float64
	ColorMetadata() *string
	SetColorMetadata(val *string)
	ColorMetadataInput() *string
	ColorSpaceSettings() MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsOutputReference
	ColorSpaceSettingsInput() *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettings
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
	FilterSettings() MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettingsOutputReference
	FilterSettingsInput() *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettings
	FixedAfd() *string
	SetFixedAfd(val *string)
	FixedAfdInput() *string
	FlickerAq() *string
	SetFlickerAq(val *string)
	FlickerAqInput() *string
	// Experimental.
	Fqn() *string
	FramerateDenominator() *float64
	SetFramerateDenominator(val *float64)
	FramerateDenominatorInput() *float64
	FramerateNumerator() *float64
	SetFramerateNumerator(val *float64)
	FramerateNumeratorInput() *float64
	GopClosedCadence() *float64
	SetGopClosedCadence(val *float64)
	GopClosedCadenceInput() *float64
	GopSize() *float64
	SetGopSize(val *float64)
	GopSizeInput() *float64
	GopSizeUnits() *string
	SetGopSizeUnits(val *string)
	GopSizeUnitsInput() *string
	InternalValue() *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265Settings
	SetInternalValue(val *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265Settings)
	Level() *string
	SetLevel(val *string)
	LevelInput() *string
	LookAheadRateControl() *string
	SetLookAheadRateControl(val *string)
	LookAheadRateControlInput() *string
	MaxBitrate() *float64
	SetMaxBitrate(val *float64)
	MaxBitrateInput() *float64
	MinIInterval() *float64
	SetMinIInterval(val *float64)
	MinIIntervalInput() *float64
	ParDenominator() *float64
	SetParDenominator(val *float64)
	ParDenominatorInput() *float64
	ParNumerator() *float64
	SetParNumerator(val *float64)
	ParNumeratorInput() *float64
	Profile() *string
	SetProfile(val *string)
	ProfileInput() *string
	QvbrQualityLevel() *float64
	SetQvbrQualityLevel(val *float64)
	QvbrQualityLevelInput() *float64
	RateControlMode() *string
	SetRateControlMode(val *string)
	RateControlModeInput() *string
	ScanType() *string
	SetScanType(val *string)
	ScanTypeInput() *string
	SceneChangeDetect() *string
	SetSceneChangeDetect(val *string)
	SceneChangeDetectInput() *string
	Slices() *float64
	SetSlices(val *float64)
	SlicesInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Tier() *string
	SetTier(val *string)
	TierInput() *string
	TimecodeBurninSettings() MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsTimecodeBurninSettingsOutputReference
	TimecodeBurninSettingsInput() *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsTimecodeBurninSettings
	TimecodeInsertion() *string
	SetTimecodeInsertion(val *string)
	TimecodeInsertionInput() *string
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
	PutColorSpaceSettings(value *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettings)
	PutFilterSettings(value *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettings)
	PutTimecodeBurninSettings(value *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsTimecodeBurninSettings)
	ResetAdaptiveQuantization()
	ResetAfdSignaling()
	ResetAlternativeTransferFunction()
	ResetBufSize()
	ResetColorMetadata()
	ResetColorSpaceSettings()
	ResetFilterSettings()
	ResetFixedAfd()
	ResetFlickerAq()
	ResetGopClosedCadence()
	ResetGopSize()
	ResetGopSizeUnits()
	ResetLevel()
	ResetLookAheadRateControl()
	ResetMaxBitrate()
	ResetMinIInterval()
	ResetParDenominator()
	ResetParNumerator()
	ResetProfile()
	ResetQvbrQualityLevel()
	ResetRateControlMode()
	ResetScanType()
	ResetSceneChangeDetect()
	ResetSlices()
	ResetTier()
	ResetTimecodeBurninSettings()
	ResetTimecodeInsertion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference
type jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) AdaptiveQuantization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adaptiveQuantization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) AdaptiveQuantizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adaptiveQuantizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) AfdSignaling() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afdSignaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) AfdSignalingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afdSignalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) AlternativeTransferFunction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alternativeTransferFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) AlternativeTransferFunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alternativeTransferFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) Bitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) BitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) BufSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) BufSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ColorMetadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ColorMetadataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ColorSpaceSettings() MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsOutputReference
	_jsii_.Get(
		j,
		"colorSpaceSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ColorSpaceSettingsInput() *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettings {
	var returns *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettings
	_jsii_.Get(
		j,
		"colorSpaceSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) FilterSettings() MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettingsOutputReference
	_jsii_.Get(
		j,
		"filterSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) FilterSettingsInput() *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettings {
	var returns *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettings
	_jsii_.Get(
		j,
		"filterSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) FixedAfd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fixedAfd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) FixedAfdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fixedAfdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) FlickerAq() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flickerAq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) FlickerAqInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flickerAqInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) FramerateDenominator() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"framerateDenominator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) FramerateDenominatorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"framerateDenominatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) FramerateNumerator() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"framerateNumerator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) FramerateNumeratorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"framerateNumeratorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) GopClosedCadence() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gopClosedCadence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) GopClosedCadenceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gopClosedCadenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) GopSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gopSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) GopSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gopSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) GopSizeUnits() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gopSizeUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) GopSizeUnitsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gopSizeUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) InternalValue() *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265Settings {
	var returns *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265Settings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) LevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"levelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) LookAheadRateControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookAheadRateControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) LookAheadRateControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookAheadRateControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) MaxBitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) MaxBitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) MinIInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) MinIIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ParDenominator() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parDenominator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ParDenominatorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parDenominatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ParNumerator() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parNumerator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ParNumeratorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parNumeratorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) QvbrQualityLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"qvbrQualityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) QvbrQualityLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"qvbrQualityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) RateControlMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateControlMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) RateControlModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateControlModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ScanType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scanType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ScanTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scanTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) SceneChangeDetect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sceneChangeDetect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) SceneChangeDetectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sceneChangeDetectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) Slices() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) SlicesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) Tier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) TierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) TimecodeBurninSettings() MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsTimecodeBurninSettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsTimecodeBurninSettingsOutputReference
	_jsii_.Get(
		j,
		"timecodeBurninSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) TimecodeBurninSettingsInput() *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsTimecodeBurninSettings {
	var returns *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsTimecodeBurninSettings
	_jsii_.Get(
		j,
		"timecodeBurninSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) TimecodeInsertion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timecodeInsertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) TimecodeInsertionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timecodeInsertionInput",
		&returns,
	)
	return returns
}


func NewMedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference_Override(m MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetAdaptiveQuantization(val *string) {
	if err := j.validateSetAdaptiveQuantizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adaptiveQuantization",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetAfdSignaling(val *string) {
	if err := j.validateSetAfdSignalingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"afdSignaling",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetAlternativeTransferFunction(val *string) {
	if err := j.validateSetAlternativeTransferFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alternativeTransferFunction",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetBitrate(val *float64) {
	if err := j.validateSetBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitrate",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetBufSize(val *float64) {
	if err := j.validateSetBufSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufSize",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetColorMetadata(val *string) {
	if err := j.validateSetColorMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"colorMetadata",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetFixedAfd(val *string) {
	if err := j.validateSetFixedAfdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fixedAfd",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetFlickerAq(val *string) {
	if err := j.validateSetFlickerAqParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flickerAq",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetFramerateDenominator(val *float64) {
	if err := j.validateSetFramerateDenominatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"framerateDenominator",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetFramerateNumerator(val *float64) {
	if err := j.validateSetFramerateNumeratorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"framerateNumerator",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetGopClosedCadence(val *float64) {
	if err := j.validateSetGopClosedCadenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gopClosedCadence",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetGopSize(val *float64) {
	if err := j.validateSetGopSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gopSize",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetGopSizeUnits(val *string) {
	if err := j.validateSetGopSizeUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gopSizeUnits",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetInternalValue(val *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265Settings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetLookAheadRateControl(val *string) {
	if err := j.validateSetLookAheadRateControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lookAheadRateControl",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetMaxBitrate(val *float64) {
	if err := j.validateSetMaxBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBitrate",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetMinIInterval(val *float64) {
	if err := j.validateSetMinIIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minIInterval",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetParDenominator(val *float64) {
	if err := j.validateSetParDenominatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parDenominator",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetParNumerator(val *float64) {
	if err := j.validateSetParNumeratorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parNumerator",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetProfile(val *string) {
	if err := j.validateSetProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profile",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetQvbrQualityLevel(val *float64) {
	if err := j.validateSetQvbrQualityLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qvbrQualityLevel",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetRateControlMode(val *string) {
	if err := j.validateSetRateControlModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rateControlMode",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetScanType(val *string) {
	if err := j.validateSetScanTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scanType",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetSceneChangeDetect(val *string) {
	if err := j.validateSetSceneChangeDetectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sceneChangeDetect",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetSlices(val *float64) {
	if err := j.validateSetSlicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slices",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetTier(val *string) {
	if err := j.validateSetTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tier",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference)SetTimecodeInsertion(val *string) {
	if err := j.validateSetTimecodeInsertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timecodeInsertion",
		val,
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) PutColorSpaceSettings(value *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettings) {
	if err := m.validatePutColorSpaceSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putColorSpaceSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) PutFilterSettings(value *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettings) {
	if err := m.validatePutFilterSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFilterSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) PutTimecodeBurninSettings(value *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsTimecodeBurninSettings) {
	if err := m.validatePutTimecodeBurninSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimecodeBurninSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetAdaptiveQuantization() {
	_jsii_.InvokeVoid(
		m,
		"resetAdaptiveQuantization",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetAfdSignaling() {
	_jsii_.InvokeVoid(
		m,
		"resetAfdSignaling",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetAlternativeTransferFunction() {
	_jsii_.InvokeVoid(
		m,
		"resetAlternativeTransferFunction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetBufSize() {
	_jsii_.InvokeVoid(
		m,
		"resetBufSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetColorMetadata() {
	_jsii_.InvokeVoid(
		m,
		"resetColorMetadata",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetColorSpaceSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetColorSpaceSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetFilterSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetFilterSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetFixedAfd() {
	_jsii_.InvokeVoid(
		m,
		"resetFixedAfd",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetFlickerAq() {
	_jsii_.InvokeVoid(
		m,
		"resetFlickerAq",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetGopClosedCadence() {
	_jsii_.InvokeVoid(
		m,
		"resetGopClosedCadence",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetGopSize() {
	_jsii_.InvokeVoid(
		m,
		"resetGopSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetGopSizeUnits() {
	_jsii_.InvokeVoid(
		m,
		"resetGopSizeUnits",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetLevel() {
	_jsii_.InvokeVoid(
		m,
		"resetLevel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetLookAheadRateControl() {
	_jsii_.InvokeVoid(
		m,
		"resetLookAheadRateControl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetMaxBitrate() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxBitrate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetMinIInterval() {
	_jsii_.InvokeVoid(
		m,
		"resetMinIInterval",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetParDenominator() {
	_jsii_.InvokeVoid(
		m,
		"resetParDenominator",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetParNumerator() {
	_jsii_.InvokeVoid(
		m,
		"resetParNumerator",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetProfile() {
	_jsii_.InvokeVoid(
		m,
		"resetProfile",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetQvbrQualityLevel() {
	_jsii_.InvokeVoid(
		m,
		"resetQvbrQualityLevel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetRateControlMode() {
	_jsii_.InvokeVoid(
		m,
		"resetRateControlMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetScanType() {
	_jsii_.InvokeVoid(
		m,
		"resetScanType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetSceneChangeDetect() {
	_jsii_.InvokeVoid(
		m,
		"resetSceneChangeDetect",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetSlices() {
	_jsii_.InvokeVoid(
		m,
		"resetSlices",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetTier() {
	_jsii_.InvokeVoid(
		m,
		"resetTier",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetTimecodeBurninSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetTimecodeBurninSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ResetTimecodeInsertion() {
	_jsii_.InvokeVoid(
		m,
		"resetTimecodeInsertion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

