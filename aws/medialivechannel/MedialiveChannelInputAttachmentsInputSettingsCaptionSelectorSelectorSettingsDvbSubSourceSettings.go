// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettingsDvbSubSourceSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.54.1/docs/resources/medialive_channel#ocr_language MedialiveChannel#ocr_language}.
	OcrLanguage *string `field:"optional" json:"ocrLanguage" yaml:"ocrLanguage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.54.1/docs/resources/medialive_channel#pid MedialiveChannel#pid}.
	Pid *float64 `field:"optional" json:"pid" yaml:"pid"`
}

