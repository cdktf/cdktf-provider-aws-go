// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelInputSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/medialive_channel#codec MedialiveChannel#codec}.
	Codec *string `field:"required" json:"codec" yaml:"codec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/medialive_channel#input_resolution MedialiveChannel#input_resolution}.
	InputResolution *string `field:"required" json:"inputResolution" yaml:"inputResolution"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/medialive_channel#maximum_bitrate MedialiveChannel#maximum_bitrate}.
	MaximumBitrate *string `field:"required" json:"maximumBitrate" yaml:"maximumBitrate"`
}

