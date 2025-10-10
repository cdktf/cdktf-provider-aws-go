// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsMultiplexOutputSettings struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/medialive_channel#destination MedialiveChannel#destination}
	Destination *MedialiveChannelEncoderSettingsOutputGroupsOutputsOutputSettingsMultiplexOutputSettingsDestination `field:"required" json:"destination" yaml:"destination"`
}

