// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalingpolicy


type AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetrics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/autoscaling_policy#id AutoscalingPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/autoscaling_policy#expression AutoscalingPolicy#expression}.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/autoscaling_policy#label AutoscalingPolicy#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// metric_stat block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/autoscaling_policy#metric_stat AutoscalingPolicy#metric_stat}
	MetricStat *AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStat `field:"optional" json:"metricStat" yaml:"metricStat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/autoscaling_policy#return_data AutoscalingPolicy#return_data}.
	ReturnData interface{} `field:"optional" json:"returnData" yaml:"returnData"`
}

