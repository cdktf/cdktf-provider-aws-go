// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasynctask


type DatasyncTaskSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/datasync_task#schedule_expression DatasyncTask#schedule_expression}.
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
}

