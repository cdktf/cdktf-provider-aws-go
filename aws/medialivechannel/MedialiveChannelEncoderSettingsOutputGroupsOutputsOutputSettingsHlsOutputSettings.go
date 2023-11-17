// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsHlsOutputSettings struct {
	// hls_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/medialive_channel#hls_settings MedialiveChannel#hls_settings}
	HlsSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsHlsOutputSettingsHlsSettings `field:"required" json:"hlsSettings" yaml:"hlsSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/medialive_channel#h265_packaging_type MedialiveChannel#h265_packaging_type}.
	H265PackagingType *string `field:"optional" json:"h265PackagingType" yaml:"h265PackagingType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/medialive_channel#name_modifier MedialiveChannel#name_modifier}.
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/medialive_channel#segment_modifier MedialiveChannel#segment_modifier}.
	SegmentModifier *string `field:"optional" json:"segmentModifier" yaml:"segmentModifier"`
}

