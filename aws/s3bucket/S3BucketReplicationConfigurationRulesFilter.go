// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketReplicationConfigurationRulesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/s3_bucket#prefix S3Bucket#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/s3_bucket#tags S3Bucket#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

