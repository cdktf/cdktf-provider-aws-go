// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ivsplaybackkeypair


type IvsPlaybackKeyPairTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ivs_playback_key_pair#create IvsPlaybackKeyPair#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ivs_playback_key_pair#delete IvsPlaybackKeyPair#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

