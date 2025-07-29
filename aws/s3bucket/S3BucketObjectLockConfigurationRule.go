// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketObjectLockConfigurationRule struct {
	// default_retention block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/s3_bucket#default_retention S3Bucket#default_retention}
	DefaultRetention *S3BucketObjectLockConfigurationRuleDefaultRetention `field:"required" json:"defaultRetention" yaml:"defaultRetention"`
}

