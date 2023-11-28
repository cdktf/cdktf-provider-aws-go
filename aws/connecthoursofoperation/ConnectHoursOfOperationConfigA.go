// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package connecthoursofoperation


type ConnectHoursOfOperationConfigA struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/connect_hours_of_operation#day ConnectHoursOfOperation#day}.
	Day *string `field:"required" json:"day" yaml:"day"`
	// end_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/connect_hours_of_operation#end_time ConnectHoursOfOperation#end_time}
	EndTime *ConnectHoursOfOperationConfigEndTime `field:"required" json:"endTime" yaml:"endTime"`
	// start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/connect_hours_of_operation#start_time ConnectHoursOfOperation#start_time}
	StartTime *ConnectHoursOfOperationConfigStartTime `field:"required" json:"startTime" yaml:"startTime"`
}

