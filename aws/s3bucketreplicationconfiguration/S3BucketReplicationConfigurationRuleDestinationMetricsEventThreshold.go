package s3bucketreplicationconfiguration


type S3BucketReplicationConfigurationRuleDestinationMetricsEventThreshold struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/s3_bucket_replication_configuration#minutes S3BucketReplicationConfigurationA#minutes}.
	Minutes *float64 `field:"required" json:"minutes" yaml:"minutes"`
}

