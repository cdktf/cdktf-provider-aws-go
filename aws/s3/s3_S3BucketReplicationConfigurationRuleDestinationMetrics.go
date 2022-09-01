package s3


type S3BucketReplicationConfigurationRuleDestinationMetrics struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_replication_configuration#status S3BucketReplicationConfigurationA#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// event_threshold block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_replication_configuration#event_threshold S3BucketReplicationConfigurationA#event_threshold}
	EventThreshold *S3BucketReplicationConfigurationRuleDestinationMetricsEventThreshold `field:"optional" json:"eventThreshold" yaml:"eventThreshold"`
}

