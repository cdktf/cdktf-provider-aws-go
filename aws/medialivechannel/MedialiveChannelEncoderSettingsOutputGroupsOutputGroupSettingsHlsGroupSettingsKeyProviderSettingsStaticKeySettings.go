// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsKeyProviderSettingsStaticKeySettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.54.1/docs/resources/medialive_channel#static_key_value MedialiveChannel#static_key_value}.
	StaticKeyValue *string `field:"required" json:"staticKeyValue" yaml:"staticKeyValue"`
	// key_provider_server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.54.1/docs/resources/medialive_channel#key_provider_server MedialiveChannel#key_provider_server}
	KeyProviderServer *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsKeyProviderSettingsStaticKeySettingsKeyProviderServer `field:"optional" json:"keyProviderServer" yaml:"keyProviderServer"`
}

