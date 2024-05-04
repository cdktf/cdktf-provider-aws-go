// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalingpolicy


type AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStat struct {
	// metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/autoscaling_policy#metric AutoscalingPolicy#metric}
	Metric *AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetric `field:"required" json:"metric" yaml:"metric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/autoscaling_policy#stat AutoscalingPolicy#stat}.
	Stat *string `field:"required" json:"stat" yaml:"stat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/autoscaling_policy#unit AutoscalingPolicy#unit}.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

