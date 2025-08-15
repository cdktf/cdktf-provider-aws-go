// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appautoscalingpolicy


type AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStat struct {
	// metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/appautoscaling_policy#metric AppautoscalingPolicy#metric}
	Metric *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatMetric `field:"required" json:"metric" yaml:"metric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/appautoscaling_policy#stat AppautoscalingPolicy#stat}.
	Stat *string `field:"required" json:"stat" yaml:"stat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/appautoscaling_policy#unit AppautoscalingPolicy#unit}.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

