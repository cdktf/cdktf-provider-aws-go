// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package connectinstancestorageconfig


type ConnectInstanceStorageConfigStorageConfigKinesisFirehoseConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/connect_instance_storage_config#firehose_arn ConnectInstanceStorageConfig#firehose_arn}.
	FirehoseArn *string `field:"required" json:"firehoseArn" yaml:"firehoseArn"`
}

