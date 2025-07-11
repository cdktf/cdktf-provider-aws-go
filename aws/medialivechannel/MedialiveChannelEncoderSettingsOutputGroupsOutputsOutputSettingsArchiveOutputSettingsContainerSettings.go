// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettings struct {
	// m2ts_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/medialive_channel#m2ts_settings MedialiveChannel#m2ts_settings}
	M2TsSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsM2TsSettings `field:"optional" json:"m2TsSettings" yaml:"m2TsSettings"`
	// raw_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/medialive_channel#raw_settings MedialiveChannel#raw_settings}
	RawSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsArchiveOutputSettingsContainerSettingsRawSettings `field:"optional" json:"rawSettings" yaml:"rawSettings"`
}

