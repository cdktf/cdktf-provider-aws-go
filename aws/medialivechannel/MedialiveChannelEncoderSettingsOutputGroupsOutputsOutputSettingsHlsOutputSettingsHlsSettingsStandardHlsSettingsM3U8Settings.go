// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsHlsOutputSettingsHlsSettingsStandardHlsSettingsM3U8Settings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/medialive_channel#audio_frames_per_pes MedialiveChannel#audio_frames_per_pes}.
	AudioFramesPerPes *float64 `field:"optional" json:"audioFramesPerPes" yaml:"audioFramesPerPes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/medialive_channel#audio_pids MedialiveChannel#audio_pids}.
	AudioPids *string `field:"optional" json:"audioPids" yaml:"audioPids"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/medialive_channel#ecm_pid MedialiveChannel#ecm_pid}.
	EcmPid *string `field:"optional" json:"ecmPid" yaml:"ecmPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/medialive_channel#nielsen_id3_behavior MedialiveChannel#nielsen_id3_behavior}.
	NielsenId3Behavior *string `field:"optional" json:"nielsenId3Behavior" yaml:"nielsenId3Behavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/medialive_channel#pat_interval MedialiveChannel#pat_interval}.
	PatInterval *float64 `field:"optional" json:"patInterval" yaml:"patInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/medialive_channel#pcr_control MedialiveChannel#pcr_control}.
	PcrControl *string `field:"optional" json:"pcrControl" yaml:"pcrControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/medialive_channel#pcr_period MedialiveChannel#pcr_period}.
	PcrPeriod *float64 `field:"optional" json:"pcrPeriod" yaml:"pcrPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/medialive_channel#pcr_pid MedialiveChannel#pcr_pid}.
	PcrPid *string `field:"optional" json:"pcrPid" yaml:"pcrPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/medialive_channel#pmt_interval MedialiveChannel#pmt_interval}.
	PmtInterval *float64 `field:"optional" json:"pmtInterval" yaml:"pmtInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/medialive_channel#pmt_pid MedialiveChannel#pmt_pid}.
	PmtPid *string `field:"optional" json:"pmtPid" yaml:"pmtPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/medialive_channel#program_num MedialiveChannel#program_num}.
	ProgramNum *float64 `field:"optional" json:"programNum" yaml:"programNum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/medialive_channel#scte35_behavior MedialiveChannel#scte35_behavior}.
	Scte35Behavior *string `field:"optional" json:"scte35Behavior" yaml:"scte35Behavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/medialive_channel#scte35_pid MedialiveChannel#scte35_pid}.
	Scte35Pid *string `field:"optional" json:"scte35Pid" yaml:"scte35Pid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/medialive_channel#timed_metadata_behavior MedialiveChannel#timed_metadata_behavior}.
	TimedMetadataBehavior *string `field:"optional" json:"timedMetadataBehavior" yaml:"timedMetadataBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/medialive_channel#timed_metadata_pid MedialiveChannel#timed_metadata_pid}.
	TimedMetadataPid *string `field:"optional" json:"timedMetadataPid" yaml:"timedMetadataPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/medialive_channel#transport_stream_id MedialiveChannel#transport_stream_id}.
	TransportStreamId *float64 `field:"optional" json:"transportStreamId" yaml:"transportStreamId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/medialive_channel#video_pid MedialiveChannel#video_pid}.
	VideoPid *string `field:"optional" json:"videoPid" yaml:"videoPid"`
}

