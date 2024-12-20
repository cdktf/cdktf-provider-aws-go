// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configconfigurationrecorder


type ConfigConfigurationRecorderRecordingModeRecordingModeOverride struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/config_configuration_recorder#recording_frequency ConfigConfigurationRecorder#recording_frequency}.
	RecordingFrequency *string `field:"required" json:"recordingFrequency" yaml:"recordingFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/config_configuration_recorder#resource_types ConfigConfigurationRecorder#resource_types}.
	ResourceTypes *[]*string `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/config_configuration_recorder#description ConfigConfigurationRecorder#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

