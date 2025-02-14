// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceS3ReferenceDataSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/kinesisanalyticsv2_application#bucket_arn Kinesisanalyticsv2Application#bucket_arn}.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/kinesisanalyticsv2_application#file_key Kinesisanalyticsv2Application#file_key}.
	FileKey *string `field:"required" json:"fileKey" yaml:"fileKey"`
}

