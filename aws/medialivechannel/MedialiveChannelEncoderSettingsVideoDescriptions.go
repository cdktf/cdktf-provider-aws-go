// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsVideoDescriptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/medialive_channel#name MedialiveChannel#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// codec_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/medialive_channel#codec_settings MedialiveChannel#codec_settings}
	CodecSettings *MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettings `field:"optional" json:"codecSettings" yaml:"codecSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/medialive_channel#height MedialiveChannel#height}.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/medialive_channel#respond_to_afd MedialiveChannel#respond_to_afd}.
	RespondToAfd *string `field:"optional" json:"respondToAfd" yaml:"respondToAfd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/medialive_channel#scaling_behavior MedialiveChannel#scaling_behavior}.
	ScalingBehavior *string `field:"optional" json:"scalingBehavior" yaml:"scalingBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/medialive_channel#sharpness MedialiveChannel#sharpness}.
	Sharpness *float64 `field:"optional" json:"sharpness" yaml:"sharpness"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/medialive_channel#width MedialiveChannel#width}.
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

