// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsTimecodeBurninSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/medialive_channel#prefix MedialiveChannel#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/medialive_channel#timecode_burnin_font_size MedialiveChannel#timecode_burnin_font_size}.
	TimecodeBurninFontSize *string `field:"optional" json:"timecodeBurninFontSize" yaml:"timecodeBurninFontSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/medialive_channel#timecode_burnin_position MedialiveChannel#timecode_burnin_position}.
	TimecodeBurninPosition *string `field:"optional" json:"timecodeBurninPosition" yaml:"timecodeBurninPosition"`
}

