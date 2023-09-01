// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schedulerschedule


type SchedulerScheduleTargetRetryPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/scheduler_schedule#maximum_event_age_in_seconds SchedulerSchedule#maximum_event_age_in_seconds}.
	MaximumEventAgeInSeconds *float64 `field:"optional" json:"maximumEventAgeInSeconds" yaml:"maximumEventAgeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/scheduler_schedule#maximum_retry_attempts SchedulerSchedule#maximum_retry_attempts}.
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
}

