// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettings struct {
	// color_space_passthrough_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/medialive_channel#color_space_passthrough_settings MedialiveChannel#color_space_passthrough_settings}
	ColorSpacePassthroughSettings *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsColorSpacePassthroughSettings `field:"optional" json:"colorSpacePassthroughSettings" yaml:"colorSpacePassthroughSettings"`
	// dolby_vision81_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/medialive_channel#dolby_vision81_settings MedialiveChannel#dolby_vision81_settings}
	DolbyVision81Settings *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsDolbyVision81Settings `field:"optional" json:"dolbyVision81Settings" yaml:"dolbyVision81Settings"`
	// hdr10_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/medialive_channel#hdr10_settings MedialiveChannel#hdr10_settings}
	Hdr10Settings *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsHdr10Settings `field:"optional" json:"hdr10Settings" yaml:"hdr10Settings"`
	// rec601_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/medialive_channel#rec601_settings MedialiveChannel#rec601_settings}
	Rec601Settings *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsRec601Settings `field:"optional" json:"rec601Settings" yaml:"rec601Settings"`
	// rec709_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/medialive_channel#rec709_settings MedialiveChannel#rec709_settings}
	Rec709Settings *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsColorSpaceSettingsRec709Settings `field:"optional" json:"rec709Settings" yaml:"rec709Settings"`
}

