// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelInputAttachmentsAutomaticInputFailoverSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/medialive_channel#secondary_input_id MedialiveChannel#secondary_input_id}.
	SecondaryInputId *string `field:"required" json:"secondaryInputId" yaml:"secondaryInputId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/medialive_channel#error_clear_time_msec MedialiveChannel#error_clear_time_msec}.
	ErrorClearTimeMsec *float64 `field:"optional" json:"errorClearTimeMsec" yaml:"errorClearTimeMsec"`
	// failover_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/medialive_channel#failover_condition MedialiveChannel#failover_condition}
	FailoverCondition interface{} `field:"optional" json:"failoverCondition" yaml:"failoverCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/medialive_channel#input_preference MedialiveChannel#input_preference}.
	InputPreference *string `field:"optional" json:"inputPreference" yaml:"inputPreference"`
}

