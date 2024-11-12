// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmmaintenancewindowtask


type SsmMaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/ssm_maintenance_window_task#notification_arn SsmMaintenanceWindowTask#notification_arn}.
	NotificationArn *string `field:"optional" json:"notificationArn" yaml:"notificationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/ssm_maintenance_window_task#notification_events SsmMaintenanceWindowTask#notification_events}.
	NotificationEvents *[]*string `field:"optional" json:"notificationEvents" yaml:"notificationEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/ssm_maintenance_window_task#notification_type SsmMaintenanceWindowTask#notification_type}.
	NotificationType *string `field:"optional" json:"notificationType" yaml:"notificationType"`
}

