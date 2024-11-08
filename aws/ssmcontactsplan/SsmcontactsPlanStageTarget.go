// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmcontactsplan


type SsmcontactsPlanStageTarget struct {
	// channel_target_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/ssmcontacts_plan#channel_target_info SsmcontactsPlan#channel_target_info}
	ChannelTargetInfo *SsmcontactsPlanStageTargetChannelTargetInfo `field:"optional" json:"channelTargetInfo" yaml:"channelTargetInfo"`
	// contact_target_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/ssmcontacts_plan#contact_target_info SsmcontactsPlan#contact_target_info}
	ContactTargetInfo *SsmcontactsPlanStageTargetContactTargetInfo `field:"optional" json:"contactTargetInfo" yaml:"contactTargetInfo"`
}

