// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package connectinstancestorageconfig


type ConnectInstanceStorageConfigStorageConfigKinesisStreamConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/connect_instance_storage_config#stream_arn ConnectInstanceStorageConfig#stream_arn}.
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
}

