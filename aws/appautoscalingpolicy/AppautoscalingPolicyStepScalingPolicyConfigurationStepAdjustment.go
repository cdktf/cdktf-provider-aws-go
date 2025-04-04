// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appautoscalingpolicy


type AppautoscalingPolicyStepScalingPolicyConfigurationStepAdjustment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/appautoscaling_policy#scaling_adjustment AppautoscalingPolicy#scaling_adjustment}.
	ScalingAdjustment *float64 `field:"required" json:"scalingAdjustment" yaml:"scalingAdjustment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/appautoscaling_policy#metric_interval_lower_bound AppautoscalingPolicy#metric_interval_lower_bound}.
	MetricIntervalLowerBound *string `field:"optional" json:"metricIntervalLowerBound" yaml:"metricIntervalLowerBound"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/appautoscaling_policy#metric_interval_upper_bound AppautoscalingPolicy#metric_interval_upper_bound}.
	MetricIntervalUpperBound *string `field:"optional" json:"metricIntervalUpperBound" yaml:"metricIntervalUpperBound"`
}

