package s3bucketlifecycleconfiguration


type S3BucketLifecycleConfigurationRuleTransition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/s3_bucket_lifecycle_configuration#storage_class S3BucketLifecycleConfiguration#storage_class}.
	StorageClass *string `field:"required" json:"storageClass" yaml:"storageClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/s3_bucket_lifecycle_configuration#date S3BucketLifecycleConfiguration#date}.
	Date *string `field:"optional" json:"date" yaml:"date"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/s3_bucket_lifecycle_configuration#days S3BucketLifecycleConfiguration#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

