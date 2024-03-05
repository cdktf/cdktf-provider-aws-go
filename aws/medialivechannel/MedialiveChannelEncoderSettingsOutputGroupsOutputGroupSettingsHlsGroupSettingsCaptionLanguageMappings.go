// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsCaptionLanguageMappings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/medialive_channel#caption_channel MedialiveChannel#caption_channel}.
	CaptionChannel *float64 `field:"required" json:"captionChannel" yaml:"captionChannel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/medialive_channel#language_code MedialiveChannel#language_code}.
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/medialive_channel#language_description MedialiveChannel#language_description}.
	LanguageDescription *string `field:"required" json:"languageDescription" yaml:"languageDescription"`
}

