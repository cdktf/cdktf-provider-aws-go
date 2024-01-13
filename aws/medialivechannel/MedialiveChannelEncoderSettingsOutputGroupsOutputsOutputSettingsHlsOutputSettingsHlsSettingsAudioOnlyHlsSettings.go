// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsHlsOutputSettingsHlsSettingsAudioOnlyHlsSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/medialive_channel#audio_group_id MedialiveChannel#audio_group_id}.
	AudioGroupId *string `field:"optional" json:"audioGroupId" yaml:"audioGroupId"`
	// audio_only_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/medialive_channel#audio_only_image MedialiveChannel#audio_only_image}
	AudioOnlyImage *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsHlsOutputSettingsHlsSettingsAudioOnlyHlsSettingsAudioOnlyImage `field:"optional" json:"audioOnlyImage" yaml:"audioOnlyImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/medialive_channel#audio_track_type MedialiveChannel#audio_track_type}.
	AudioTrackType *string `field:"optional" json:"audioTrackType" yaml:"audioTrackType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/medialive_channel#segment_type MedialiveChannel#segment_type}.
	SegmentType *string `field:"optional" json:"segmentType" yaml:"segmentType"`
}

