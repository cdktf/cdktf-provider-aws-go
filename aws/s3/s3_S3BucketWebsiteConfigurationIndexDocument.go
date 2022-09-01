package s3


type S3BucketWebsiteConfigurationIndexDocument struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_website_configuration#suffix S3BucketWebsiteConfiguration#suffix}.
	Suffix *string `field:"required" json:"suffix" yaml:"suffix"`
}

