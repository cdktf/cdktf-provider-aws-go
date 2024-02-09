// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketlogging


type S3BucketLoggingTargetObjectKeyFormatPartitionedPrefix struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/s3_bucket_logging#partition_date_source S3BucketLoggingA#partition_date_source}.
	PartitionDateSource *string `field:"required" json:"partitionDateSource" yaml:"partitionDateSource"`
}

