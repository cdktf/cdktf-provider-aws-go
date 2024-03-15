// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configconfigurationrecorder


type ConfigConfigurationRecorderRecordingMode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/config_configuration_recorder#recording_frequency ConfigConfigurationRecorder#recording_frequency}.
	RecordingFrequency *string `field:"optional" json:"recordingFrequency" yaml:"recordingFrequency"`
	// recording_mode_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/config_configuration_recorder#recording_mode_override ConfigConfigurationRecorder#recording_mode_override}
	RecordingModeOverride *ConfigConfigurationRecorderRecordingModeRecordingModeOverride `field:"optional" json:"recordingModeOverride" yaml:"recordingModeOverride"`
}

