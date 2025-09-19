// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsapplication


type KinesisAnalyticsApplicationReferenceDataSourcesS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/kinesis_analytics_application#bucket_arn KinesisAnalyticsApplication#bucket_arn}.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/kinesis_analytics_application#file_key KinesisAnalyticsApplication#file_key}.
	FileKey *string `field:"required" json:"fileKey" yaml:"fileKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/kinesis_analytics_application#role_arn KinesisAnalyticsApplication#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

