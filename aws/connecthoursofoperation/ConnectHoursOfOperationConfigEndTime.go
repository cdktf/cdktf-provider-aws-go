// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package connecthoursofoperation


type ConnectHoursOfOperationConfigEndTime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/connect_hours_of_operation#hours ConnectHoursOfOperation#hours}.
	Hours *float64 `field:"required" json:"hours" yaml:"hours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/connect_hours_of_operation#minutes ConnectHoursOfOperation#minutes}.
	Minutes *float64 `field:"required" json:"minutes" yaml:"minutes"`
}

