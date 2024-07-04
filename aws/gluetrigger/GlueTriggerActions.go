// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluetrigger


type GlueTriggerActions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/glue_trigger#arguments GlueTrigger#arguments}.
	Arguments *map[string]*string `field:"optional" json:"arguments" yaml:"arguments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/glue_trigger#crawler_name GlueTrigger#crawler_name}.
	CrawlerName *string `field:"optional" json:"crawlerName" yaml:"crawlerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/glue_trigger#job_name GlueTrigger#job_name}.
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// notification_property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/glue_trigger#notification_property GlueTrigger#notification_property}
	NotificationProperty *GlueTriggerActionsNotificationProperty `field:"optional" json:"notificationProperty" yaml:"notificationProperty"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/glue_trigger#security_configuration GlueTrigger#security_configuration}.
	SecurityConfiguration *string `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/glue_trigger#timeout GlueTrigger#timeout}.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

