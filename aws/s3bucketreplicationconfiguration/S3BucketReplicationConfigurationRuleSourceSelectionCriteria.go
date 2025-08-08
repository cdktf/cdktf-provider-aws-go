// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketreplicationconfiguration


type S3BucketReplicationConfigurationRuleSourceSelectionCriteria struct {
	// replica_modifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/s3_bucket_replication_configuration#replica_modifications S3BucketReplicationConfigurationA#replica_modifications}
	ReplicaModifications *S3BucketReplicationConfigurationRuleSourceSelectionCriteriaReplicaModifications `field:"optional" json:"replicaModifications" yaml:"replicaModifications"`
	// sse_kms_encrypted_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/s3_bucket_replication_configuration#sse_kms_encrypted_objects S3BucketReplicationConfigurationA#sse_kms_encrypted_objects}
	SseKmsEncryptedObjects *S3BucketReplicationConfigurationRuleSourceSelectionCriteriaSseKmsEncryptedObjects `field:"optional" json:"sseKmsEncryptedObjects" yaml:"sseKmsEncryptedObjects"`
}

