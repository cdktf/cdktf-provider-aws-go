// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverCondition struct {
	// failover_condition_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/medialive_channel#failover_condition_settings MedialiveChannel#failover_condition_settings}
	FailoverConditionSettings *MedialiveChannelInputAttachmentsAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettings `field:"optional" json:"failoverConditionSettings" yaml:"failoverConditionSettings"`
}

