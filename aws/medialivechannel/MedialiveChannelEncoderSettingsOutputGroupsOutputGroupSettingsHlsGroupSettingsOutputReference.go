package medialivechannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v13/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v13/medialivechannel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference interface {
	cdktf.ComplexObject
	AdMarkers() *[]*string
	SetAdMarkers(val *[]*string)
	AdMarkersInput() *[]*string
	BaseUrlContent() *string
	SetBaseUrlContent(val *string)
	BaseUrlContent1() *string
	SetBaseUrlContent1(val *string)
	BaseUrlContent1Input() *string
	BaseUrlContentInput() *string
	BaseUrlManifest() *string
	SetBaseUrlManifest(val *string)
	BaseUrlManifest1() *string
	SetBaseUrlManifest1(val *string)
	BaseUrlManifest1Input() *string
	BaseUrlManifestInput() *string
	CaptionLanguageMappings() MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsCaptionLanguageMappingsList
	CaptionLanguageMappingsInput() interface{}
	CaptionLanguageSetting() *string
	SetCaptionLanguageSetting(val *string)
	CaptionLanguageSettingInput() *string
	ClientCache() *string
	SetClientCache(val *string)
	ClientCacheInput() *string
	CodecSpecification() *string
	SetCodecSpecification(val *string)
	CodecSpecificationInput() *string
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
	ConstantIv() *string
	SetConstantIv(val *string)
	ConstantIvInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Destination() MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestinationOutputReference
	DestinationInput() *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestination
	DirectoryStructure() *string
	SetDirectoryStructure(val *string)
	DirectoryStructureInput() *string
	DiscontinuityTags() *string
	SetDiscontinuityTags(val *string)
	DiscontinuityTagsInput() *string
	EncryptionType() *string
	SetEncryptionType(val *string)
	EncryptionTypeInput() *string
	// Experimental.
	Fqn() *string
	HlsCdnSettings() MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsHlsCdnSettingsList
	HlsCdnSettingsInput() interface{}
	HlsId3SegmentTagging() *string
	SetHlsId3SegmentTagging(val *string)
	HlsId3SegmentTaggingInput() *string
	IframeOnlyPlaylists() *string
	SetIframeOnlyPlaylists(val *string)
	IframeOnlyPlaylistsInput() *string
	IncompleteSegmentBehavior() *string
	SetIncompleteSegmentBehavior(val *string)
	IncompleteSegmentBehaviorInput() *string
	IndexNSegments() *float64
	SetIndexNSegments(val *float64)
	IndexNSegmentsInput() *float64
	InputLossAction() *string
	SetInputLossAction(val *string)
	InputLossActionInput() *string
	InternalValue() *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettings
	SetInternalValue(val *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettings)
	IvInManifest() *string
	SetIvInManifest(val *string)
	IvInManifestInput() *string
	IvSource() *string
	SetIvSource(val *string)
	IvSourceInput() *string
	KeepSegments() *float64
	SetKeepSegments(val *float64)
	KeepSegmentsInput() *float64
	KeyFormat() *string
	SetKeyFormat(val *string)
	KeyFormatInput() *string
	KeyFormatVersions() *string
	SetKeyFormatVersions(val *string)
	KeyFormatVersionsInput() *string
	KeyProviderSettings() MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsKeyProviderSettingsOutputReference
	KeyProviderSettingsInput() *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsKeyProviderSettings
	ManifestCompression() *string
	SetManifestCompression(val *string)
	ManifestCompressionInput() *string
	ManifestDurationFormat() *string
	SetManifestDurationFormat(val *string)
	ManifestDurationFormatInput() *string
	MinSegmentLength() *float64
	SetMinSegmentLength(val *float64)
	MinSegmentLengthInput() *float64
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	OutputSelection() *string
	SetOutputSelection(val *string)
	OutputSelectionInput() *string
	ProgramDateTime() *string
	SetProgramDateTime(val *string)
	ProgramDateTimeClock() *string
	SetProgramDateTimeClock(val *string)
	ProgramDateTimeClockInput() *string
	ProgramDateTimeInput() *string
	ProgramDateTimePeriod() *float64
	SetProgramDateTimePeriod(val *float64)
	ProgramDateTimePeriodInput() *float64
	RedundantManifest() *string
	SetRedundantManifest(val *string)
	RedundantManifestInput() *string
	SegmentLength() *float64
	SetSegmentLength(val *float64)
	SegmentLengthInput() *float64
	SegmentsPerSubdirectory() *float64
	SetSegmentsPerSubdirectory(val *float64)
	SegmentsPerSubdirectoryInput() *float64
	StreamInfResolution() *string
	SetStreamInfResolution(val *string)
	StreamInfResolutionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimedMetadataId3Frame() *string
	SetTimedMetadataId3Frame(val *string)
	TimedMetadataId3FrameInput() *string
	TimedMetadataId3Period() *float64
	SetTimedMetadataId3Period(val *float64)
	TimedMetadataId3PeriodInput() *float64
	TimestampDeltaMilliseconds() *float64
	SetTimestampDeltaMilliseconds(val *float64)
	TimestampDeltaMillisecondsInput() *float64
	TsFileMode() *string
	SetTsFileMode(val *string)
	TsFileModeInput() *string
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
	PutCaptionLanguageMappings(value interface{})
	PutDestination(value *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestination)
	PutHlsCdnSettings(value interface{})
	PutKeyProviderSettings(value *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsKeyProviderSettings)
	ResetAdMarkers()
	ResetBaseUrlContent()
	ResetBaseUrlContent1()
	ResetBaseUrlManifest()
	ResetBaseUrlManifest1()
	ResetCaptionLanguageMappings()
	ResetCaptionLanguageSetting()
	ResetClientCache()
	ResetCodecSpecification()
	ResetConstantIv()
	ResetDirectoryStructure()
	ResetDiscontinuityTags()
	ResetEncryptionType()
	ResetHlsCdnSettings()
	ResetHlsId3SegmentTagging()
	ResetIframeOnlyPlaylists()
	ResetIncompleteSegmentBehavior()
	ResetIndexNSegments()
	ResetInputLossAction()
	ResetIvInManifest()
	ResetIvSource()
	ResetKeepSegments()
	ResetKeyFormat()
	ResetKeyFormatVersions()
	ResetKeyProviderSettings()
	ResetManifestCompression()
	ResetManifestDurationFormat()
	ResetMinSegmentLength()
	ResetMode()
	ResetOutputSelection()
	ResetProgramDateTime()
	ResetProgramDateTimeClock()
	ResetProgramDateTimePeriod()
	ResetRedundantManifest()
	ResetSegmentLength()
	ResetSegmentsPerSubdirectory()
	ResetStreamInfResolution()
	ResetTimedMetadataId3Frame()
	ResetTimedMetadataId3Period()
	ResetTimestampDeltaMilliseconds()
	ResetTsFileMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference
type jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) AdMarkers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adMarkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) AdMarkersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adMarkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) BaseUrlContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) BaseUrlContent1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlContent1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) BaseUrlContent1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlContent1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) BaseUrlContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) BaseUrlManifest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlManifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) BaseUrlManifest1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlManifest1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) BaseUrlManifest1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlManifest1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) BaseUrlManifestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlManifestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) CaptionLanguageMappings() MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsCaptionLanguageMappingsList {
	var returns MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsCaptionLanguageMappingsList
	_jsii_.Get(
		j,
		"captionLanguageMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) CaptionLanguageMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"captionLanguageMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) CaptionLanguageSetting() *string {
	var returns *string
	_jsii_.Get(
		j,
		"captionLanguageSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) CaptionLanguageSettingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"captionLanguageSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ClientCache() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ClientCacheInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) CodecSpecification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codecSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) CodecSpecificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codecSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ConstantIv() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constantIv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ConstantIvInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constantIvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) Destination() MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestinationOutputReference {
	var returns MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestinationOutputReference
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) DestinationInput() *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestination {
	var returns *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestination
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) DirectoryStructure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryStructure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) DirectoryStructureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryStructureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) DiscontinuityTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discontinuityTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) DiscontinuityTagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discontinuityTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) EncryptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) EncryptionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) HlsCdnSettings() MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsHlsCdnSettingsList {
	var returns MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsHlsCdnSettingsList
	_jsii_.Get(
		j,
		"hlsCdnSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) HlsCdnSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hlsCdnSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) HlsId3SegmentTagging() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hlsId3SegmentTagging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) HlsId3SegmentTaggingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hlsId3SegmentTaggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) IframeOnlyPlaylists() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iframeOnlyPlaylists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) IframeOnlyPlaylistsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iframeOnlyPlaylistsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) IncompleteSegmentBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"incompleteSegmentBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) IncompleteSegmentBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"incompleteSegmentBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) IndexNSegments() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indexNSegments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) IndexNSegmentsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indexNSegmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) InputLossAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputLossAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) InputLossActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputLossActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) InternalValue() *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettings {
	var returns *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) IvInManifest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ivInManifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) IvInManifestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ivInManifestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) IvSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ivSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) IvSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ivSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) KeepSegments() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepSegments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) KeepSegmentsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepSegmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) KeyFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) KeyFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) KeyFormatVersions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyFormatVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) KeyFormatVersionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyFormatVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) KeyProviderSettings() MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsKeyProviderSettingsOutputReference {
	var returns MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsKeyProviderSettingsOutputReference
	_jsii_.Get(
		j,
		"keyProviderSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) KeyProviderSettingsInput() *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsKeyProviderSettings {
	var returns *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsKeyProviderSettings
	_jsii_.Get(
		j,
		"keyProviderSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ManifestCompression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestCompression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ManifestCompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestCompressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ManifestDurationFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestDurationFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ManifestDurationFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestDurationFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) MinSegmentLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSegmentLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) MinSegmentLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSegmentLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) OutputSelection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) OutputSelectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ProgramDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"programDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ProgramDateTimeClock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"programDateTimeClock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ProgramDateTimeClockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"programDateTimeClockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ProgramDateTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"programDateTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ProgramDateTimePeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programDateTimePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ProgramDateTimePeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programDateTimePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) RedundantManifest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redundantManifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) RedundantManifestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redundantManifestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) SegmentLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) SegmentLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) SegmentsPerSubdirectory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentsPerSubdirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) SegmentsPerSubdirectoryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentsPerSubdirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) StreamInfResolution() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamInfResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) StreamInfResolutionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamInfResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) TimedMetadataId3Frame() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timedMetadataId3Frame",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) TimedMetadataId3FrameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timedMetadataId3FrameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) TimedMetadataId3Period() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timedMetadataId3Period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) TimedMetadataId3PeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timedMetadataId3PeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) TimestampDeltaMilliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timestampDeltaMilliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) TimestampDeltaMillisecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timestampDeltaMillisecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) TsFileMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tsFileMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) TsFileModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tsFileModeInput",
		&returns,
	)
	return returns
}


func NewMedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewMedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference_Override(m MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.medialiveChannel.MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetAdMarkers(val *[]*string) {
	if err := j.validateSetAdMarkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adMarkers",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetBaseUrlContent(val *string) {
	if err := j.validateSetBaseUrlContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseUrlContent",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetBaseUrlContent1(val *string) {
	if err := j.validateSetBaseUrlContent1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseUrlContent1",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetBaseUrlManifest(val *string) {
	if err := j.validateSetBaseUrlManifestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseUrlManifest",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetBaseUrlManifest1(val *string) {
	if err := j.validateSetBaseUrlManifest1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseUrlManifest1",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetCaptionLanguageSetting(val *string) {
	if err := j.validateSetCaptionLanguageSettingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"captionLanguageSetting",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetClientCache(val *string) {
	if err := j.validateSetClientCacheParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCache",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetCodecSpecification(val *string) {
	if err := j.validateSetCodecSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codecSpecification",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetConstantIv(val *string) {
	if err := j.validateSetConstantIvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"constantIv",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetDirectoryStructure(val *string) {
	if err := j.validateSetDirectoryStructureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directoryStructure",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetDiscontinuityTags(val *string) {
	if err := j.validateSetDiscontinuityTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"discontinuityTags",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetEncryptionType(val *string) {
	if err := j.validateSetEncryptionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionType",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetHlsId3SegmentTagging(val *string) {
	if err := j.validateSetHlsId3SegmentTaggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hlsId3SegmentTagging",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetIframeOnlyPlaylists(val *string) {
	if err := j.validateSetIframeOnlyPlaylistsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iframeOnlyPlaylists",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetIncompleteSegmentBehavior(val *string) {
	if err := j.validateSetIncompleteSegmentBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"incompleteSegmentBehavior",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetIndexNSegments(val *float64) {
	if err := j.validateSetIndexNSegmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexNSegments",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetInputLossAction(val *string) {
	if err := j.validateSetInputLossActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputLossAction",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetInternalValue(val *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetIvInManifest(val *string) {
	if err := j.validateSetIvInManifestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ivInManifest",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetIvSource(val *string) {
	if err := j.validateSetIvSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ivSource",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetKeepSegments(val *float64) {
	if err := j.validateSetKeepSegmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepSegments",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetKeyFormat(val *string) {
	if err := j.validateSetKeyFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyFormat",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetKeyFormatVersions(val *string) {
	if err := j.validateSetKeyFormatVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyFormatVersions",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetManifestCompression(val *string) {
	if err := j.validateSetManifestCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifestCompression",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetManifestDurationFormat(val *string) {
	if err := j.validateSetManifestDurationFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifestDurationFormat",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetMinSegmentLength(val *float64) {
	if err := j.validateSetMinSegmentLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSegmentLength",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetOutputSelection(val *string) {
	if err := j.validateSetOutputSelectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputSelection",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetProgramDateTime(val *string) {
	if err := j.validateSetProgramDateTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"programDateTime",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetProgramDateTimeClock(val *string) {
	if err := j.validateSetProgramDateTimeClockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"programDateTimeClock",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetProgramDateTimePeriod(val *float64) {
	if err := j.validateSetProgramDateTimePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"programDateTimePeriod",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetRedundantManifest(val *string) {
	if err := j.validateSetRedundantManifestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redundantManifest",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetSegmentLength(val *float64) {
	if err := j.validateSetSegmentLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentLength",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetSegmentsPerSubdirectory(val *float64) {
	if err := j.validateSetSegmentsPerSubdirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentsPerSubdirectory",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetStreamInfResolution(val *string) {
	if err := j.validateSetStreamInfResolutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamInfResolution",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetTimedMetadataId3Frame(val *string) {
	if err := j.validateSetTimedMetadataId3FrameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timedMetadataId3Frame",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetTimedMetadataId3Period(val *float64) {
	if err := j.validateSetTimedMetadataId3PeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timedMetadataId3Period",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetTimestampDeltaMilliseconds(val *float64) {
	if err := j.validateSetTimestampDeltaMillisecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampDeltaMilliseconds",
		val,
	)
}

func (j *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference)SetTsFileMode(val *string) {
	if err := j.validateSetTsFileModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tsFileMode",
		val,
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) PutCaptionLanguageMappings(value interface{}) {
	if err := m.validatePutCaptionLanguageMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putCaptionLanguageMappings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) PutDestination(value *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestination) {
	if err := m.validatePutDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDestination",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) PutHlsCdnSettings(value interface{}) {
	if err := m.validatePutHlsCdnSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putHlsCdnSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) PutKeyProviderSettings(value *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsKeyProviderSettings) {
	if err := m.validatePutKeyProviderSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putKeyProviderSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetAdMarkers() {
	_jsii_.InvokeVoid(
		m,
		"resetAdMarkers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetBaseUrlContent() {
	_jsii_.InvokeVoid(
		m,
		"resetBaseUrlContent",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetBaseUrlContent1() {
	_jsii_.InvokeVoid(
		m,
		"resetBaseUrlContent1",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetBaseUrlManifest() {
	_jsii_.InvokeVoid(
		m,
		"resetBaseUrlManifest",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetBaseUrlManifest1() {
	_jsii_.InvokeVoid(
		m,
		"resetBaseUrlManifest1",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetCaptionLanguageMappings() {
	_jsii_.InvokeVoid(
		m,
		"resetCaptionLanguageMappings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetCaptionLanguageSetting() {
	_jsii_.InvokeVoid(
		m,
		"resetCaptionLanguageSetting",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetClientCache() {
	_jsii_.InvokeVoid(
		m,
		"resetClientCache",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetCodecSpecification() {
	_jsii_.InvokeVoid(
		m,
		"resetCodecSpecification",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetConstantIv() {
	_jsii_.InvokeVoid(
		m,
		"resetConstantIv",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetDirectoryStructure() {
	_jsii_.InvokeVoid(
		m,
		"resetDirectoryStructure",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetDiscontinuityTags() {
	_jsii_.InvokeVoid(
		m,
		"resetDiscontinuityTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetEncryptionType() {
	_jsii_.InvokeVoid(
		m,
		"resetEncryptionType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetHlsCdnSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetHlsCdnSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetHlsId3SegmentTagging() {
	_jsii_.InvokeVoid(
		m,
		"resetHlsId3SegmentTagging",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetIframeOnlyPlaylists() {
	_jsii_.InvokeVoid(
		m,
		"resetIframeOnlyPlaylists",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetIncompleteSegmentBehavior() {
	_jsii_.InvokeVoid(
		m,
		"resetIncompleteSegmentBehavior",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetIndexNSegments() {
	_jsii_.InvokeVoid(
		m,
		"resetIndexNSegments",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetInputLossAction() {
	_jsii_.InvokeVoid(
		m,
		"resetInputLossAction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetIvInManifest() {
	_jsii_.InvokeVoid(
		m,
		"resetIvInManifest",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetIvSource() {
	_jsii_.InvokeVoid(
		m,
		"resetIvSource",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetKeepSegments() {
	_jsii_.InvokeVoid(
		m,
		"resetKeepSegments",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetKeyFormat() {
	_jsii_.InvokeVoid(
		m,
		"resetKeyFormat",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetKeyFormatVersions() {
	_jsii_.InvokeVoid(
		m,
		"resetKeyFormatVersions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetKeyProviderSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetKeyProviderSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetManifestCompression() {
	_jsii_.InvokeVoid(
		m,
		"resetManifestCompression",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetManifestDurationFormat() {
	_jsii_.InvokeVoid(
		m,
		"resetManifestDurationFormat",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetMinSegmentLength() {
	_jsii_.InvokeVoid(
		m,
		"resetMinSegmentLength",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		m,
		"resetMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetOutputSelection() {
	_jsii_.InvokeVoid(
		m,
		"resetOutputSelection",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetProgramDateTime() {
	_jsii_.InvokeVoid(
		m,
		"resetProgramDateTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetProgramDateTimeClock() {
	_jsii_.InvokeVoid(
		m,
		"resetProgramDateTimeClock",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetProgramDateTimePeriod() {
	_jsii_.InvokeVoid(
		m,
		"resetProgramDateTimePeriod",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetRedundantManifest() {
	_jsii_.InvokeVoid(
		m,
		"resetRedundantManifest",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetSegmentLength() {
	_jsii_.InvokeVoid(
		m,
		"resetSegmentLength",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetSegmentsPerSubdirectory() {
	_jsii_.InvokeVoid(
		m,
		"resetSegmentsPerSubdirectory",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetStreamInfResolution() {
	_jsii_.InvokeVoid(
		m,
		"resetStreamInfResolution",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetTimedMetadataId3Frame() {
	_jsii_.InvokeVoid(
		m,
		"resetTimedMetadataId3Frame",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetTimedMetadataId3Period() {
	_jsii_.InvokeVoid(
		m,
		"resetTimedMetadataId3Period",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetTimestampDeltaMilliseconds() {
	_jsii_.InvokeVoid(
		m,
		"resetTimestampDeltaMilliseconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ResetTsFileMode() {
	_jsii_.InvokeVoid(
		m,
		"resetTsFileMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

