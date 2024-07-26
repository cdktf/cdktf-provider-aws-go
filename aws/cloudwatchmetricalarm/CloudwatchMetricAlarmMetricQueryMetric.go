// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatchmetricalarm


type CloudwatchMetricAlarmMetricQueryMetric struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/cloudwatch_metric_alarm#metric_name CloudwatchMetricAlarm#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/cloudwatch_metric_alarm#period CloudwatchMetricAlarm#period}.
	Period *float64 `field:"required" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/cloudwatch_metric_alarm#stat CloudwatchMetricAlarm#stat}.
	Stat *string `field:"required" json:"stat" yaml:"stat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/cloudwatch_metric_alarm#dimensions CloudwatchMetricAlarm#dimensions}.
	Dimensions *map[string]*string `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/cloudwatch_metric_alarm#namespace CloudwatchMetricAlarm#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/cloudwatch_metric_alarm#unit CloudwatchMetricAlarm#unit}.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

