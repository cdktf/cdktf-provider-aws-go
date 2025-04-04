// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketObjectLockConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/s3_bucket#object_lock_enabled S3Bucket#object_lock_enabled}.
	ObjectLockEnabled *string `field:"optional" json:"objectLockEnabled" yaml:"objectLockEnabled"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/s3_bucket#rule S3Bucket#rule}
	Rule *S3BucketObjectLockConfigurationRule `field:"optional" json:"rule" yaml:"rule"`
}

