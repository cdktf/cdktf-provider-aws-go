// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehavior struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.19.0/docs/resources/medialive_channel#black_frame_msec MedialiveChannel#black_frame_msec}.
	BlackFrameMsec *float64 `field:"optional" json:"blackFrameMsec" yaml:"blackFrameMsec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.19.0/docs/resources/medialive_channel#input_loss_image_color MedialiveChannel#input_loss_image_color}.
	InputLossImageColor *string `field:"optional" json:"inputLossImageColor" yaml:"inputLossImageColor"`
	// input_loss_image_slate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.19.0/docs/resources/medialive_channel#input_loss_image_slate MedialiveChannel#input_loss_image_slate}
	InputLossImageSlate *MedialiveChannelEncoderSettingsGlobalConfigurationInputLossBehaviorInputLossImageSlate `field:"optional" json:"inputLossImageSlate" yaml:"inputLossImageSlate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.19.0/docs/resources/medialive_channel#input_loss_image_type MedialiveChannel#input_loss_image_type}.
	InputLossImageType *string `field:"optional" json:"inputLossImageType" yaml:"inputLossImageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.19.0/docs/resources/medialive_channel#repeat_frame_msec MedialiveChannel#repeat_frame_msec}.
	RepeatFrameMsec *float64 `field:"optional" json:"repeatFrameMsec" yaml:"repeatFrameMsec"`
}

