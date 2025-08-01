// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ivsrecordingconfiguration


type IvsRecordingConfigurationDestinationConfigurationS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ivs_recording_configuration#bucket_name IvsRecordingConfiguration#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
}

