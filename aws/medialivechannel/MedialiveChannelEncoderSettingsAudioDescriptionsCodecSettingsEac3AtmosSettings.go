// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsEac3AtmosSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/medialive_channel#bitrate MedialiveChannel#bitrate}.
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/medialive_channel#coding_mode MedialiveChannel#coding_mode}.
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/medialive_channel#dialnorm MedialiveChannel#dialnorm}.
	Dialnorm *float64 `field:"optional" json:"dialnorm" yaml:"dialnorm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/medialive_channel#drc_line MedialiveChannel#drc_line}.
	DrcLine *string `field:"optional" json:"drcLine" yaml:"drcLine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/medialive_channel#drc_rf MedialiveChannel#drc_rf}.
	DrcRf *string `field:"optional" json:"drcRf" yaml:"drcRf"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/medialive_channel#height_trim MedialiveChannel#height_trim}.
	HeightTrim *float64 `field:"optional" json:"heightTrim" yaml:"heightTrim"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/medialive_channel#surround_trim MedialiveChannel#surround_trim}.
	SurroundTrim *float64 `field:"optional" json:"surroundTrim" yaml:"surroundTrim"`
}

