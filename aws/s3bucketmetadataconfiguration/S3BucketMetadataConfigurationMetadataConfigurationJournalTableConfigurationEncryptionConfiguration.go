// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketmetadataconfiguration


type S3BucketMetadataConfigurationMetadataConfigurationJournalTableConfigurationEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/s3_bucket_metadata_configuration#sse_algorithm S3BucketMetadataConfiguration#sse_algorithm}.
	SseAlgorithm *string `field:"required" json:"sseAlgorithm" yaml:"sseAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/s3_bucket_metadata_configuration#kms_key_arn S3BucketMetadataConfiguration#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

