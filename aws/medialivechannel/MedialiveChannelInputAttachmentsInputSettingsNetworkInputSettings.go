// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelInputAttachmentsInputSettingsNetworkInputSettings struct {
	// hls_input_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/medialive_channel#hls_input_settings MedialiveChannel#hls_input_settings}
	HlsInputSettings *MedialiveChannelInputAttachmentsInputSettingsNetworkInputSettingsHlsInputSettings `field:"optional" json:"hlsInputSettings" yaml:"hlsInputSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.25.0/docs/resources/medialive_channel#server_validation MedialiveChannel#server_validation}.
	ServerValidation *string `field:"optional" json:"serverValidation" yaml:"serverValidation"`
}

