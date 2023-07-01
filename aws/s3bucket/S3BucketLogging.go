package s3bucket


type S3BucketLogging struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/s3_bucket#target_bucket S3Bucket#target_bucket}.
	TargetBucket *string `field:"required" json:"targetBucket" yaml:"targetBucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/s3_bucket#target_prefix S3Bucket#target_prefix}.
	TargetPrefix *string `field:"optional" json:"targetPrefix" yaml:"targetPrefix"`
}

