// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devopsgurunotificationchannel


type DevopsguruNotificationChannelSns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/devopsguru_notification_channel#topic_arn DevopsguruNotificationChannel#topic_arn}.
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
}

