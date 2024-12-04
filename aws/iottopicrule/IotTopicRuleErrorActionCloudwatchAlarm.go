// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iottopicrule


type IotTopicRuleErrorActionCloudwatchAlarm struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/iot_topic_rule#alarm_name IotTopicRule#alarm_name}.
	AlarmName *string `field:"required" json:"alarmName" yaml:"alarmName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/iot_topic_rule#role_arn IotTopicRule#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/iot_topic_rule#state_reason IotTopicRule#state_reason}.
	StateReason *string `field:"required" json:"stateReason" yaml:"stateReason"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/iot_topic_rule#state_value IotTopicRule#state_value}.
	StateValue *string `field:"required" json:"stateValue" yaml:"stateValue"`
}

