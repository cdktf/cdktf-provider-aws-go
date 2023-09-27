// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormat struct {
	// mapping_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/kinesisanalyticsv2_application#mapping_parameters Kinesisanalyticsv2Application#mapping_parameters}
	MappingParameters *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParameters `field:"required" json:"mappingParameters" yaml:"mappingParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/kinesisanalyticsv2_application#record_format_type Kinesisanalyticsv2Application#record_format_type}.
	RecordFormatType *string `field:"required" json:"recordFormatType" yaml:"recordFormatType"`
}

