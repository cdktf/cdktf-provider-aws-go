package s3


type S3BucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultA struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_server_side_encryption_configuration#sse_algorithm S3BucketServerSideEncryptionConfigurationA#sse_algorithm}.
	SseAlgorithm *string `field:"required" json:"sseAlgorithm" yaml:"sseAlgorithm"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3_bucket_server_side_encryption_configuration#kms_master_key_id S3BucketServerSideEncryptionConfigurationA#kms_master_key_id}.
	KmsMasterKeyId *string `field:"optional" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
}

