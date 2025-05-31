// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatchmetricalarm


type CloudwatchMetricAlarmMetricQuery struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cloudwatch_metric_alarm#id CloudwatchMetricAlarm#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cloudwatch_metric_alarm#account_id CloudwatchMetricAlarm#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cloudwatch_metric_alarm#expression CloudwatchMetricAlarm#expression}.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cloudwatch_metric_alarm#label CloudwatchMetricAlarm#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cloudwatch_metric_alarm#metric CloudwatchMetricAlarm#metric}
	Metric *CloudwatchMetricAlarmMetricQueryMetric `field:"optional" json:"metric" yaml:"metric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cloudwatch_metric_alarm#period CloudwatchMetricAlarm#period}.
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cloudwatch_metric_alarm#return_data CloudwatchMetricAlarm#return_data}.
	ReturnData interface{} `field:"optional" json:"returnData" yaml:"returnData"`
}

