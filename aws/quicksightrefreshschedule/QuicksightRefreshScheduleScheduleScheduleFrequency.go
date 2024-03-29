// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightrefreshschedule


type QuicksightRefreshScheduleScheduleScheduleFrequency struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/quicksight_refresh_schedule#interval QuicksightRefreshSchedule#interval}.
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// refresh_on_day block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/quicksight_refresh_schedule#refresh_on_day QuicksightRefreshSchedule#refresh_on_day}
	RefreshOnDay interface{} `field:"optional" json:"refreshOnDay" yaml:"refreshOnDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/quicksight_refresh_schedule#time_of_the_day QuicksightRefreshSchedule#time_of_the_day}.
	TimeOfTheDay *string `field:"optional" json:"timeOfTheDay" yaml:"timeOfTheDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/quicksight_refresh_schedule#timezone QuicksightRefreshSchedule#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

