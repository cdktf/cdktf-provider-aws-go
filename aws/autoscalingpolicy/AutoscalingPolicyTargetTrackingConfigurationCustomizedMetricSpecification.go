// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalingpolicy


type AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecification struct {
	// metric_dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/autoscaling_policy#metric_dimension AutoscalingPolicy#metric_dimension}
	MetricDimension interface{} `field:"optional" json:"metricDimension" yaml:"metricDimension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/autoscaling_policy#metric_name AutoscalingPolicy#metric_name}.
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/autoscaling_policy#metrics AutoscalingPolicy#metrics}
	Metrics interface{} `field:"optional" json:"metrics" yaml:"metrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/autoscaling_policy#namespace AutoscalingPolicy#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/autoscaling_policy#statistic AutoscalingPolicy#statistic}.
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/autoscaling_policy#unit AutoscalingPolicy#unit}.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

