// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsNielsenConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/medialive_channel#distributor_id MedialiveChannel#distributor_id}.
	DistributorId *string `field:"optional" json:"distributorId" yaml:"distributorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/medialive_channel#nielsen_pcm_to_id3_tagging MedialiveChannel#nielsen_pcm_to_id3_tagging}.
	NielsenPcmToId3Tagging *string `field:"optional" json:"nielsenPcmToId3Tagging" yaml:"nielsenPcmToId3Tagging"`
}

