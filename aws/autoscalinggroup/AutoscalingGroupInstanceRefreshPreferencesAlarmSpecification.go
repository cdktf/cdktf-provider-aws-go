// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup


type AutoscalingGroupInstanceRefreshPreferencesAlarmSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.61.0/docs/resources/autoscaling_group#alarms AutoscalingGroup#alarms}.
	Alarms *[]*string `field:"optional" json:"alarms" yaml:"alarms"`
}

