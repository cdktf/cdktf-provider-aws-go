package s3


type S3BucketInventoryDestinationBucketEncryption struct {
	// sse_kms block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_inventory#sse_kms S3BucketInventory#sse_kms}
	SseKms *S3BucketInventoryDestinationBucketEncryptionSseKms `field:"optional" json:"sseKms" yaml:"sseKms"`
	// sse_s3 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_inventory#sse_s3 S3BucketInventory#sse_s3}
	SseS3 *S3BucketInventoryDestinationBucketEncryptionSseS3 `field:"optional" json:"sseS3" yaml:"sseS3"`
}

