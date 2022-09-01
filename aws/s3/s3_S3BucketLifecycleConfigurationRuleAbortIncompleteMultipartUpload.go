package s3


type S3BucketLifecycleConfigurationRuleAbortIncompleteMultipartUpload struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_lifecycle_configuration#days_after_initiation S3BucketLifecycleConfiguration#days_after_initiation}.
	DaysAfterInitiation *float64 `field:"optional" json:"daysAfterInitiation" yaml:"daysAfterInitiation"`
}

