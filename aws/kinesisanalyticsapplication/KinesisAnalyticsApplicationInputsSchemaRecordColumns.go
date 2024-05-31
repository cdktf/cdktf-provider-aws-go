// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsapplication


type KinesisAnalyticsApplicationInputsSchemaRecordColumns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/kinesis_analytics_application#name KinesisAnalyticsApplication#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/kinesis_analytics_application#sql_type KinesisAnalyticsApplication#sql_type}.
	SqlType *string `field:"required" json:"sqlType" yaml:"sqlType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/kinesis_analytics_application#mapping KinesisAnalyticsApplication#mapping}.
	Mapping *string `field:"optional" json:"mapping" yaml:"mapping"`
}

