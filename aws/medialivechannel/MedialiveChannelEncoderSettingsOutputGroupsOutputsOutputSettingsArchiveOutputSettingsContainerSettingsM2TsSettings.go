// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2TsSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#absent_input_audio_behavior MedialiveChannel#absent_input_audio_behavior}.
	AbsentInputAudioBehavior *string `field:"optional" json:"absentInputAudioBehavior" yaml:"absentInputAudioBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#arib MedialiveChannel#arib}.
	Arib *string `field:"optional" json:"arib" yaml:"arib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#arib_captions_pid MedialiveChannel#arib_captions_pid}.
	AribCaptionsPid *string `field:"optional" json:"aribCaptionsPid" yaml:"aribCaptionsPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#arib_captions_pid_control MedialiveChannel#arib_captions_pid_control}.
	AribCaptionsPidControl *string `field:"optional" json:"aribCaptionsPidControl" yaml:"aribCaptionsPidControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#audio_buffer_model MedialiveChannel#audio_buffer_model}.
	AudioBufferModel *string `field:"optional" json:"audioBufferModel" yaml:"audioBufferModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#audio_frames_per_pes MedialiveChannel#audio_frames_per_pes}.
	AudioFramesPerPes *float64 `field:"optional" json:"audioFramesPerPes" yaml:"audioFramesPerPes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#audio_pids MedialiveChannel#audio_pids}.
	AudioPids *string `field:"optional" json:"audioPids" yaml:"audioPids"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#audio_stream_type MedialiveChannel#audio_stream_type}.
	AudioStreamType *string `field:"optional" json:"audioStreamType" yaml:"audioStreamType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#bitrate MedialiveChannel#bitrate}.
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#buffer_model MedialiveChannel#buffer_model}.
	BufferModel *string `field:"optional" json:"bufferModel" yaml:"bufferModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#cc_descriptor MedialiveChannel#cc_descriptor}.
	CcDescriptor *string `field:"optional" json:"ccDescriptor" yaml:"ccDescriptor"`
	// dvb_nit_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#dvb_nit_settings MedialiveChannel#dvb_nit_settings}
	DvbNitSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2TsSettingsDvbNitSettings `field:"optional" json:"dvbNitSettings" yaml:"dvbNitSettings"`
	// dvb_sdt_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#dvb_sdt_settings MedialiveChannel#dvb_sdt_settings}
	DvbSdtSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2TsSettingsDvbSdtSettings `field:"optional" json:"dvbSdtSettings" yaml:"dvbSdtSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#dvb_sub_pids MedialiveChannel#dvb_sub_pids}.
	DvbSubPids *string `field:"optional" json:"dvbSubPids" yaml:"dvbSubPids"`
	// dvb_tdt_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#dvb_tdt_settings MedialiveChannel#dvb_tdt_settings}
	DvbTdtSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2TsSettingsDvbTdtSettings `field:"optional" json:"dvbTdtSettings" yaml:"dvbTdtSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#dvb_teletext_pid MedialiveChannel#dvb_teletext_pid}.
	DvbTeletextPid *string `field:"optional" json:"dvbTeletextPid" yaml:"dvbTeletextPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#ebif MedialiveChannel#ebif}.
	Ebif *string `field:"optional" json:"ebif" yaml:"ebif"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#ebp_audio_interval MedialiveChannel#ebp_audio_interval}.
	EbpAudioInterval *string `field:"optional" json:"ebpAudioInterval" yaml:"ebpAudioInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#ebp_lookahead_ms MedialiveChannel#ebp_lookahead_ms}.
	EbpLookaheadMs *float64 `field:"optional" json:"ebpLookaheadMs" yaml:"ebpLookaheadMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#ebp_placement MedialiveChannel#ebp_placement}.
	EbpPlacement *string `field:"optional" json:"ebpPlacement" yaml:"ebpPlacement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#ecm_pid MedialiveChannel#ecm_pid}.
	EcmPid *string `field:"optional" json:"ecmPid" yaml:"ecmPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#es_rate_in_pes MedialiveChannel#es_rate_in_pes}.
	EsRateInPes *string `field:"optional" json:"esRateInPes" yaml:"esRateInPes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#etv_platform_pid MedialiveChannel#etv_platform_pid}.
	EtvPlatformPid *string `field:"optional" json:"etvPlatformPid" yaml:"etvPlatformPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#etv_signal_pid MedialiveChannel#etv_signal_pid}.
	EtvSignalPid *string `field:"optional" json:"etvSignalPid" yaml:"etvSignalPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#fragment_time MedialiveChannel#fragment_time}.
	FragmentTime *float64 `field:"optional" json:"fragmentTime" yaml:"fragmentTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#klv MedialiveChannel#klv}.
	Klv *string `field:"optional" json:"klv" yaml:"klv"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#klv_data_pids MedialiveChannel#klv_data_pids}.
	KlvDataPids *string `field:"optional" json:"klvDataPids" yaml:"klvDataPids"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#nielsen_id3_behavior MedialiveChannel#nielsen_id3_behavior}.
	NielsenId3Behavior *string `field:"optional" json:"nielsenId3Behavior" yaml:"nielsenId3Behavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#null_packet_bitrate MedialiveChannel#null_packet_bitrate}.
	NullPacketBitrate *float64 `field:"optional" json:"nullPacketBitrate" yaml:"nullPacketBitrate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#pat_interval MedialiveChannel#pat_interval}.
	PatInterval *float64 `field:"optional" json:"patInterval" yaml:"patInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#pcr_control MedialiveChannel#pcr_control}.
	PcrControl *string `field:"optional" json:"pcrControl" yaml:"pcrControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#pcr_period MedialiveChannel#pcr_period}.
	PcrPeriod *float64 `field:"optional" json:"pcrPeriod" yaml:"pcrPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#pcr_pid MedialiveChannel#pcr_pid}.
	PcrPid *string `field:"optional" json:"pcrPid" yaml:"pcrPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#pmt_interval MedialiveChannel#pmt_interval}.
	PmtInterval *float64 `field:"optional" json:"pmtInterval" yaml:"pmtInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#pmt_pid MedialiveChannel#pmt_pid}.
	PmtPid *string `field:"optional" json:"pmtPid" yaml:"pmtPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#program_num MedialiveChannel#program_num}.
	ProgramNum *float64 `field:"optional" json:"programNum" yaml:"programNum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#rate_mode MedialiveChannel#rate_mode}.
	RateMode *string `field:"optional" json:"rateMode" yaml:"rateMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#scte27_pids MedialiveChannel#scte27_pids}.
	Scte27Pids *string `field:"optional" json:"scte27Pids" yaml:"scte27Pids"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#scte35_control MedialiveChannel#scte35_control}.
	Scte35Control *string `field:"optional" json:"scte35Control" yaml:"scte35Control"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#scte35_pid MedialiveChannel#scte35_pid}.
	Scte35Pid *string `field:"optional" json:"scte35Pid" yaml:"scte35Pid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#segmentation_markers MedialiveChannel#segmentation_markers}.
	SegmentationMarkers *string `field:"optional" json:"segmentationMarkers" yaml:"segmentationMarkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#segmentation_style MedialiveChannel#segmentation_style}.
	SegmentationStyle *string `field:"optional" json:"segmentationStyle" yaml:"segmentationStyle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#segmentation_time MedialiveChannel#segmentation_time}.
	SegmentationTime *float64 `field:"optional" json:"segmentationTime" yaml:"segmentationTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#timed_metadata_behavior MedialiveChannel#timed_metadata_behavior}.
	TimedMetadataBehavior *string `field:"optional" json:"timedMetadataBehavior" yaml:"timedMetadataBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#timed_metadata_pid MedialiveChannel#timed_metadata_pid}.
	TimedMetadataPid *string `field:"optional" json:"timedMetadataPid" yaml:"timedMetadataPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#transport_stream_id MedialiveChannel#transport_stream_id}.
	TransportStreamId *float64 `field:"optional" json:"transportStreamId" yaml:"transportStreamId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/medialive_channel#video_pid MedialiveChannel#video_pid}.
	VideoPid *string `field:"optional" json:"videoPid" yaml:"videoPid"`
}

