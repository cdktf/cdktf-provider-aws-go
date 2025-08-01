// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketlogging


type S3BucketLoggingTargetGrant struct {
	// grantee block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/s3_bucket_logging#grantee S3BucketLoggingA#grantee}
	Grantee *S3BucketLoggingTargetGrantGrantee `field:"required" json:"grantee" yaml:"grantee"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/s3_bucket_logging#permission S3BucketLoggingA#permission}.
	Permission *string `field:"required" json:"permission" yaml:"permission"`
}

