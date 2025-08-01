// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsFrameCaptureGroupSettingsFrameCaptureCdnSettings struct {
	// frame_capture_s3_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/medialive_channel#frame_capture_s3_settings MedialiveChannel#frame_capture_s3_settings}
	FrameCaptureS3Settings *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsFrameCaptureGroupSettingsFrameCaptureCdnSettingsFrameCaptureS3Settings `field:"optional" json:"frameCaptureS3Settings" yaml:"frameCaptureS3Settings"`
}

