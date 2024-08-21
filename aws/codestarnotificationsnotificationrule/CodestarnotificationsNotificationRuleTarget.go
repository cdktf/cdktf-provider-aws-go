// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codestarnotificationsnotificationrule


type CodestarnotificationsNotificationRuleTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/codestarnotifications_notification_rule#address CodestarnotificationsNotificationRule#address}.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/codestarnotifications_notification_rule#type CodestarnotificationsNotificationRule#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

