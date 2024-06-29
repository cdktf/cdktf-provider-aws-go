// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3controlmultiregionaccesspoint


type S3ControlMultiRegionAccessPointDetailsRegion struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/s3control_multi_region_access_point#bucket S3ControlMultiRegionAccessPoint#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/s3control_multi_region_access_point#bucket_account_id S3ControlMultiRegionAccessPoint#bucket_account_id}.
	BucketAccountId *string `field:"optional" json:"bucketAccountId" yaml:"bucketAccountId"`
}

