// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsFecOutputSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/medialive_channel#column_depth MedialiveChannel#column_depth}.
	ColumnDepth *float64 `field:"optional" json:"columnDepth" yaml:"columnDepth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/medialive_channel#include_fec MedialiveChannel#include_fec}.
	IncludeFec *string `field:"optional" json:"includeFec" yaml:"includeFec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/medialive_channel#row_length MedialiveChannel#row_length}.
	RowLength *float64 `field:"optional" json:"rowLength" yaml:"rowLength"`
}

