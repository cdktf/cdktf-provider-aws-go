// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketmetadataconfiguration


type S3BucketMetadataConfigurationMetadataConfigurationJournalTableConfigurationRecordExpiration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/s3_bucket_metadata_configuration#expiration S3BucketMetadataConfiguration#expiration}.
	Expiration *string `field:"required" json:"expiration" yaml:"expiration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/s3_bucket_metadata_configuration#days S3BucketMetadataConfiguration#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

