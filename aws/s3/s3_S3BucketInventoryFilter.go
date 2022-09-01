package s3


type S3BucketInventoryFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_inventory#prefix S3BucketInventory#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

