// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettings struct {
	// container_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/medialive_channel#container_settings MedialiveChannel#container_settings}
	ContainerSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettings `field:"optional" json:"containerSettings" yaml:"containerSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/medialive_channel#extension MedialiveChannel#extension}.
	Extension *string `field:"optional" json:"extension" yaml:"extension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/medialive_channel#name_modifier MedialiveChannel#name_modifier}.
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
}

