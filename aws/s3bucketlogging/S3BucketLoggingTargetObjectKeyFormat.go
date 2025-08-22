// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketlogging


type S3BucketLoggingTargetObjectKeyFormat struct {
	// partitioned_prefix block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/s3_bucket_logging#partitioned_prefix S3BucketLoggingA#partitioned_prefix}
	PartitionedPrefix *S3BucketLoggingTargetObjectKeyFormatPartitionedPrefix `field:"optional" json:"partitionedPrefix" yaml:"partitionedPrefix"`
	// simple_prefix block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/s3_bucket_logging#simple_prefix S3BucketLoggingA#simple_prefix}
	SimplePrefix *S3BucketLoggingTargetObjectKeyFormatSimplePrefix `field:"optional" json:"simplePrefix" yaml:"simplePrefix"`
}

