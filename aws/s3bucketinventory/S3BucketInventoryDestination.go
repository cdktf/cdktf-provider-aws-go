// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketinventory


type S3BucketInventoryDestination struct {
	// bucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/s3_bucket_inventory#bucket S3BucketInventory#bucket}
	Bucket *S3BucketInventoryDestinationBucket `field:"required" json:"bucket" yaml:"bucket"`
}

