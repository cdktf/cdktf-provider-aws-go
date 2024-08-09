// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsapplication


type KinesisAnalyticsApplicationInputsSchemaRecordFormat struct {
	// mapping_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/kinesis_analytics_application#mapping_parameters KinesisAnalyticsApplication#mapping_parameters}
	MappingParameters *KinesisAnalyticsApplicationInputsSchemaRecordFormatMappingParameters `field:"optional" json:"mappingParameters" yaml:"mappingParameters"`
}

