// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appautoscalingpolicy


type AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedLoadMetricSpecification struct {
	// metric_data_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/appautoscaling_policy#metric_data_query AppautoscalingPolicy#metric_data_query}
	MetricDataQuery interface{} `field:"required" json:"metricDataQuery" yaml:"metricDataQuery"`
}

