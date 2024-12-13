// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsAudioDescriptionsAudioWatermarkSettingsNielsenWatermarksSettings struct {
	// nielsen_cbet_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/medialive_channel#nielsen_cbet_settings MedialiveChannel#nielsen_cbet_settings}
	NielsenCbetSettings *MedialiveChannelEncoderSettingsAudioDescriptionsAudioWatermarkSettingsNielsenWatermarksSettingsNielsenCbetSettings `field:"optional" json:"nielsenCbetSettings" yaml:"nielsenCbetSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/medialive_channel#nielsen_distribution_type MedialiveChannel#nielsen_distribution_type}.
	NielsenDistributionType *string `field:"optional" json:"nielsenDistributionType" yaml:"nielsenDistributionType"`
	// nielsen_naes_ii_nw_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/medialive_channel#nielsen_naes_ii_nw_settings MedialiveChannel#nielsen_naes_ii_nw_settings}
	NielsenNaesIiNwSettings interface{} `field:"optional" json:"nielsenNaesIiNwSettings" yaml:"nielsenNaesIiNwSettings"`
}

