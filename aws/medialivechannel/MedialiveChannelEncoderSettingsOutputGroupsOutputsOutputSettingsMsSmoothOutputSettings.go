// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsMsSmoothOutputSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/medialive_channel#h265_packaging_type MedialiveChannel#h265_packaging_type}.
	H265PackagingType *string `field:"optional" json:"h265PackagingType" yaml:"h265PackagingType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/medialive_channel#name_modifier MedialiveChannel#name_modifier}.
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
}

