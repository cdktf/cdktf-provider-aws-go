// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketReplicationConfigurationRulesDestinationReplicationTime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/s3_bucket#minutes S3Bucket#minutes}.
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/s3_bucket#status S3Bucket#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

