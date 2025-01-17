// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsFrameCaptureGroupSettings struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/medialive_channel#destination MedialiveChannel#destination}
	Destination *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsFrameCaptureGroupSettingsDestination `field:"required" json:"destination" yaml:"destination"`
	// frame_capture_cdn_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/resources/medialive_channel#frame_capture_cdn_settings MedialiveChannel#frame_capture_cdn_settings}
	FrameCaptureCdnSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsFrameCaptureGroupSettingsFrameCaptureCdnSettings `field:"optional" json:"frameCaptureCdnSettings" yaml:"frameCaptureCdnSettings"`
}

