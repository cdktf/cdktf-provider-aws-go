// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3accesspoint


type S3AccessPointPublicAccessBlockConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/s3_access_point#block_public_acls S3AccessPoint#block_public_acls}.
	BlockPublicAcls interface{} `field:"optional" json:"blockPublicAcls" yaml:"blockPublicAcls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/s3_access_point#block_public_policy S3AccessPoint#block_public_policy}.
	BlockPublicPolicy interface{} `field:"optional" json:"blockPublicPolicy" yaml:"blockPublicPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/s3_access_point#ignore_public_acls S3AccessPoint#ignore_public_acls}.
	IgnorePublicAcls interface{} `field:"optional" json:"ignorePublicAcls" yaml:"ignorePublicAcls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/s3_access_point#restrict_public_buckets S3AccessPoint#restrict_public_buckets}.
	RestrictPublicBuckets interface{} `field:"optional" json:"restrictPublicBuckets" yaml:"restrictPublicBuckets"`
}

