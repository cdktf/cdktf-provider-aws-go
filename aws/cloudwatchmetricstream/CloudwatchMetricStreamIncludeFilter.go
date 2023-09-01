// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatchmetricstream


type CloudwatchMetricStreamIncludeFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/cloudwatch_metric_stream#namespace CloudwatchMetricStream#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/cloudwatch_metric_stream#metric_names CloudwatchMetricStream#metric_names}.
	MetricNames *[]*string `field:"optional" json:"metricNames" yaml:"metricNames"`
}

