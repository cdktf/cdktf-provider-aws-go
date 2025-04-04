// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAc3Settings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/medialive_channel#bitrate MedialiveChannel#bitrate}.
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/medialive_channel#bitstream_mode MedialiveChannel#bitstream_mode}.
	BitstreamMode *string `field:"optional" json:"bitstreamMode" yaml:"bitstreamMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/medialive_channel#coding_mode MedialiveChannel#coding_mode}.
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/medialive_channel#dialnorm MedialiveChannel#dialnorm}.
	Dialnorm *float64 `field:"optional" json:"dialnorm" yaml:"dialnorm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/medialive_channel#drc_profile MedialiveChannel#drc_profile}.
	DrcProfile *string `field:"optional" json:"drcProfile" yaml:"drcProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/medialive_channel#lfe_filter MedialiveChannel#lfe_filter}.
	LfeFilter *string `field:"optional" json:"lfeFilter" yaml:"lfeFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/medialive_channel#metadata_control MedialiveChannel#metadata_control}.
	MetadataControl *string `field:"optional" json:"metadataControl" yaml:"metadataControl"`
}

