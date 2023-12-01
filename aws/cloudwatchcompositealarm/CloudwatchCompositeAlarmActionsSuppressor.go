// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatchcompositealarm


type CloudwatchCompositeAlarmActionsSuppressor struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/cloudwatch_composite_alarm#alarm CloudwatchCompositeAlarm#alarm}.
	Alarm *string `field:"required" json:"alarm" yaml:"alarm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/cloudwatch_composite_alarm#extension_period CloudwatchCompositeAlarm#extension_period}.
	ExtensionPeriod *float64 `field:"required" json:"extensionPeriod" yaml:"extensionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/cloudwatch_composite_alarm#wait_period CloudwatchCompositeAlarm#wait_period}.
	WaitPeriod *float64 `field:"required" json:"waitPeriod" yaml:"waitPeriod"`
}

