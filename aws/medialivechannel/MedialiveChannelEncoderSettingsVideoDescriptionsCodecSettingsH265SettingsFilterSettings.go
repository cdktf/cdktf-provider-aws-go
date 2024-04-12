// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettings struct {
	// temporal_filter_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#temporal_filter_settings MedialiveChannel#temporal_filter_settings}
	TemporalFilterSettings *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettingsTemporalFilterSettings `field:"optional" json:"temporalFilterSettings" yaml:"temporalFilterSettings"`
}

