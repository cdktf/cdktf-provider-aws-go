package s3


type S3BucketReplicationConfigurationRuleDestinationEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_replication_configuration#replica_kms_key_id S3BucketReplicationConfigurationA#replica_kms_key_id}.
	ReplicaKmsKeyId *string `field:"required" json:"replicaKmsKeyId" yaml:"replicaKmsKeyId"`
}

