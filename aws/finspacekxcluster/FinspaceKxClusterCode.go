// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package finspacekxcluster


type FinspaceKxClusterCode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/finspace_kx_cluster#s3_bucket FinspaceKxCluster#s3_bucket}.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/finspace_kx_cluster#s3_key FinspaceKxCluster#s3_key}.
	S3Key *string `field:"required" json:"s3Key" yaml:"s3Key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/finspace_kx_cluster#s3_object_version FinspaceKxCluster#s3_object_version}.
	S3ObjectVersion *string `field:"optional" json:"s3ObjectVersion" yaml:"s3ObjectVersion"`
}

