package s3


type S3BucketLifecycleConfigurationRuleFilterTag struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_lifecycle_configuration#key S3BucketLifecycleConfiguration#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_lifecycle_configuration#value S3BucketLifecycleConfiguration#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

