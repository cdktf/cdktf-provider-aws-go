// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketreplicationconfiguration


type S3BucketReplicationConfigurationRuleFilter struct {
	// and block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/s3_bucket_replication_configuration#and S3BucketReplicationConfigurationA#and}
	And *S3BucketReplicationConfigurationRuleFilterAnd `field:"optional" json:"and" yaml:"and"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/s3_bucket_replication_configuration#prefix S3BucketReplicationConfigurationA#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/s3_bucket_replication_configuration#tag S3BucketReplicationConfigurationA#tag}
	Tag *S3BucketReplicationConfigurationRuleFilterTag `field:"optional" json:"tag" yaml:"tag"`
}

