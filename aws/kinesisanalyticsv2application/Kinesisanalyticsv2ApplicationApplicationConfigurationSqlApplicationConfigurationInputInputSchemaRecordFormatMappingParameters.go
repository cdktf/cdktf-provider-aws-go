// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParameters struct {
	// csv_mapping_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/kinesisanalyticsv2_application#csv_mapping_parameters Kinesisanalyticsv2Application#csv_mapping_parameters}
	CsvMappingParameters *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersCsvMappingParameters `field:"optional" json:"csvMappingParameters" yaml:"csvMappingParameters"`
	// json_mapping_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/kinesisanalyticsv2_application#json_mapping_parameters Kinesisanalyticsv2Application#json_mapping_parameters}
	JsonMappingParameters *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersJsonMappingParameters `field:"optional" json:"jsonMappingParameters" yaml:"jsonMappingParameters"`
}

