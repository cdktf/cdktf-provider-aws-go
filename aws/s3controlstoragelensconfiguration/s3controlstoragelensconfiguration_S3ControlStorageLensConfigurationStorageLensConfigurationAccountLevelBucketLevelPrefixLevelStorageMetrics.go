package s3controlstoragelensconfiguration


type S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetrics struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3control_storage_lens_configuration#enabled S3ControlStorageLensConfiguration#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// selection_criteria block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3control_storage_lens_configuration#selection_criteria S3ControlStorageLensConfiguration#selection_criteria}
	SelectionCriteria *S3ControlStorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsSelectionCriteria `field:"optional" json:"selectionCriteria" yaml:"selectionCriteria"`
}

