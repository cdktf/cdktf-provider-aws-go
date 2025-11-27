// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketmetadataconfiguration


type S3BucketMetadataConfigurationMetadataConfigurationJournalTableConfiguration struct {
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/s3_bucket_metadata_configuration#encryption_configuration S3BucketMetadataConfiguration#encryption_configuration}
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// record_expiration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/s3_bucket_metadata_configuration#record_expiration S3BucketMetadataConfiguration#record_expiration}
	RecordExpiration interface{} `field:"optional" json:"recordExpiration" yaml:"recordExpiration"`
}

