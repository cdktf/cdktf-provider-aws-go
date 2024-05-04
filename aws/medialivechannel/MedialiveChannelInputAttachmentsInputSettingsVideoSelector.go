// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelInputAttachmentsInputSettingsVideoSelector struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/medialive_channel#color_space MedialiveChannel#color_space}.
	ColorSpace *string `field:"optional" json:"colorSpace" yaml:"colorSpace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/medialive_channel#color_space_usage MedialiveChannel#color_space_usage}.
	ColorSpaceUsage *string `field:"optional" json:"colorSpaceUsage" yaml:"colorSpaceUsage"`
}

