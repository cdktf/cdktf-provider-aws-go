// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketReplicationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/s3_bucket#role S3Bucket#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/s3_bucket#rules S3Bucket#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

