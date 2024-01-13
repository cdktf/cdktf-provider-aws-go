// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsHlsOutputSettingsHlsSettingsStandardHlsSettings struct {
	// m3u8_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/medialive_channel#m3u8_settings MedialiveChannel#m3u8_settings}
	M3U8Settings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsHlsOutputSettingsHlsSettingsStandardHlsSettingsM3U8Settings `field:"required" json:"m3U8Settings" yaml:"m3U8Settings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/medialive_channel#audio_rendition_sets MedialiveChannel#audio_rendition_sets}.
	AudioRenditionSets *string `field:"optional" json:"audioRenditionSets" yaml:"audioRenditionSets"`
}

