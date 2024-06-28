// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightrefreshschedule


type QuicksightRefreshScheduleSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.0/docs/resources/quicksight_refresh_schedule#refresh_type QuicksightRefreshSchedule#refresh_type}.
	RefreshType *string `field:"required" json:"refreshType" yaml:"refreshType"`
	// schedule_frequency block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.0/docs/resources/quicksight_refresh_schedule#schedule_frequency QuicksightRefreshSchedule#schedule_frequency}
	ScheduleFrequency interface{} `field:"optional" json:"scheduleFrequency" yaml:"scheduleFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.0/docs/resources/quicksight_refresh_schedule#start_after_date_time QuicksightRefreshSchedule#start_after_date_time}.
	StartAfterDateTime *string `field:"optional" json:"startAfterDateTime" yaml:"startAfterDateTime"`
}

