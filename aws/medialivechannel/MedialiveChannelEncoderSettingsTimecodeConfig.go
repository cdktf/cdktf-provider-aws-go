// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsTimecodeConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/medialive_channel#source MedialiveChannel#source}.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/medialive_channel#sync_threshold MedialiveChannel#sync_threshold}.
	SyncThreshold *float64 `field:"optional" json:"syncThreshold" yaml:"syncThreshold"`
}

