package medialivechannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/medialivechannel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference interface {
	cdktf.ComplexObject
	AdaptiveQuantization() *string
	SetAdaptiveQuantization(val *string)
	AdaptiveQuantizationInput() *string
	AfdSignaling() *string
	SetAfdSignaling(val *string)
	AfdSignalingInput() *string
	Bitrate() *float64
	SetBitrate(val *float64)
	BitrateInput() *float64
	BufFillPct() *float64
	SetBufFillPct(val *float64)
	BufFillPctInput() *float64
	BufSize() *float64
	SetBufSize(val *float64)
	BufSizeInput() *float64
	ColorMetadata() *string
	SetColorMetadata(val *string)
	ColorMetadataInput() *string
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
	EntropyEncoding() *string
	SetEntropyEncoding(val *string)
	EntropyEncodingInput() *string
	FilterSettings() MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsFilterSettingsOutputReference
	FilterSettingsInput() *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsFilterSettings
	FixedAfd() *string
	SetFixedAfd(val *string)
	FixedAfdInput() *string
	FlickerAq() *string
	SetFlickerAq(val *string)
	FlickerAqInput() *string
	ForceFieldPictures() *string
	SetForceFieldPictures(val *string)
	ForceFieldPicturesInput() *string
	// Experimental.
	Fqn() *string
	FramerateControl() *string
	SetFramerateControl(val *string)
	FramerateControlInput() *string
	FramerateDenominator() *float64
	SetFramerateDenominator(val *float64)
	FramerateDenominatorInput() *float64
	FramerateNumerator() *float64
	SetFramerateNumerator(val *float64)
	FramerateNumeratorInput() *float64
	GopBReference() *string
	SetGopBReference(val *string)
	GopBReferenceInput() *string
	GopClosedCadence() *float64
	SetGopClosedCadence(val *float64)
	GopClosedCadenceInput() *float64
	GopNumBFrames() *float64
	SetGopNumBFrames(val *float64)
	GopNumBFramesInput() *float64
	GopSize() *float64
	SetGopSize(val *float64)
	GopSizeInput() *float64
	GopSizeUnits() *string
	SetGopSizeUnits(val *string)
	GopSizeUnitsInput() *string
	InternalValue() *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264Settings
	SetInternalValue(val *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264Settings)
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
	NumRefFrames() *float64
	SetNumRefFrames(val *float64)
	NumRefFramesInput() *float64
	ParControl() *string
	SetParControl(val *string)
	ParControlInput() *string
	ParDenominator() *float64
	SetParDenominator(val *float64)
	ParDenominatorInput() *float64
	ParNumerator() *float64
	SetParNumerator(val *float64)
	ParNumeratorInput() *float64
	Profile() *string
	SetProfile(val *string)
	ProfileInput() *string
	QualityLevel() *string
	SetQualityLevel(val *string)
	QualityLevelInput() *string
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
	Softness() *float64
	SetSoftness(val *float64)
	SoftnessInput() *float64
	SpatialAq() *string
	SetSpatialAq(val *string)
	SpatialAqInput() *string
	SubgopLength() *string
	SetSubgopLength(val *string)
	SubgopLengthInput() *string
	Syntax() *string
	SetSyntax(val *string)
	SyntaxInput() *string
	TemporalAq() *string
	SetTemporalAq(val *string)
	TemporalAqInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutFilterSettings(value *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsFilterSettings)
	ResetAdaptiveQuantization()
	ResetAfdSignaling()
	ResetBitrate()
	ResetBufFillPct()
	ResetBufSize()
	ResetColorMetadata()
	ResetEntropyEncoding()
	ResetFilterSettings()
	ResetFixedAfd()
	ResetFlickerAq()
	ResetForceFieldPictures()
	ResetFramerateControl()
	ResetFramerateDenominator()
	ResetFramerateNumerator()
	ResetGopBReference()
	ResetGopClosedCadence()
	ResetGopNumBFrames()
	ResetGopSize()
	ResetGopSizeUnits()
	ResetLevel()
	ResetLookAheadRateControl()
	ResetMaxBitrate()
	ResetMinIInterval()
	ResetNumRefFrames()
	ResetParControl()
	ResetParDenominator()
	ResetParNumerator()
	ResetProfile()
	ResetQualityLevel()
	ResetQvbrQualityLevel()
	ResetRateControlMode()
	ResetScanType()
	ResetSceneChangeDetect()
	ResetSlices()
	ResetSoftness()
	ResetSpatialAq()
	ResetSubgopLength()
	ResetSyntax()
	ResetTemporalAq()
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

// The jsii proxy struct for MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference
type jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) AdaptiveQuantization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adaptiveQuantization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) AdaptiveQuantizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adaptiveQuantizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) AfdSignaling() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afdSignaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) AfdSignalingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afdSignalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) Bitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) BitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) BufFillPct() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufFillPct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) BufFillPctInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufFillPctInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) BufSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) BufSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ColorMetadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ColorMetadataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"colorMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) EntropyEncoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entropyEncoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) EntropyEncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entropyEncodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) FilterSettings() MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsFilterSettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsFilterSettingsOutputReference
	_jsii_.Get(
		j,
		"filterSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) FilterSettingsInput() *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsFilterSettings {
	var returns *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsFilterSettings
	_jsii_.Get(
		j,
		"filterSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) FixedAfd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fixedAfd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) FixedAfdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fixedAfdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) FlickerAq() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flickerAq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) FlickerAqInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flickerAqInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ForceFieldPictures() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forceFieldPictures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ForceFieldPicturesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forceFieldPicturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) FramerateControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"framerateControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) FramerateControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"framerateControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) FramerateDenominator() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"framerateDenominator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) FramerateDenominatorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"framerateDenominatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) FramerateNumerator() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"framerateNumerator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) FramerateNumeratorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"framerateNumeratorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) GopBReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gopBReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) GopBReferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gopBReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) GopClosedCadence() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gopClosedCadence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) GopClosedCadenceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gopClosedCadenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) GopNumBFrames() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gopNumBFrames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) GopNumBFramesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gopNumBFramesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) GopSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gopSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) GopSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gopSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) GopSizeUnits() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gopSizeUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) GopSizeUnitsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gopSizeUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) InternalValue() *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264Settings {
	var returns *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264Settings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) LevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"levelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) LookAheadRateControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookAheadRateControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) LookAheadRateControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookAheadRateControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) MaxBitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) MaxBitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) MinIInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) MinIIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) NumRefFrames() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numRefFrames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) NumRefFramesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numRefFramesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ParControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ParControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ParDenominator() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parDenominator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ParDenominatorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parDenominatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ParNumerator() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parNumerator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ParNumeratorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parNumeratorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) QualityLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) QualityLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) QvbrQualityLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"qvbrQualityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) QvbrQualityLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"qvbrQualityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) RateControlMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateControlMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) RateControlModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateControlModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ScanType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scanType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ScanTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scanTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) SceneChangeDetect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sceneChangeDetect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) SceneChangeDetectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sceneChangeDetectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) Slices() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) SlicesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) Softness() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"softness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) SoftnessInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"softnessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) SpatialAq() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spatialAq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) SpatialAqInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spatialAqInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) SubgopLength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subgopLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) SubgopLengthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subgopLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) Syntax() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syntax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) SyntaxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syntaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) TemporalAq() *string {
	var returns *string
	_jsii_.Get(
		j,
		"temporalAq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) TemporalAqInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"temporalAqInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) TimecodeInsertion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timecodeInsertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) TimecodeInsertionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timecodeInsertionInput",
		&returns,
	)
	return returns
}


func NewMedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference_Override(m MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetAdaptiveQuantization(val *string) {
	if err := j.validateSetAdaptiveQuantizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adaptiveQuantization",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetAfdSignaling(val *string) {
	if err := j.validateSetAfdSignalingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"afdSignaling",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetBitrate(val *float64) {
	if err := j.validateSetBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitrate",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetBufFillPct(val *float64) {
	if err := j.validateSetBufFillPctParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufFillPct",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetBufSize(val *float64) {
	if err := j.validateSetBufSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufSize",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetColorMetadata(val *string) {
	if err := j.validateSetColorMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"colorMetadata",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetEntropyEncoding(val *string) {
	if err := j.validateSetEntropyEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entropyEncoding",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetFixedAfd(val *string) {
	if err := j.validateSetFixedAfdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fixedAfd",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetFlickerAq(val *string) {
	if err := j.validateSetFlickerAqParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flickerAq",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetForceFieldPictures(val *string) {
	if err := j.validateSetForceFieldPicturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceFieldPictures",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetFramerateControl(val *string) {
	if err := j.validateSetFramerateControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"framerateControl",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetFramerateDenominator(val *float64) {
	if err := j.validateSetFramerateDenominatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"framerateDenominator",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetFramerateNumerator(val *float64) {
	if err := j.validateSetFramerateNumeratorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"framerateNumerator",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetGopBReference(val *string) {
	if err := j.validateSetGopBReferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gopBReference",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetGopClosedCadence(val *float64) {
	if err := j.validateSetGopClosedCadenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gopClosedCadence",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetGopNumBFrames(val *float64) {
	if err := j.validateSetGopNumBFramesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gopNumBFrames",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetGopSize(val *float64) {
	if err := j.validateSetGopSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gopSize",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetGopSizeUnits(val *string) {
	if err := j.validateSetGopSizeUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gopSizeUnits",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetInternalValue(val *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264Settings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetLookAheadRateControl(val *string) {
	if err := j.validateSetLookAheadRateControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lookAheadRateControl",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetMaxBitrate(val *float64) {
	if err := j.validateSetMaxBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBitrate",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetMinIInterval(val *float64) {
	if err := j.validateSetMinIIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minIInterval",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetNumRefFrames(val *float64) {
	if err := j.validateSetNumRefFramesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numRefFrames",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetParControl(val *string) {
	if err := j.validateSetParControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parControl",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetParDenominator(val *float64) {
	if err := j.validateSetParDenominatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parDenominator",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetParNumerator(val *float64) {
	if err := j.validateSetParNumeratorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parNumerator",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetProfile(val *string) {
	if err := j.validateSetProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profile",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetQualityLevel(val *string) {
	if err := j.validateSetQualityLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qualityLevel",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetQvbrQualityLevel(val *float64) {
	if err := j.validateSetQvbrQualityLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qvbrQualityLevel",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetRateControlMode(val *string) {
	if err := j.validateSetRateControlModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rateControlMode",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetScanType(val *string) {
	if err := j.validateSetScanTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scanType",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetSceneChangeDetect(val *string) {
	if err := j.validateSetSceneChangeDetectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sceneChangeDetect",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetSlices(val *float64) {
	if err := j.validateSetSlicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slices",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetSoftness(val *float64) {
	if err := j.validateSetSoftnessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"softness",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetSpatialAq(val *string) {
	if err := j.validateSetSpatialAqParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spatialAq",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetSubgopLength(val *string) {
	if err := j.validateSetSubgopLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subgopLength",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetSyntax(val *string) {
	if err := j.validateSetSyntaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syntax",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetTemporalAq(val *string) {
	if err := j.validateSetTemporalAqParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"temporalAq",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference)SetTimecodeInsertion(val *string) {
	if err := j.validateSetTimecodeInsertionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timecodeInsertion",
		val,
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) PutFilterSettings(value *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsFilterSettings) {
	if err := m.validatePutFilterSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putFilterSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetAdaptiveQuantization() {
	_jsii_.InvokeVoid(
		m,
		"resetAdaptiveQuantization",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetAfdSignaling() {
	_jsii_.InvokeVoid(
		m,
		"resetAfdSignaling",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetBitrate() {
	_jsii_.InvokeVoid(
		m,
		"resetBitrate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetBufFillPct() {
	_jsii_.InvokeVoid(
		m,
		"resetBufFillPct",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetBufSize() {
	_jsii_.InvokeVoid(
		m,
		"resetBufSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetColorMetadata() {
	_jsii_.InvokeVoid(
		m,
		"resetColorMetadata",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetEntropyEncoding() {
	_jsii_.InvokeVoid(
		m,
		"resetEntropyEncoding",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetFilterSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetFilterSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetFixedAfd() {
	_jsii_.InvokeVoid(
		m,
		"resetFixedAfd",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetFlickerAq() {
	_jsii_.InvokeVoid(
		m,
		"resetFlickerAq",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetForceFieldPictures() {
	_jsii_.InvokeVoid(
		m,
		"resetForceFieldPictures",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetFramerateControl() {
	_jsii_.InvokeVoid(
		m,
		"resetFramerateControl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetFramerateDenominator() {
	_jsii_.InvokeVoid(
		m,
		"resetFramerateDenominator",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetFramerateNumerator() {
	_jsii_.InvokeVoid(
		m,
		"resetFramerateNumerator",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetGopBReference() {
	_jsii_.InvokeVoid(
		m,
		"resetGopBReference",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetGopClosedCadence() {
	_jsii_.InvokeVoid(
		m,
		"resetGopClosedCadence",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetGopNumBFrames() {
	_jsii_.InvokeVoid(
		m,
		"resetGopNumBFrames",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetGopSize() {
	_jsii_.InvokeVoid(
		m,
		"resetGopSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetGopSizeUnits() {
	_jsii_.InvokeVoid(
		m,
		"resetGopSizeUnits",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetLevel() {
	_jsii_.InvokeVoid(
		m,
		"resetLevel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetLookAheadRateControl() {
	_jsii_.InvokeVoid(
		m,
		"resetLookAheadRateControl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetMaxBitrate() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxBitrate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetMinIInterval() {
	_jsii_.InvokeVoid(
		m,
		"resetMinIInterval",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetNumRefFrames() {
	_jsii_.InvokeVoid(
		m,
		"resetNumRefFrames",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetParControl() {
	_jsii_.InvokeVoid(
		m,
		"resetParControl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetParDenominator() {
	_jsii_.InvokeVoid(
		m,
		"resetParDenominator",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetParNumerator() {
	_jsii_.InvokeVoid(
		m,
		"resetParNumerator",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetProfile() {
	_jsii_.InvokeVoid(
		m,
		"resetProfile",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetQualityLevel() {
	_jsii_.InvokeVoid(
		m,
		"resetQualityLevel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetQvbrQualityLevel() {
	_jsii_.InvokeVoid(
		m,
		"resetQvbrQualityLevel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetRateControlMode() {
	_jsii_.InvokeVoid(
		m,
		"resetRateControlMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetScanType() {
	_jsii_.InvokeVoid(
		m,
		"resetScanType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetSceneChangeDetect() {
	_jsii_.InvokeVoid(
		m,
		"resetSceneChangeDetect",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetSlices() {
	_jsii_.InvokeVoid(
		m,
		"resetSlices",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetSoftness() {
	_jsii_.InvokeVoid(
		m,
		"resetSoftness",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetSpatialAq() {
	_jsii_.InvokeVoid(
		m,
		"resetSpatialAq",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetSubgopLength() {
	_jsii_.InvokeVoid(
		m,
		"resetSubgopLength",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetSyntax() {
	_jsii_.InvokeVoid(
		m,
		"resetSyntax",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetTemporalAq() {
	_jsii_.InvokeVoid(
		m,
		"resetTemporalAq",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ResetTimecodeInsertion() {
	_jsii_.InvokeVoid(
		m,
		"resetTimecodeInsertion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH264SettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

