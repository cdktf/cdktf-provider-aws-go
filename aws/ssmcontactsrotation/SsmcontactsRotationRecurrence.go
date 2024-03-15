// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmcontactsrotation


type SsmcontactsRotationRecurrence struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/ssmcontacts_rotation#number_of_on_calls SsmcontactsRotation#number_of_on_calls}.
	NumberOfOnCalls *float64 `field:"required" json:"numberOfOnCalls" yaml:"numberOfOnCalls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/ssmcontacts_rotation#recurrence_multiplier SsmcontactsRotation#recurrence_multiplier}.
	RecurrenceMultiplier *float64 `field:"required" json:"recurrenceMultiplier" yaml:"recurrenceMultiplier"`
	// daily_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/ssmcontacts_rotation#daily_settings SsmcontactsRotation#daily_settings}
	DailySettings interface{} `field:"optional" json:"dailySettings" yaml:"dailySettings"`
	// monthly_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/ssmcontacts_rotation#monthly_settings SsmcontactsRotation#monthly_settings}
	MonthlySettings interface{} `field:"optional" json:"monthlySettings" yaml:"monthlySettings"`
	// shift_coverages block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/ssmcontacts_rotation#shift_coverages SsmcontactsRotation#shift_coverages}
	ShiftCoverages interface{} `field:"optional" json:"shiftCoverages" yaml:"shiftCoverages"`
	// weekly_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/ssmcontacts_rotation#weekly_settings SsmcontactsRotation#weekly_settings}
	WeeklySettings interface{} `field:"optional" json:"weeklySettings" yaml:"weeklySettings"`
}

