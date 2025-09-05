// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package timestreamqueryscheduledquery


type TimestreamqueryScheduledQueryTargetConfigurationTimestreamConfigurationMultiMeasureMappings struct {
	// multi_measure_attribute_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/timestreamquery_scheduled_query#multi_measure_attribute_mapping TimestreamqueryScheduledQuery#multi_measure_attribute_mapping}
	MultiMeasureAttributeMapping interface{} `field:"optional" json:"multiMeasureAttributeMapping" yaml:"multiMeasureAttributeMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/timestreamquery_scheduled_query#target_multi_measure_name TimestreamqueryScheduledQuery#target_multi_measure_name}.
	TargetMultiMeasureName *string `field:"optional" json:"targetMultiMeasureName" yaml:"targetMultiMeasureName"`
}

