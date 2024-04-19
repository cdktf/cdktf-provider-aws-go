// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettings struct {
	// output_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/medialive_channel#output_groups MedialiveChannel#output_groups}
	OutputGroups interface{} `field:"required" json:"outputGroups" yaml:"outputGroups"`
	// timecode_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/medialive_channel#timecode_config MedialiveChannel#timecode_config}
	TimecodeConfig *MedialiveChannelEncoderSettingsTimecodeConfig `field:"required" json:"timecodeConfig" yaml:"timecodeConfig"`
	// audio_descriptions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/medialive_channel#audio_descriptions MedialiveChannel#audio_descriptions}
	AudioDescriptions interface{} `field:"optional" json:"audioDescriptions" yaml:"audioDescriptions"`
	// avail_blanking block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/medialive_channel#avail_blanking MedialiveChannel#avail_blanking}
	AvailBlanking *MedialiveChannelEncoderSettingsAvailBlanking `field:"optional" json:"availBlanking" yaml:"availBlanking"`
	// caption_descriptions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/medialive_channel#caption_descriptions MedialiveChannel#caption_descriptions}
	CaptionDescriptions interface{} `field:"optional" json:"captionDescriptions" yaml:"captionDescriptions"`
	// global_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/medialive_channel#global_configuration MedialiveChannel#global_configuration}
	GlobalConfiguration *MedialiveChannelEncoderSettingsGlobalConfiguration `field:"optional" json:"globalConfiguration" yaml:"globalConfiguration"`
	// motion_graphics_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/medialive_channel#motion_graphics_configuration MedialiveChannel#motion_graphics_configuration}
	MotionGraphicsConfiguration *MedialiveChannelEncoderSettingsMotionGraphicsConfiguration `field:"optional" json:"motionGraphicsConfiguration" yaml:"motionGraphicsConfiguration"`
	// nielsen_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/medialive_channel#nielsen_configuration MedialiveChannel#nielsen_configuration}
	NielsenConfiguration *MedialiveChannelEncoderSettingsNielsenConfiguration `field:"optional" json:"nielsenConfiguration" yaml:"nielsenConfiguration"`
	// video_descriptions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/medialive_channel#video_descriptions MedialiveChannel#video_descriptions}
	VideoDescriptions interface{} `field:"optional" json:"videoDescriptions" yaml:"videoDescriptions"`
}

