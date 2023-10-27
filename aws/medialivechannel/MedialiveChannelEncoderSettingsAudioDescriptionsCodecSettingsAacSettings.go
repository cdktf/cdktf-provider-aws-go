// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsAudioDescriptionsCodecSettingsAacSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/medialive_channel#bitrate MedialiveChannel#bitrate}.
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/medialive_channel#coding_mode MedialiveChannel#coding_mode}.
	CodingMode *string `field:"optional" json:"codingMode" yaml:"codingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/medialive_channel#input_type MedialiveChannel#input_type}.
	InputType *string `field:"optional" json:"inputType" yaml:"inputType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/medialive_channel#profile MedialiveChannel#profile}.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/medialive_channel#rate_control_mode MedialiveChannel#rate_control_mode}.
	RateControlMode *string `field:"optional" json:"rateControlMode" yaml:"rateControlMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/medialive_channel#raw_format MedialiveChannel#raw_format}.
	RawFormat *string `field:"optional" json:"rawFormat" yaml:"rawFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/medialive_channel#sample_rate MedialiveChannel#sample_rate}.
	SampleRate *float64 `field:"optional" json:"sampleRate" yaml:"sampleRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/medialive_channel#spec MedialiveChannel#spec}.
	Spec *string `field:"optional" json:"spec" yaml:"spec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/medialive_channel#vbr_quality MedialiveChannel#vbr_quality}.
	VbrQuality *string `field:"optional" json:"vbrQuality" yaml:"vbrQuality"`
}

