// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketlifecycleconfiguration


type S3BucketLifecycleConfigurationRuleNoncurrentVersionExpiration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/s3_bucket_lifecycle_configuration#noncurrent_days S3BucketLifecycleConfiguration#noncurrent_days}.
	NoncurrentDays *float64 `field:"required" json:"noncurrentDays" yaml:"noncurrentDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/s3_bucket_lifecycle_configuration#newer_noncurrent_versions S3BucketLifecycleConfiguration#newer_noncurrent_versions}.
	NewerNoncurrentVersions *float64 `field:"optional" json:"newerNoncurrentVersions" yaml:"newerNoncurrentVersions"`
}

