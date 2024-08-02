// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettings struct {
	// aac_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.61.0/docs/resources/medialive_channel#aac_settings MedialiveChannel#aac_settings}
	AacSettings *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettings `field:"optional" json:"aacSettings" yaml:"aacSettings"`
	// ac3_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.61.0/docs/resources/medialive_channel#ac3_settings MedialiveChannel#ac3_settings}
	Ac3Settings *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAc3Settings `field:"optional" json:"ac3Settings" yaml:"ac3Settings"`
	// eac3_atmos_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.61.0/docs/resources/medialive_channel#eac3_atmos_settings MedialiveChannel#eac3_atmos_settings}
	Eac3AtmosSettings *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettings `field:"optional" json:"eac3AtmosSettings" yaml:"eac3AtmosSettings"`
	// eac3_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.61.0/docs/resources/medialive_channel#eac3_settings MedialiveChannel#eac3_settings}
	Eac3Settings *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3Settings `field:"optional" json:"eac3Settings" yaml:"eac3Settings"`
	// mp2_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.61.0/docs/resources/medialive_channel#mp2_settings MedialiveChannel#mp2_settings}
	Mp2Settings *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsMp2Settings `field:"optional" json:"mp2Settings" yaml:"mp2Settings"`
	// pass_through_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.61.0/docs/resources/medialive_channel#pass_through_settings MedialiveChannel#pass_through_settings}
	PassThroughSettings *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsPassThroughSettings `field:"optional" json:"passThroughSettings" yaml:"passThroughSettings"`
	// wav_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.61.0/docs/resources/medialive_channel#wav_settings MedialiveChannel#wav_settings}
	WavSettings *MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsWavSettings `field:"optional" json:"wavSettings" yaml:"wavSettings"`
}

