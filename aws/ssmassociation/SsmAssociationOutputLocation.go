// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmassociation


type SsmAssociationOutputLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/ssm_association#s3_bucket_name SsmAssociation#s3_bucket_name}.
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/ssm_association#s3_key_prefix SsmAssociation#s3_key_prefix}.
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/ssm_association#s3_region SsmAssociation#s3_region}.
	S3Region *string `field:"optional" json:"s3Region" yaml:"s3Region"`
}

