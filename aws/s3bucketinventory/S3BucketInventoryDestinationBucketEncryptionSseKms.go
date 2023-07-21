package s3bucketinventory


type S3BucketInventoryDestinationBucketEncryptionSseKms struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/s3_bucket_inventory#key_id S3BucketInventory#key_id}.
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
}

