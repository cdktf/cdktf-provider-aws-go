package s3controlstoragelensconfiguration


type S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevel struct {
	// activity_metrics block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3control_storage_lens_configuration#activity_metrics S3ControlStorageLensConfiguration#activity_metrics}
	ActivityMetrics *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelActivityMetrics `field:"optional" json:"activityMetrics" yaml:"activityMetrics"`
	// prefix_level block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3control_storage_lens_configuration#prefix_level S3ControlStorageLensConfiguration#prefix_level}
	PrefixLevel *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelPrefixLevel `field:"optional" json:"prefixLevel" yaml:"prefixLevel"`
}

