// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketreplicationconfiguration


type S3BucketReplicationConfigurationRuleDestinationAccessControlTranslation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/s3_bucket_replication_configuration#owner S3BucketReplicationConfigurationA#owner}.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
}

