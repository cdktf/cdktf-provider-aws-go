package s3bucket


type S3BucketObjectLockConfigurationRule struct {
	// default_retention block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket#default_retention S3Bucket#default_retention}
	DefaultRetention *S3BucketObjectLockConfigurationRuleDefaultRetention `field:"required" json:"defaultRetention" yaml:"defaultRetention"`
}

