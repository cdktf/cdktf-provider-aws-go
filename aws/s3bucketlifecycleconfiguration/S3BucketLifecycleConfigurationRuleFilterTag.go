// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketlifecycleconfiguration


type S3BucketLifecycleConfigurationRuleFilterTag struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/s3_bucket_lifecycle_configuration#key S3BucketLifecycleConfiguration#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/s3_bucket_lifecycle_configuration#value S3BucketLifecycleConfiguration#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

