// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iottopicrule


type IotTopicRuleCloudwatchMetric struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/iot_topic_rule#metric_name IotTopicRule#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/iot_topic_rule#metric_namespace IotTopicRule#metric_namespace}.
	MetricNamespace *string `field:"required" json:"metricNamespace" yaml:"metricNamespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/iot_topic_rule#metric_unit IotTopicRule#metric_unit}.
	MetricUnit *string `field:"required" json:"metricUnit" yaml:"metricUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/iot_topic_rule#metric_value IotTopicRule#metric_value}.
	MetricValue *string `field:"required" json:"metricValue" yaml:"metricValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/iot_topic_rule#role_arn IotTopicRule#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/iot_topic_rule#metric_timestamp IotTopicRule#metric_timestamp}.
	MetricTimestamp *string `field:"optional" json:"metricTimestamp" yaml:"metricTimestamp"`
}

