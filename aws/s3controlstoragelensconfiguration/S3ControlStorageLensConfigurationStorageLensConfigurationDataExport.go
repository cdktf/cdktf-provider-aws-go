package s3controlstoragelensconfiguration


type S3ControlStorageLensConfigurationStorageLensConfigurationDataExport struct {
	// cloud_watch_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/s3control_storage_lens_configuration#cloud_watch_metrics S3ControlStorageLensConfiguration#cloud_watch_metrics}
	CloudWatchMetrics *S3ControlStorageLensConfigurationStorageLensConfigurationDataExportCloudWatchMetrics `field:"optional" json:"cloudWatchMetrics" yaml:"cloudWatchMetrics"`
	// s3_bucket_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/s3control_storage_lens_configuration#s3_bucket_destination S3ControlStorageLensConfiguration#s3_bucket_destination}
	S3BucketDestination *S3ControlStorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestination `field:"optional" json:"s3BucketDestination" yaml:"s3BucketDestination"`
}

