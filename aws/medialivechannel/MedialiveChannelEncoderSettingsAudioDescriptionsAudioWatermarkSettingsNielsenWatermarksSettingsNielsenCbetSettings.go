// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsAudioDescriptionsAudioWatermarkSettingsNielsenWatermarksSettingsNielsenCbetSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/medialive_channel#cbet_check_digit_string MedialiveChannel#cbet_check_digit_string}.
	CbetCheckDigitString *string `field:"required" json:"cbetCheckDigitString" yaml:"cbetCheckDigitString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/medialive_channel#cbet_stepaside MedialiveChannel#cbet_stepaside}.
	CbetStepaside *string `field:"required" json:"cbetStepaside" yaml:"cbetStepaside"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/medialive_channel#csid MedialiveChannel#csid}.
	Csid *string `field:"required" json:"csid" yaml:"csid"`
}

