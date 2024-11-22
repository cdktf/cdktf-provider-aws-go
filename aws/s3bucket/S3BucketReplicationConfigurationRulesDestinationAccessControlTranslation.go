// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketReplicationConfigurationRulesDestinationAccessControlTranslation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.77.0/docs/resources/s3_bucket#owner S3Bucket#owner}.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
}

