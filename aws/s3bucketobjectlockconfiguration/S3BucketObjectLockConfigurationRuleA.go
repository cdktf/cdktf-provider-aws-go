// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketobjectlockconfiguration


type S3BucketObjectLockConfigurationRuleA struct {
	// default_retention block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/s3_bucket_object_lock_configuration#default_retention S3BucketObjectLockConfigurationA#default_retention}
	DefaultRetention *S3BucketObjectLockConfigurationRuleDefaultRetentionA `field:"required" json:"defaultRetention" yaml:"defaultRetention"`
}

