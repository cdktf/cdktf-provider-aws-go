// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketlifecycleconfiguration


type S3BucketLifecycleConfigurationRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/s3_bucket_lifecycle_configuration#id S3BucketLifecycleConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/s3_bucket_lifecycle_configuration#status S3BucketLifecycleConfiguration#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// abort_incomplete_multipart_upload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/s3_bucket_lifecycle_configuration#abort_incomplete_multipart_upload S3BucketLifecycleConfiguration#abort_incomplete_multipart_upload}
	AbortIncompleteMultipartUpload *S3BucketLifecycleConfigurationRuleAbortIncompleteMultipartUpload `field:"optional" json:"abortIncompleteMultipartUpload" yaml:"abortIncompleteMultipartUpload"`
	// expiration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/s3_bucket_lifecycle_configuration#expiration S3BucketLifecycleConfiguration#expiration}
	Expiration *S3BucketLifecycleConfigurationRuleExpiration `field:"optional" json:"expiration" yaml:"expiration"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/s3_bucket_lifecycle_configuration#filter S3BucketLifecycleConfiguration#filter}
	Filter *S3BucketLifecycleConfigurationRuleFilter `field:"optional" json:"filter" yaml:"filter"`
	// noncurrent_version_expiration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/s3_bucket_lifecycle_configuration#noncurrent_version_expiration S3BucketLifecycleConfiguration#noncurrent_version_expiration}
	NoncurrentVersionExpiration *S3BucketLifecycleConfigurationRuleNoncurrentVersionExpiration `field:"optional" json:"noncurrentVersionExpiration" yaml:"noncurrentVersionExpiration"`
	// noncurrent_version_transition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/s3_bucket_lifecycle_configuration#noncurrent_version_transition S3BucketLifecycleConfiguration#noncurrent_version_transition}
	NoncurrentVersionTransition interface{} `field:"optional" json:"noncurrentVersionTransition" yaml:"noncurrentVersionTransition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/s3_bucket_lifecycle_configuration#prefix S3BucketLifecycleConfiguration#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// transition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/s3_bucket_lifecycle_configuration#transition S3BucketLifecycleConfiguration#transition}
	Transition interface{} `field:"optional" json:"transition" yaml:"transition"`
}

