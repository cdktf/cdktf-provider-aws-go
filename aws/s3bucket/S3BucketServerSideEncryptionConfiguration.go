package s3bucket


type S3BucketServerSideEncryptionConfiguration struct {
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/s3_bucket#rule S3Bucket#rule}
	Rule *S3BucketServerSideEncryptionConfigurationRule `field:"required" json:"rule" yaml:"rule"`
}

