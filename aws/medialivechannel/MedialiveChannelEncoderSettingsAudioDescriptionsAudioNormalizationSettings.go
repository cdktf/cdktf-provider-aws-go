// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsAudioDescriptionsAudioNormalizationSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/medialive_channel#algorithm MedialiveChannel#algorithm}.
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/medialive_channel#algorithm_control MedialiveChannel#algorithm_control}.
	AlgorithmControl *string `field:"optional" json:"algorithmControl" yaml:"algorithmControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/medialive_channel#target_lkfs MedialiveChannel#target_lkfs}.
	TargetLkfs *float64 `field:"optional" json:"targetLkfs" yaml:"targetLkfs"`
}

