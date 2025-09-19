// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettings struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#destination MedialiveChannel#destination}
	Destination *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestination `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#ad_markers MedialiveChannel#ad_markers}.
	AdMarkers *[]*string `field:"optional" json:"adMarkers" yaml:"adMarkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#base_url_content MedialiveChannel#base_url_content}.
	BaseUrlContent *string `field:"optional" json:"baseUrlContent" yaml:"baseUrlContent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#base_url_content1 MedialiveChannel#base_url_content1}.
	BaseUrlContent1 *string `field:"optional" json:"baseUrlContent1" yaml:"baseUrlContent1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#base_url_manifest MedialiveChannel#base_url_manifest}.
	BaseUrlManifest *string `field:"optional" json:"baseUrlManifest" yaml:"baseUrlManifest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#base_url_manifest1 MedialiveChannel#base_url_manifest1}.
	BaseUrlManifest1 *string `field:"optional" json:"baseUrlManifest1" yaml:"baseUrlManifest1"`
	// caption_language_mappings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#caption_language_mappings MedialiveChannel#caption_language_mappings}
	CaptionLanguageMappings interface{} `field:"optional" json:"captionLanguageMappings" yaml:"captionLanguageMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#caption_language_setting MedialiveChannel#caption_language_setting}.
	CaptionLanguageSetting *string `field:"optional" json:"captionLanguageSetting" yaml:"captionLanguageSetting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#client_cache MedialiveChannel#client_cache}.
	ClientCache *string `field:"optional" json:"clientCache" yaml:"clientCache"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#codec_specification MedialiveChannel#codec_specification}.
	CodecSpecification *string `field:"optional" json:"codecSpecification" yaml:"codecSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#constant_iv MedialiveChannel#constant_iv}.
	ConstantIv *string `field:"optional" json:"constantIv" yaml:"constantIv"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#directory_structure MedialiveChannel#directory_structure}.
	DirectoryStructure *string `field:"optional" json:"directoryStructure" yaml:"directoryStructure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#discontinuity_tags MedialiveChannel#discontinuity_tags}.
	DiscontinuityTags *string `field:"optional" json:"discontinuityTags" yaml:"discontinuityTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#encryption_type MedialiveChannel#encryption_type}.
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// hls_cdn_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#hls_cdn_settings MedialiveChannel#hls_cdn_settings}
	HlsCdnSettings interface{} `field:"optional" json:"hlsCdnSettings" yaml:"hlsCdnSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#hls_id3_segment_tagging MedialiveChannel#hls_id3_segment_tagging}.
	HlsId3SegmentTagging *string `field:"optional" json:"hlsId3SegmentTagging" yaml:"hlsId3SegmentTagging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#iframe_only_playlists MedialiveChannel#iframe_only_playlists}.
	IframeOnlyPlaylists *string `field:"optional" json:"iframeOnlyPlaylists" yaml:"iframeOnlyPlaylists"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#incomplete_segment_behavior MedialiveChannel#incomplete_segment_behavior}.
	IncompleteSegmentBehavior *string `field:"optional" json:"incompleteSegmentBehavior" yaml:"incompleteSegmentBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#index_n_segments MedialiveChannel#index_n_segments}.
	IndexNSegments *float64 `field:"optional" json:"indexNSegments" yaml:"indexNSegments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#input_loss_action MedialiveChannel#input_loss_action}.
	InputLossAction *string `field:"optional" json:"inputLossAction" yaml:"inputLossAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#iv_in_manifest MedialiveChannel#iv_in_manifest}.
	IvInManifest *string `field:"optional" json:"ivInManifest" yaml:"ivInManifest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#iv_source MedialiveChannel#iv_source}.
	IvSource *string `field:"optional" json:"ivSource" yaml:"ivSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#keep_segments MedialiveChannel#keep_segments}.
	KeepSegments *float64 `field:"optional" json:"keepSegments" yaml:"keepSegments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#key_format MedialiveChannel#key_format}.
	KeyFormat *string `field:"optional" json:"keyFormat" yaml:"keyFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#key_format_versions MedialiveChannel#key_format_versions}.
	KeyFormatVersions *string `field:"optional" json:"keyFormatVersions" yaml:"keyFormatVersions"`
	// key_provider_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#key_provider_settings MedialiveChannel#key_provider_settings}
	KeyProviderSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsKeyProviderSettings `field:"optional" json:"keyProviderSettings" yaml:"keyProviderSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#manifest_compression MedialiveChannel#manifest_compression}.
	ManifestCompression *string `field:"optional" json:"manifestCompression" yaml:"manifestCompression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#manifest_duration_format MedialiveChannel#manifest_duration_format}.
	ManifestDurationFormat *string `field:"optional" json:"manifestDurationFormat" yaml:"manifestDurationFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#min_segment_length MedialiveChannel#min_segment_length}.
	MinSegmentLength *float64 `field:"optional" json:"minSegmentLength" yaml:"minSegmentLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#mode MedialiveChannel#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#output_selection MedialiveChannel#output_selection}.
	OutputSelection *string `field:"optional" json:"outputSelection" yaml:"outputSelection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#program_date_time MedialiveChannel#program_date_time}.
	ProgramDateTime *string `field:"optional" json:"programDateTime" yaml:"programDateTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#program_date_time_clock MedialiveChannel#program_date_time_clock}.
	ProgramDateTimeClock *string `field:"optional" json:"programDateTimeClock" yaml:"programDateTimeClock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#program_date_time_period MedialiveChannel#program_date_time_period}.
	ProgramDateTimePeriod *float64 `field:"optional" json:"programDateTimePeriod" yaml:"programDateTimePeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#redundant_manifest MedialiveChannel#redundant_manifest}.
	RedundantManifest *string `field:"optional" json:"redundantManifest" yaml:"redundantManifest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#segment_length MedialiveChannel#segment_length}.
	SegmentLength *float64 `field:"optional" json:"segmentLength" yaml:"segmentLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#segments_per_subdirectory MedialiveChannel#segments_per_subdirectory}.
	SegmentsPerSubdirectory *float64 `field:"optional" json:"segmentsPerSubdirectory" yaml:"segmentsPerSubdirectory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#stream_inf_resolution MedialiveChannel#stream_inf_resolution}.
	StreamInfResolution *string `field:"optional" json:"streamInfResolution" yaml:"streamInfResolution"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#timed_metadata_id3_frame MedialiveChannel#timed_metadata_id3_frame}.
	TimedMetadataId3Frame *string `field:"optional" json:"timedMetadataId3Frame" yaml:"timedMetadataId3Frame"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#timed_metadata_id3_period MedialiveChannel#timed_metadata_id3_period}.
	TimedMetadataId3Period *float64 `field:"optional" json:"timedMetadataId3Period" yaml:"timedMetadataId3Period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#timestamp_delta_milliseconds MedialiveChannel#timestamp_delta_milliseconds}.
	TimestampDeltaMilliseconds *float64 `field:"optional" json:"timestampDeltaMilliseconds" yaml:"timestampDeltaMilliseconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/medialive_channel#ts_file_mode MedialiveChannel#ts_file_mode}.
	TsFileMode *string `field:"optional" json:"tsFileMode" yaml:"tsFileMode"`
}

