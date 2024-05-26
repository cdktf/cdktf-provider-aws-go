// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputs struct {
	// output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/medialive_channel#output_settings MedialiveChannel#output_settings}
	OutputSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettings `field:"required" json:"outputSettings" yaml:"outputSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/medialive_channel#audio_description_names MedialiveChannel#audio_description_names}.
	AudioDescriptionNames *[]*string `field:"optional" json:"audioDescriptionNames" yaml:"audioDescriptionNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/medialive_channel#caption_description_names MedialiveChannel#caption_description_names}.
	CaptionDescriptionNames *[]*string `field:"optional" json:"captionDescriptionNames" yaml:"captionDescriptionNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/medialive_channel#output_name MedialiveChannel#output_name}.
	OutputName *string `field:"optional" json:"outputName" yaml:"outputName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/medialive_channel#video_description_name MedialiveChannel#video_description_name}.
	VideoDescriptionName *string `field:"optional" json:"videoDescriptionName" yaml:"videoDescriptionName"`
}

