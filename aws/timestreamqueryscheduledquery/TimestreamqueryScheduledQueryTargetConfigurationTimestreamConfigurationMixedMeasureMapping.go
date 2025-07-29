// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package timestreamqueryscheduledquery


type TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMapping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/timestreamquery_scheduled_query#measure_value_type TimestreamqueryScheduledQuery#measure_value_type}.
	MeasureValueType *string `field:"required" json:"measureValueType" yaml:"measureValueType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/timestreamquery_scheduled_query#measure_name TimestreamqueryScheduledQuery#measure_name}.
	MeasureName *string `field:"optional" json:"measureName" yaml:"measureName"`
	// multi_measure_attribute_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/timestreamquery_scheduled_query#multi_measure_attribute_mapping TimestreamqueryScheduledQuery#multi_measure_attribute_mapping}
	MultiMeasureAttributeMapping interface{} `field:"optional" json:"multiMeasureAttributeMapping" yaml:"multiMeasureAttributeMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/timestreamquery_scheduled_query#source_column TimestreamqueryScheduledQuery#source_column}.
	SourceColumn *string `field:"optional" json:"sourceColumn" yaml:"sourceColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/timestreamquery_scheduled_query#target_measure_name TimestreamqueryScheduledQuery#target_measure_name}.
	TargetMeasureName *string `field:"optional" json:"targetMeasureName" yaml:"targetMeasureName"`
}

