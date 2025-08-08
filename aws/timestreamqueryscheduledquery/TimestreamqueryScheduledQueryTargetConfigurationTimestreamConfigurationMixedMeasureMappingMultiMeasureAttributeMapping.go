// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package timestreamqueryscheduledquery


type TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMixedMeasureMappingMultiMeasureAttributeMapping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/timestreamquery_scheduled_query#measure_value_type TimestreamqueryScheduledQuery#measure_value_type}.
	MeasureValueType *string `field:"required" json:"measureValueType" yaml:"measureValueType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/timestreamquery_scheduled_query#source_column TimestreamqueryScheduledQuery#source_column}.
	SourceColumn *string `field:"required" json:"sourceColumn" yaml:"sourceColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/timestreamquery_scheduled_query#target_multi_measure_attribute_name TimestreamqueryScheduledQuery#target_multi_measure_attribute_name}.
	TargetMultiMeasureAttributeName *string `field:"optional" json:"targetMultiMeasureAttributeName" yaml:"targetMultiMeasureAttributeName"`
}

