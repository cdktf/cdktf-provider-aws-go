// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsapplication


type KinesisAnalyticsApplicationReferenceDataSourcesSchema struct {
	// record_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/kinesis_analytics_application#record_columns KinesisAnalyticsApplication#record_columns}
	RecordColumns interface{} `field:"required" json:"recordColumns" yaml:"recordColumns"`
	// record_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/kinesis_analytics_application#record_format KinesisAnalyticsApplication#record_format}
	RecordFormat *KinesisAnalyticsApplicationReferenceDataSourcesSchemaRecordFormat `field:"required" json:"recordFormat" yaml:"recordFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/kinesis_analytics_application#record_encoding KinesisAnalyticsApplication#record_encoding}.
	RecordEncoding *string `field:"optional" json:"recordEncoding" yaml:"recordEncoding"`
}

