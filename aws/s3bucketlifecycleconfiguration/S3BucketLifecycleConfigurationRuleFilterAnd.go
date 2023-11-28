// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketlifecycleconfiguration


type S3BucketLifecycleConfigurationRuleFilterAnd struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/s3_bucket_lifecycle_configuration#object_size_greater_than S3BucketLifecycleConfiguration#object_size_greater_than}.
	ObjectSizeGreaterThan *float64 `field:"optional" json:"objectSizeGreaterThan" yaml:"objectSizeGreaterThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/s3_bucket_lifecycle_configuration#object_size_less_than S3BucketLifecycleConfiguration#object_size_less_than}.
	ObjectSizeLessThan *float64 `field:"optional" json:"objectSizeLessThan" yaml:"objectSizeLessThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/s3_bucket_lifecycle_configuration#prefix S3BucketLifecycleConfiguration#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/s3_bucket_lifecycle_configuration#tags S3BucketLifecycleConfiguration#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

