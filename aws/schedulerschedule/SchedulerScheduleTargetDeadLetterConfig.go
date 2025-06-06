// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schedulerschedule


type SchedulerScheduleTargetDeadLetterConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/scheduler_schedule#arn SchedulerSchedule#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

