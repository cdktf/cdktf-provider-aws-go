// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketownershipcontrols


type S3BucketOwnershipControlsRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/s3_bucket_ownership_controls#object_ownership S3BucketOwnershipControls#object_ownership}.
	ObjectOwnership *string `field:"required" json:"objectOwnership" yaml:"objectOwnership"`
}

