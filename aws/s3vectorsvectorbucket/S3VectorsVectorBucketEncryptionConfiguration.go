// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3vectorsvectorbucket


type S3VectorsVectorBucketEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/s3vectors_vector_bucket#kms_key_arn S3VectorsVectorBucket#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/s3vectors_vector_bucket#sse_type S3VectorsVectorBucket#sse_type}.
	SseType *string `field:"optional" json:"sseType" yaml:"sseType"`
}

