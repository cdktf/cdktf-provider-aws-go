// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsTeletextSourceSettings struct {
	// output_rectangle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/medialive_channel#output_rectangle MedialiveChannel#output_rectangle}
	OutputRectangle *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsTeletextSourceSettingsOutputRectangle `field:"optional" json:"outputRectangle" yaml:"outputRectangle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/medialive_channel#page_number MedialiveChannel#page_number}.
	PageNumber *string `field:"optional" json:"pageNumber" yaml:"pageNumber"`
}

