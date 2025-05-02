// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettings struct {
	// archive_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/medialive_channel#archive_output_settings MedialiveChannel#archive_output_settings}
	ArchiveOutputSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettings `field:"optional" json:"archiveOutputSettings" yaml:"archiveOutputSettings"`
	// frame_capture_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/medialive_channel#frame_capture_output_settings MedialiveChannel#frame_capture_output_settings}
	FrameCaptureOutputSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsFrameCaptureOutputSettings `field:"optional" json:"frameCaptureOutputSettings" yaml:"frameCaptureOutputSettings"`
	// hls_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/medialive_channel#hls_output_settings MedialiveChannel#hls_output_settings}
	HlsOutputSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsHlsOutputSettings `field:"optional" json:"hlsOutputSettings" yaml:"hlsOutputSettings"`
	// media_package_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/medialive_channel#media_package_output_settings MedialiveChannel#media_package_output_settings}
	MediaPackageOutputSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsMediaPackageOutputSettings `field:"optional" json:"mediaPackageOutputSettings" yaml:"mediaPackageOutputSettings"`
	// ms_smooth_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/medialive_channel#ms_smooth_output_settings MedialiveChannel#ms_smooth_output_settings}
	MsSmoothOutputSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsMsSmoothOutputSettings `field:"optional" json:"msSmoothOutputSettings" yaml:"msSmoothOutputSettings"`
	// multiplex_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/medialive_channel#multiplex_output_settings MedialiveChannel#multiplex_output_settings}
	MultiplexOutputSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsMultiplexOutputSettings `field:"optional" json:"multiplexOutputSettings" yaml:"multiplexOutputSettings"`
	// rtmp_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/medialive_channel#rtmp_output_settings MedialiveChannel#rtmp_output_settings}
	RtmpOutputSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsRtmpOutputSettings `field:"optional" json:"rtmpOutputSettings" yaml:"rtmpOutputSettings"`
	// udp_output_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/medialive_channel#udp_output_settings MedialiveChannel#udp_output_settings}
	UdpOutputSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettings `field:"optional" json:"udpOutputSettings" yaml:"udpOutputSettings"`
}

