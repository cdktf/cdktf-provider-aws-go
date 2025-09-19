// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmcontactsplan


type SsmcontactsPlanStageTargetChannelTargetInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/ssmcontacts_plan#contact_channel_id SsmcontactsPlan#contact_channel_id}.
	ContactChannelId *string `field:"required" json:"contactChannelId" yaml:"contactChannelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/ssmcontacts_plan#retry_interval_in_minutes SsmcontactsPlan#retry_interval_in_minutes}.
	RetryIntervalInMinutes *float64 `field:"optional" json:"retryIntervalInMinutes" yaml:"retryIntervalInMinutes"`
}

