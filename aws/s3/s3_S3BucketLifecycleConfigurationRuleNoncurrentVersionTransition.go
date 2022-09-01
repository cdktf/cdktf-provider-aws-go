package s3


type S3BucketLifecycleConfigurationRuleNoncurrentVersionTransition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_lifecycle_configuration#storage_class S3BucketLifecycleConfiguration#storage_class}.
	StorageClass *string `field:"required" json:"storageClass" yaml:"storageClass"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_lifecycle_configuration#newer_noncurrent_versions S3BucketLifecycleConfiguration#newer_noncurrent_versions}.
	NewerNoncurrentVersions *string `field:"optional" json:"newerNoncurrentVersions" yaml:"newerNoncurrentVersions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_lifecycle_configuration#noncurrent_days S3BucketLifecycleConfiguration#noncurrent_days}.
	NoncurrentDays *float64 `field:"optional" json:"noncurrentDays" yaml:"noncurrentDays"`
}

