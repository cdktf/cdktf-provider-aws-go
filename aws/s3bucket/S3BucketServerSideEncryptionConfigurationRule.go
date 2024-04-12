// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketServerSideEncryptionConfigurationRule struct {
	// apply_server_side_encryption_by_default block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/s3_bucket#apply_server_side_encryption_by_default S3Bucket#apply_server_side_encryption_by_default}
	ApplyServerSideEncryptionByDefault *S3BucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault `field:"required" json:"applyServerSideEncryptionByDefault" yaml:"applyServerSideEncryptionByDefault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/s3_bucket#bucket_key_enabled S3Bucket#bucket_key_enabled}.
	BucketKeyEnabled interface{} `field:"optional" json:"bucketKeyEnabled" yaml:"bucketKeyEnabled"`
}

