package s3


type S3BucketObjectLockConfigurationRuleDefaultRetentionA struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_object_lock_configuration#days S3BucketObjectLockConfigurationA#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_object_lock_configuration#mode S3BucketObjectLockConfigurationA#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_object_lock_configuration#years S3BucketObjectLockConfigurationA#years}.
	Years *float64 `field:"optional" json:"years" yaml:"years"`
}

