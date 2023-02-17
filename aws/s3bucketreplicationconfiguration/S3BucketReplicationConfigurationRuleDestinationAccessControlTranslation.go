package s3bucketreplicationconfiguration


type S3BucketReplicationConfigurationRuleDestinationAccessControlTranslation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_replication_configuration#owner S3BucketReplicationConfigurationA#owner}.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
}

