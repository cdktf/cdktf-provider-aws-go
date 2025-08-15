// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devopsgurunotificationchannel


type DevopsguruNotificationChannelFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/devopsguru_notification_channel#message_types DevopsguruNotificationChannel#message_types}.
	MessageTypes *[]*string `field:"optional" json:"messageTypes" yaml:"messageTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/devopsguru_notification_channel#severities DevopsguruNotificationChannel#severities}.
	Severities *[]*string `field:"optional" json:"severities" yaml:"severities"`
}

