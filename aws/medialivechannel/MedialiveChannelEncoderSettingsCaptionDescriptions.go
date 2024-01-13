// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsCaptionDescriptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/medialive_channel#caption_selector_name MedialiveChannel#caption_selector_name}.
	CaptionSelectorName *string `field:"required" json:"captionSelectorName" yaml:"captionSelectorName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/medialive_channel#name MedialiveChannel#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/medialive_channel#accessibility MedialiveChannel#accessibility}.
	Accessibility *string `field:"optional" json:"accessibility" yaml:"accessibility"`
	// destination_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/medialive_channel#destination_settings MedialiveChannel#destination_settings}
	DestinationSettings *MedialiveChannelEncoderSettingsCaptionDescriptionsDestinationSettings `field:"optional" json:"destinationSettings" yaml:"destinationSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/medialive_channel#language_code MedialiveChannel#language_code}.
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/medialive_channel#language_description MedialiveChannel#language_description}.
	LanguageDescription *string `field:"optional" json:"languageDescription" yaml:"languageDescription"`
}

