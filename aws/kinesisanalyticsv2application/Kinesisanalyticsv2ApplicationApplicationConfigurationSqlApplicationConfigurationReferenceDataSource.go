// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSource struct {
	// reference_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/kinesisanalyticsv2_application#reference_schema Kinesisanalyticsv2Application#reference_schema}
	ReferenceSchema *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchema `field:"required" json:"referenceSchema" yaml:"referenceSchema"`
	// s3_reference_data_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/kinesisanalyticsv2_application#s3_reference_data_source Kinesisanalyticsv2Application#s3_reference_data_source}
	S3ReferenceDataSource *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceS3ReferenceDataSource `field:"required" json:"s3ReferenceDataSource" yaml:"s3ReferenceDataSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/kinesisanalyticsv2_application#table_name Kinesisanalyticsv2Application#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

