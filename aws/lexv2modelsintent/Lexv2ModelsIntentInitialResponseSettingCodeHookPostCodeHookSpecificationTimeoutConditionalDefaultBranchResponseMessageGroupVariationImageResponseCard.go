// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationTimeoutConditionalDefaultBranchResponseMessageGroupVariationImageResponseCard struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/lexv2models_intent#title Lexv2ModelsIntent#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// button block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/lexv2models_intent#button Lexv2ModelsIntent#button}
	Button interface{} `field:"optional" json:"button" yaml:"button"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/lexv2models_intent#image_url Lexv2ModelsIntent#image_url}.
	ImageUrl *string `field:"optional" json:"imageUrl" yaml:"imageUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/lexv2models_intent#subtitle Lexv2ModelsIntent#subtitle}.
	Subtitle *string `field:"optional" json:"subtitle" yaml:"subtitle"`
}

