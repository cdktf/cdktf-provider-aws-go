package s3bucketinventory


type S3BucketInventoryDestinationBucketEncryptionSseKms struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_inventory#key_id S3BucketInventory#key_id}.
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
}

