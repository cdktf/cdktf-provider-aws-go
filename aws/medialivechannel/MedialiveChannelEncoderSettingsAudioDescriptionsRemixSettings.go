// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsAudioDescriptionsRemixSettings struct {
	// channel_mappings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/medialive_channel#channel_mappings MedialiveChannel#channel_mappings}
	ChannelMappings interface{} `field:"required" json:"channelMappings" yaml:"channelMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/medialive_channel#channels_in MedialiveChannel#channels_in}.
	ChannelsIn *float64 `field:"optional" json:"channelsIn" yaml:"channelsIn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/medialive_channel#channels_out MedialiveChannel#channels_out}.
	ChannelsOut *float64 `field:"optional" json:"channelsOut" yaml:"channelsOut"`
}

