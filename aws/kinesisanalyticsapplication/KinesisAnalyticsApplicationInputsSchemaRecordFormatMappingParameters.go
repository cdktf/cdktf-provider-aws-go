// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsapplication


type KinesisAnalyticsApplicationInputsSchemaRecordFormatMappingParameters struct {
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/kinesis_analytics_application#csv KinesisAnalyticsApplication#csv}
	Csv *KinesisAnalyticsApplicationInputsSchemaRecordFormatMappingParametersCsv `field:"optional" json:"csv" yaml:"csv"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/kinesis_analytics_application#json KinesisAnalyticsApplication#json}
	Json *KinesisAnalyticsApplicationInputsSchemaRecordFormatMappingParametersJson `field:"optional" json:"json" yaml:"json"`
}

