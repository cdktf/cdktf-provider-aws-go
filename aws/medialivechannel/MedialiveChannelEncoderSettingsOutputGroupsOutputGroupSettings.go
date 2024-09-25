// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettings struct {
	// archive_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/medialive_channel#archive_group_settings MedialiveChannel#archive_group_settings}
	ArchiveGroupSettings interface{} `field:"optional" json:"archiveGroupSettings" yaml:"archiveGroupSettings"`
	// frame_capture_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/medialive_channel#frame_capture_group_settings MedialiveChannel#frame_capture_group_settings}
	FrameCaptureGroupSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsFrameCaptureGroupSettings `field:"optional" json:"frameCaptureGroupSettings" yaml:"frameCaptureGroupSettings"`
	// hls_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/medialive_channel#hls_group_settings MedialiveChannel#hls_group_settings}
	HlsGroupSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettings `field:"optional" json:"hlsGroupSettings" yaml:"hlsGroupSettings"`
	// media_package_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/medialive_channel#media_package_group_settings MedialiveChannel#media_package_group_settings}
	MediaPackageGroupSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsMediaPackageGroupSettings `field:"optional" json:"mediaPackageGroupSettings" yaml:"mediaPackageGroupSettings"`
	// ms_smooth_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/medialive_channel#ms_smooth_group_settings MedialiveChannel#ms_smooth_group_settings}
	MsSmoothGroupSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettings `field:"optional" json:"msSmoothGroupSettings" yaml:"msSmoothGroupSettings"`
	// multiplex_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/medialive_channel#multiplex_group_settings MedialiveChannel#multiplex_group_settings}
	MultiplexGroupSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsMultiplexGroupSettings `field:"optional" json:"multiplexGroupSettings" yaml:"multiplexGroupSettings"`
	// rtmp_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/medialive_channel#rtmp_group_settings MedialiveChannel#rtmp_group_settings}
	RtmpGroupSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsRtmpGroupSettings `field:"optional" json:"rtmpGroupSettings" yaml:"rtmpGroupSettings"`
	// udp_group_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/medialive_channel#udp_group_settings MedialiveChannel#udp_group_settings}
	UdpGroupSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsUdpGroupSettings `field:"optional" json:"udpGroupSettings" yaml:"udpGroupSettings"`
}

