// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package timestreamqueryscheduledquery


type TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/timestreamquery_scheduled_query#database_name TimestreamqueryScheduledQuery#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/timestreamquery_scheduled_query#table_name TimestreamqueryScheduledQuery#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/timestreamquery_scheduled_query#time_column TimestreamqueryScheduledQuery#time_column}.
	TimeColumn *string `field:"required" json:"timeColumn" yaml:"timeColumn"`
	// dimension_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/timestreamquery_scheduled_query#dimension_mapping TimestreamqueryScheduledQuery#dimension_mapping}
	DimensionMapping interface{} `field:"optional" json:"dimensionMapping" yaml:"dimensionMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/timestreamquery_scheduled_query#measure_name_column TimestreamqueryScheduledQuery#measure_name_column}.
	MeasureNameColumn *string `field:"optional" json:"measureNameColumn" yaml:"measureNameColumn"`
	// mixed_measure_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/timestreamquery_scheduled_query#mixed_measure_mapping TimestreamqueryScheduledQuery#mixed_measure_mapping}
	MixedMeasureMapping interface{} `field:"optional" json:"mixedMeasureMapping" yaml:"mixedMeasureMapping"`
	// multi_measure_mappings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/timestreamquery_scheduled_query#multi_measure_mappings TimestreamqueryScheduledQuery#multi_measure_mappings}
	MultiMeasureMappings interface{} `field:"optional" json:"multiMeasureMappings" yaml:"multiMeasureMappings"`
}

