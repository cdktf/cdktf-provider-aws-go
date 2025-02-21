// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsWavSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/medialive_channel#bit_depth MedialiveChannel#bit_depth}.
	BitDepth *float64 `field:"optional" json:"bitDepth" yaml:"bitDepth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/medialive_channel#coding_mode MedialiveChannel#coding_mode}.
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/medialive_channel#sample_rate MedialiveChannel#sample_rate}.
	SampleRate *float64 `field:"optional" json:"sampleRate" yaml:"sampleRate"`
}

