// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketmetadataconfiguration


type S3BucketMetadataConfigurationMetadataConfigurationInventoryTableConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/s3_bucket_metadata_configuration#configuration_state S3BucketMetadataConfiguration#configuration_state}.
	ConfigurationState *string `field:"required" json:"configurationState" yaml:"configurationState"`
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/s3_bucket_metadata_configuration#encryption_configuration S3BucketMetadataConfiguration#encryption_configuration}
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
}

