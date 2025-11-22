// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appautoscalingpolicy


type AppautoscalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationCustomizedCapacityMetricSpecificationMetricDataQueryMetricStatMetric struct {
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/appautoscaling_policy#dimension AppautoscalingPolicy#dimension}
	Dimension interface{} `field:"optional" json:"dimension" yaml:"dimension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/appautoscaling_policy#metric_name AppautoscalingPolicy#metric_name}.
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/appautoscaling_policy#namespace AppautoscalingPolicy#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

