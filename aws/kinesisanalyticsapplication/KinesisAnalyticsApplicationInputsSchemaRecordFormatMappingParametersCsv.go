// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsapplication


type KinesisAnalyticsApplicationInputsSchemaRecordFormatMappingParametersCsv struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kinesis_analytics_application#record_column_delimiter KinesisAnalyticsApplication#record_column_delimiter}.
	RecordColumnDelimiter *string `field:"required" json:"recordColumnDelimiter" yaml:"recordColumnDelimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/kinesis_analytics_application#record_row_delimiter KinesisAnalyticsApplication#record_row_delimiter}.
	RecordRowDelimiter *string `field:"required" json:"recordRowDelimiter" yaml:"recordRowDelimiter"`
}

