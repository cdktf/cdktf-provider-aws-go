// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package connectinstancestorageconfig


type ConnectInstanceStorageConfigStorageConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.0/docs/resources/connect_instance_storage_config#storage_type ConnectInstanceStorageConfig#storage_type}.
	StorageType *string `field:"required" json:"storageType" yaml:"storageType"`
	// kinesis_firehose_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.0/docs/resources/connect_instance_storage_config#kinesis_firehose_config ConnectInstanceStorageConfig#kinesis_firehose_config}
	KinesisFirehoseConfig *ConnectInstanceStorageConfigStorageConfigKinesisFirehoseConfig `field:"optional" json:"kinesisFirehoseConfig" yaml:"kinesisFirehoseConfig"`
	// kinesis_stream_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.0/docs/resources/connect_instance_storage_config#kinesis_stream_config ConnectInstanceStorageConfig#kinesis_stream_config}
	KinesisStreamConfig *ConnectInstanceStorageConfigStorageConfigKinesisStreamConfig `field:"optional" json:"kinesisStreamConfig" yaml:"kinesisStreamConfig"`
	// kinesis_video_stream_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.0/docs/resources/connect_instance_storage_config#kinesis_video_stream_config ConnectInstanceStorageConfig#kinesis_video_stream_config}
	KinesisVideoStreamConfig *ConnectInstanceStorageConfigStorageConfigKinesisVideoStreamConfig `field:"optional" json:"kinesisVideoStreamConfig" yaml:"kinesisVideoStreamConfig"`
	// s3_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.0/docs/resources/connect_instance_storage_config#s3_config ConnectInstanceStorageConfig#s3_config}
	S3Config *ConnectInstanceStorageConfigStorageConfigS3Config `field:"optional" json:"s3Config" yaml:"s3Config"`
}

