package s3


type S3BucketReplicationConfigurationRuleDestinationMetricsEventThreshold struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_replication_configuration#minutes S3BucketReplicationConfigurationA#minutes}.
	Minutes *float64 `field:"required" json:"minutes" yaml:"minutes"`
}

