package s3bucketinventory


type S3BucketInventorySchedule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_inventory#frequency S3BucketInventory#frequency}.
	Frequency *string `field:"required" json:"frequency" yaml:"frequency"`
}

