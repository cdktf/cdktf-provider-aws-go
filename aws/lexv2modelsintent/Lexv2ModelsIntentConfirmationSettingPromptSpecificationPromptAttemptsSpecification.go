// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/lexv2models_intent#map_block_key Lexv2ModelsIntent#map_block_key}.
	MapBlockKey *string `field:"required" json:"mapBlockKey" yaml:"mapBlockKey"`
	// allowed_input_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/lexv2models_intent#allowed_input_types Lexv2ModelsIntent#allowed_input_types}
	AllowedInputTypes interface{} `field:"optional" json:"allowedInputTypes" yaml:"allowedInputTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/lexv2models_intent#allow_interrupt Lexv2ModelsIntent#allow_interrupt}.
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
	// audio_and_dtmf_input_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/lexv2models_intent#audio_and_dtmf_input_specification Lexv2ModelsIntent#audio_and_dtmf_input_specification}
	AudioAndDtmfInputSpecification interface{} `field:"optional" json:"audioAndDtmfInputSpecification" yaml:"audioAndDtmfInputSpecification"`
	// text_input_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/lexv2models_intent#text_input_specification Lexv2ModelsIntent#text_input_specification}
	TextInputSpecification interface{} `field:"optional" json:"textInputSpecification" yaml:"textInputSpecification"`
}

