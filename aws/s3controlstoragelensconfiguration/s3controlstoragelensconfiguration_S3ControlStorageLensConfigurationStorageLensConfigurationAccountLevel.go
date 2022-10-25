package s3controlstoragelensconfiguration


type S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevel struct {
	// bucket_level block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3control_storage_lens_configuration#bucket_level S3ControlStorageLensConfiguration#bucket_level}
	BucketLevel *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevel `field:"required" json:"bucketLevel" yaml:"bucketLevel"`
	// activity_metrics block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3control_storage_lens_configuration#activity_metrics S3ControlStorageLensConfiguration#activity_metrics}
	ActivityMetrics *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelActivityMetrics `field:"optional" json:"activityMetrics" yaml:"activityMetrics"`
}

