package s3


type S3BucketReplicationConfigurationRulesSourceSelectionCriteria struct {
	// sse_kms_encrypted_objects block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket#sse_kms_encrypted_objects S3Bucket#sse_kms_encrypted_objects}
	SseKmsEncryptedObjects *S3BucketReplicationConfigurationRulesSourceSelectionCriteriaSseKmsEncryptedObjects `field:"optional" json:"sseKmsEncryptedObjects" yaml:"sseKmsEncryptedObjects"`
}

