// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appautoscalingpolicy


type AppautoscalingPolicyPredictiveScalingPolicyConfiguration struct {
	// metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/appautoscaling_policy#metric_specification AppautoscalingPolicy#metric_specification}
	MetricSpecification interface{} `field:"required" json:"metricSpecification" yaml:"metricSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/appautoscaling_policy#max_capacity_breach_behavior AppautoscalingPolicy#max_capacity_breach_behavior}.
	MaxCapacityBreachBehavior *string `field:"optional" json:"maxCapacityBreachBehavior" yaml:"maxCapacityBreachBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/appautoscaling_policy#max_capacity_buffer AppautoscalingPolicy#max_capacity_buffer}.
	MaxCapacityBuffer *float64 `field:"optional" json:"maxCapacityBuffer" yaml:"maxCapacityBuffer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/appautoscaling_policy#mode AppautoscalingPolicy#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/appautoscaling_policy#scheduling_buffer_time AppautoscalingPolicy#scheduling_buffer_time}.
	SchedulingBufferTime *float64 `field:"optional" json:"schedulingBufferTime" yaml:"schedulingBufferTime"`
}

