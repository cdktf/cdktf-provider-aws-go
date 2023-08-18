package s3bucketlogging


type S3BucketLoggingTargetGrantGrantee struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/s3_bucket_logging#type S3BucketLoggingA#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/s3_bucket_logging#email_address S3BucketLoggingA#email_address}.
	EmailAddress *string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/s3_bucket_logging#id S3BucketLoggingA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/s3_bucket_logging#uri S3BucketLoggingA#uri}.
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

