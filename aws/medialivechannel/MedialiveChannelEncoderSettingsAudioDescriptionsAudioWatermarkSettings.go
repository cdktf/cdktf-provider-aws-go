// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsAudioDescriptionsAudioWatermarkSettings struct {
	// nielsen_watermarks_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/medialive_channel#nielsen_watermarks_settings MedialiveChannel#nielsen_watermarks_settings}
	NielsenWatermarksSettings *MedialiveChannelEncoderSettingsAudioDescriptionsAudioWatermarkSettingsNielsenWatermarksSettings `field:"optional" json:"nielsenWatermarksSettings" yaml:"nielsenWatermarksSettings"`
}

