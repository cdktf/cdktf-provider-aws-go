// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketReplicationConfigurationRulesDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.96.0/docs/resources/s3_bucket#bucket S3Bucket#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// access_control_translation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.96.0/docs/resources/s3_bucket#access_control_translation S3Bucket#access_control_translation}
	AccessControlTranslation *S3BucketReplicationConfigurationRulesDestinationAccessControlTranslation `field:"optional" json:"accessControlTranslation" yaml:"accessControlTranslation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.96.0/docs/resources/s3_bucket#account_id S3Bucket#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.96.0/docs/resources/s3_bucket#metrics S3Bucket#metrics}
	Metrics *S3BucketReplicationConfigurationRulesDestinationMetrics `field:"optional" json:"metrics" yaml:"metrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.96.0/docs/resources/s3_bucket#replica_kms_key_id S3Bucket#replica_kms_key_id}.
	ReplicaKmsKeyId *string `field:"optional" json:"replicaKmsKeyId" yaml:"replicaKmsKeyId"`
	// replication_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.96.0/docs/resources/s3_bucket#replication_time S3Bucket#replication_time}
	ReplicationTime *S3BucketReplicationConfigurationRulesDestinationReplicationTime `field:"optional" json:"replicationTime" yaml:"replicationTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.96.0/docs/resources/s3_bucket#storage_class S3Bucket#storage_class}.
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

