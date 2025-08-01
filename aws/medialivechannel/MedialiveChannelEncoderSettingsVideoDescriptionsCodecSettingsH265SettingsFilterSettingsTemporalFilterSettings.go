// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettingsTemporalFilterSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/medialive_channel#post_filter_sharpening MedialiveChannel#post_filter_sharpening}.
	PostFilterSharpening *string `field:"optional" json:"postFilterSharpening" yaml:"postFilterSharpening"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/medialive_channel#strength MedialiveChannel#strength}.
	Strength *string `field:"optional" json:"strength" yaml:"strength"`
}

