// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsKeyProviderSettingsStaticKeySettingsKeyProviderServer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/medialive_channel#uri MedialiveChannel#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/medialive_channel#password_param MedialiveChannel#password_param}.
	PasswordParam *string `field:"optional" json:"passwordParam" yaml:"passwordParam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/medialive_channel#username MedialiveChannel#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

