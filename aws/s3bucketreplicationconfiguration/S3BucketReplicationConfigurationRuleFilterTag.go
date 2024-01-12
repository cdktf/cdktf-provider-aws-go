// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketreplicationconfiguration


type S3BucketReplicationConfigurationRuleFilterTag struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/s3_bucket_replication_configuration#key S3BucketReplicationConfigurationA#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/s3_bucket_replication_configuration#value S3BucketReplicationConfigurationA#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

