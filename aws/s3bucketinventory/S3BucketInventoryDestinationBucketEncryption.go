// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketinventory


type S3BucketInventoryDestinationBucketEncryption struct {
	// sse_kms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.1/docs/resources/s3_bucket_inventory#sse_kms S3BucketInventory#sse_kms}
	SseKms *S3BucketInventoryDestinationBucketEncryptionSseKms `field:"optional" json:"sseKms" yaml:"sseKms"`
	// sse_s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.1/docs/resources/s3_bucket_inventory#sse_s3 S3BucketInventory#sse_s3}
	SseS3 *S3BucketInventoryDestinationBucketEncryptionSseS3 `field:"optional" json:"sseS3" yaml:"sseS3"`
}

