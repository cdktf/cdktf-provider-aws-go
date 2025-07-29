// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettings struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/medialive_channel#destination MedialiveChannel#destination}
	Destination *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestination `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/medialive_channel#acquisition_point_id MedialiveChannel#acquisition_point_id}.
	AcquisitionPointId *string `field:"optional" json:"acquisitionPointId" yaml:"acquisitionPointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/medialive_channel#audio_only_timecode_control MedialiveChannel#audio_only_timecode_control}.
	AudioOnlyTimecodeControl *string `field:"optional" json:"audioOnlyTimecodeControl" yaml:"audioOnlyTimecodeControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/medialive_channel#certificate_mode MedialiveChannel#certificate_mode}.
	CertificateMode *string `field:"optional" json:"certificateMode" yaml:"certificateMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/medialive_channel#connection_retry_interval MedialiveChannel#connection_retry_interval}.
	ConnectionRetryInterval *float64 `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/medialive_channel#event_id MedialiveChannel#event_id}.
	EventId *string `field:"optional" json:"eventId" yaml:"eventId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/medialive_channel#event_id_mode MedialiveChannel#event_id_mode}.
	EventIdMode *string `field:"optional" json:"eventIdMode" yaml:"eventIdMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/medialive_channel#event_stop_behavior MedialiveChannel#event_stop_behavior}.
	EventStopBehavior *string `field:"optional" json:"eventStopBehavior" yaml:"eventStopBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/medialive_channel#filecache_duration MedialiveChannel#filecache_duration}.
	FilecacheDuration *float64 `field:"optional" json:"filecacheDuration" yaml:"filecacheDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/medialive_channel#fragment_length MedialiveChannel#fragment_length}.
	FragmentLength *float64 `field:"optional" json:"fragmentLength" yaml:"fragmentLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/medialive_channel#input_loss_action MedialiveChannel#input_loss_action}.
	InputLossAction *string `field:"optional" json:"inputLossAction" yaml:"inputLossAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/medialive_channel#num_retries MedialiveChannel#num_retries}.
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/medialive_channel#restart_delay MedialiveChannel#restart_delay}.
	RestartDelay *float64 `field:"optional" json:"restartDelay" yaml:"restartDelay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/medialive_channel#segmentation_mode MedialiveChannel#segmentation_mode}.
	SegmentationMode *string `field:"optional" json:"segmentationMode" yaml:"segmentationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/medialive_channel#send_delay_ms MedialiveChannel#send_delay_ms}.
	SendDelayMs *float64 `field:"optional" json:"sendDelayMs" yaml:"sendDelayMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/medialive_channel#sparse_track_type MedialiveChannel#sparse_track_type}.
	SparseTrackType *string `field:"optional" json:"sparseTrackType" yaml:"sparseTrackType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/medialive_channel#stream_manifest_behavior MedialiveChannel#stream_manifest_behavior}.
	StreamManifestBehavior *string `field:"optional" json:"streamManifestBehavior" yaml:"streamManifestBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/medialive_channel#timestamp_offset MedialiveChannel#timestamp_offset}.
	TimestampOffset *string `field:"optional" json:"timestampOffset" yaml:"timestampOffset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/medialive_channel#timestamp_offset_mode MedialiveChannel#timestamp_offset_mode}.
	TimestampOffsetMode *string `field:"optional" json:"timestampOffsetMode" yaml:"timestampOffsetMode"`
}

