// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmcontactsrotation


type SsmcontactsRotationRecurrenceMonthlySettingsHandOffTime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/ssmcontacts_rotation#hour_of_day SsmcontactsRotation#hour_of_day}.
	HourOfDay *float64 `field:"required" json:"hourOfDay" yaml:"hourOfDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/ssmcontacts_rotation#minute_of_hour SsmcontactsRotation#minute_of_hour}.
	MinuteOfHour *float64 `field:"required" json:"minuteOfHour" yaml:"minuteOfHour"`
}

