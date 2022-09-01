package s3


type S3BucketServerSideEncryptionConfiguration struct {
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket#rule S3Bucket#rule}
	Rule *S3BucketServerSideEncryptionConfigurationRule `field:"required" json:"rule" yaml:"rule"`
}

