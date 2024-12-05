// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ivsrecordingconfiguration


type IvsRecordingConfigurationDestinationConfiguration struct {
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/ivs_recording_configuration#s3 IvsRecordingConfiguration#s3}
	S3 *IvsRecordingConfigurationDestinationConfigurationS3 `field:"required" json:"s3" yaml:"s3"`
}

