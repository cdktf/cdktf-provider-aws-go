// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265Settings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#bitrate MedialiveChannel#bitrate}.
	Bitrate *float64 `field:"required" json:"bitrate" yaml:"bitrate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#framerate_denominator MedialiveChannel#framerate_denominator}.
	FramerateDenominator *float64 `field:"required" json:"framerateDenominator" yaml:"framerateDenominator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#framerate_numerator MedialiveChannel#framerate_numerator}.
	FramerateNumerator *float64 `field:"required" json:"framerateNumerator" yaml:"framerateNumerator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#adaptive_quantization MedialiveChannel#adaptive_quantization}.
	AdaptiveQuantization *string `field:"optional" json:"adaptiveQuantization" yaml:"adaptiveQuantization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#afd_signaling MedialiveChannel#afd_signaling}.
	AfdSignaling *string `field:"optional" json:"afdSignaling" yaml:"afdSignaling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#alternative_transfer_function MedialiveChannel#alternative_transfer_function}.
	AlternativeTransferFunction *string `field:"optional" json:"alternativeTransferFunction" yaml:"alternativeTransferFunction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#buf_size MedialiveChannel#buf_size}.
	BufSize *float64 `field:"optional" json:"bufSize" yaml:"bufSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#color_metadata MedialiveChannel#color_metadata}.
	ColorMetadata *string `field:"optional" json:"colorMetadata" yaml:"colorMetadata"`
	// color_space_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#color_space_settings MedialiveChannel#color_space_settings}
	ColorSpaceSettings *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettings `field:"optional" json:"colorSpaceSettings" yaml:"colorSpaceSettings"`
	// filter_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#filter_settings MedialiveChannel#filter_settings}
	FilterSettings *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettings `field:"optional" json:"filterSettings" yaml:"filterSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#fixed_afd MedialiveChannel#fixed_afd}.
	FixedAfd *string `field:"optional" json:"fixedAfd" yaml:"fixedAfd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#flicker_aq MedialiveChannel#flicker_aq}.
	FlickerAq *string `field:"optional" json:"flickerAq" yaml:"flickerAq"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#gop_closed_cadence MedialiveChannel#gop_closed_cadence}.
	GopClosedCadence *float64 `field:"optional" json:"gopClosedCadence" yaml:"gopClosedCadence"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#gop_size MedialiveChannel#gop_size}.
	GopSize *float64 `field:"optional" json:"gopSize" yaml:"gopSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#gop_size_units MedialiveChannel#gop_size_units}.
	GopSizeUnits *string `field:"optional" json:"gopSizeUnits" yaml:"gopSizeUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#level MedialiveChannel#level}.
	Level *string `field:"optional" json:"level" yaml:"level"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#look_ahead_rate_control MedialiveChannel#look_ahead_rate_control}.
	LookAheadRateControl *string `field:"optional" json:"lookAheadRateControl" yaml:"lookAheadRateControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#max_bitrate MedialiveChannel#max_bitrate}.
	MaxBitrate *float64 `field:"optional" json:"maxBitrate" yaml:"maxBitrate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#min_i_interval MedialiveChannel#min_i_interval}.
	MinIInterval *float64 `field:"optional" json:"minIInterval" yaml:"minIInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#par_denominator MedialiveChannel#par_denominator}.
	ParDenominator *float64 `field:"optional" json:"parDenominator" yaml:"parDenominator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#par_numerator MedialiveChannel#par_numerator}.
	ParNumerator *float64 `field:"optional" json:"parNumerator" yaml:"parNumerator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#profile MedialiveChannel#profile}.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#qvbr_quality_level MedialiveChannel#qvbr_quality_level}.
	QvbrQualityLevel *float64 `field:"optional" json:"qvbrQualityLevel" yaml:"qvbrQualityLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#rate_control_mode MedialiveChannel#rate_control_mode}.
	RateControlMode *string `field:"optional" json:"rateControlMode" yaml:"rateControlMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#scan_type MedialiveChannel#scan_type}.
	ScanType *string `field:"optional" json:"scanType" yaml:"scanType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#scene_change_detect MedialiveChannel#scene_change_detect}.
	SceneChangeDetect *string `field:"optional" json:"sceneChangeDetect" yaml:"sceneChangeDetect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#slices MedialiveChannel#slices}.
	Slices *float64 `field:"optional" json:"slices" yaml:"slices"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#tier MedialiveChannel#tier}.
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
	// timecode_burnin_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#timecode_burnin_settings MedialiveChannel#timecode_burnin_settings}
	TimecodeBurninSettings *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsTimecodeBurninSettings `field:"optional" json:"timecodeBurninSettings" yaml:"timecodeBurninSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/medialive_channel#timecode_insertion MedialiveChannel#timecode_insertion}.
	TimecodeInsertion *string `field:"optional" json:"timecodeInsertion" yaml:"timecodeInsertion"`
}

