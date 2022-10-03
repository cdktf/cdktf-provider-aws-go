package s3bucketownershipcontrols


type S3BucketOwnershipControlsRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_ownership_controls#object_ownership S3BucketOwnershipControls#object_ownership}.
	ObjectOwnership *string `field:"required" json:"objectOwnership" yaml:"objectOwnership"`
}

