package s3bucketinventory


type S3BucketInventoryDestination struct {
	// bucket block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_inventory#bucket S3BucketInventory#bucket}
	Bucket *S3BucketInventoryDestinationBucket `field:"required" json:"bucket" yaml:"bucket"`
}

