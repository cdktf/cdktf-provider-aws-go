// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package xraygroup


type XrayGroupInsightsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/xray_group#insights_enabled XrayGroup#insights_enabled}.
	InsightsEnabled interface{} `field:"required" json:"insightsEnabled" yaml:"insightsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/xray_group#notifications_enabled XrayGroup#notifications_enabled}.
	NotificationsEnabled interface{} `field:"optional" json:"notificationsEnabled" yaml:"notificationsEnabled"`
}

