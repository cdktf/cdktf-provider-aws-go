// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketLifecycleRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/s3_bucket#enabled S3Bucket#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/s3_bucket#abort_incomplete_multipart_upload_days S3Bucket#abort_incomplete_multipart_upload_days}.
	AbortIncompleteMultipartUploadDays *float64 `field:"optional" json:"abortIncompleteMultipartUploadDays" yaml:"abortIncompleteMultipartUploadDays"`
	// expiration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/s3_bucket#expiration S3Bucket#expiration}
	Expiration *S3BucketLifecycleRuleExpiration `field:"optional" json:"expiration" yaml:"expiration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/s3_bucket#id S3Bucket#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// noncurrent_version_expiration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/s3_bucket#noncurrent_version_expiration S3Bucket#noncurrent_version_expiration}
	NoncurrentVersionExpiration *S3BucketLifecycleRuleNoncurrentVersionExpiration `field:"optional" json:"noncurrentVersionExpiration" yaml:"noncurrentVersionExpiration"`
	// noncurrent_version_transition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/s3_bucket#noncurrent_version_transition S3Bucket#noncurrent_version_transition}
	NoncurrentVersionTransition interface{} `field:"optional" json:"noncurrentVersionTransition" yaml:"noncurrentVersionTransition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/s3_bucket#prefix S3Bucket#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/s3_bucket#tags S3Bucket#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// transition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/s3_bucket#transition S3Bucket#transition}
	Transition interface{} `field:"optional" json:"transition" yaml:"transition"`
}

