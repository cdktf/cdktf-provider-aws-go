package s3bucketreplicationconfiguration


type S3BucketReplicationConfigurationRuleDestinationEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/s3_bucket_replication_configuration#replica_kms_key_id S3BucketReplicationConfigurationA#replica_kms_key_id}.
	ReplicaKmsKeyId *string `field:"required" json:"replicaKmsKeyId" yaml:"replicaKmsKeyId"`
}

