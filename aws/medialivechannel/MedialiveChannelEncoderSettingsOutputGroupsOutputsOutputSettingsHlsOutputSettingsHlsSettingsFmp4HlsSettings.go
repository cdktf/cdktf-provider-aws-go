// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsHlsOutputSettingsHlsSettingsFmp4HlsSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/medialive_channel#audio_rendition_sets MedialiveChannel#audio_rendition_sets}.
	AudioRenditionSets *string `field:"optional" json:"audioRenditionSets" yaml:"audioRenditionSets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/medialive_channel#nielsen_id3_behavior MedialiveChannel#nielsen_id3_behavior}.
	NielsenId3Behavior *string `field:"optional" json:"nielsenId3Behavior" yaml:"nielsenId3Behavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/medialive_channel#timed_metadata_behavior MedialiveChannel#timed_metadata_behavior}.
	TimedMetadataBehavior *string `field:"optional" json:"timedMetadataBehavior" yaml:"timedMetadataBehavior"`
}

