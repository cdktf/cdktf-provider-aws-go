package s3


type S3BucketReplicationConfigurationRuleSourceSelectionCriteriaReplicaModifications struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_replication_configuration#status S3BucketReplicationConfigurationA#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
}
