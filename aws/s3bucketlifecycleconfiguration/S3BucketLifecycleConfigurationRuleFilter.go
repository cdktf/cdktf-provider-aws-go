// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketlifecycleconfiguration


type S3BucketLifecycleConfigurationRuleFilter struct {
	// and block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/s3_bucket_lifecycle_configuration#and S3BucketLifecycleConfiguration#and}
	And *S3BucketLifecycleConfigurationRuleFilterAnd `field:"optional" json:"and" yaml:"and"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/s3_bucket_lifecycle_configuration#object_size_greater_than S3BucketLifecycleConfiguration#object_size_greater_than}.
	ObjectSizeGreaterThan *string `field:"optional" json:"objectSizeGreaterThan" yaml:"objectSizeGreaterThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/s3_bucket_lifecycle_configuration#object_size_less_than S3BucketLifecycleConfiguration#object_size_less_than}.
	ObjectSizeLessThan *string `field:"optional" json:"objectSizeLessThan" yaml:"objectSizeLessThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/s3_bucket_lifecycle_configuration#prefix S3BucketLifecycleConfiguration#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/s3_bucket_lifecycle_configuration#tag S3BucketLifecycleConfiguration#tag}
	Tag *S3BucketLifecycleConfigurationRuleFilterTag `field:"optional" json:"tag" yaml:"tag"`
}

