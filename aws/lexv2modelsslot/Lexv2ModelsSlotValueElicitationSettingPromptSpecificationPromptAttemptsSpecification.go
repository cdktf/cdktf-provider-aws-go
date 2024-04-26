// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsslot


type Lexv2ModelsSlotValueElicitationSettingPromptSpecificationPromptAttemptsSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/lexv2models_slot#map_block_key Lexv2ModelsSlot#map_block_key}.
	MapBlockKey *string `field:"required" json:"mapBlockKey" yaml:"mapBlockKey"`
	// allowed_input_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/lexv2models_slot#allowed_input_types Lexv2ModelsSlot#allowed_input_types}
	AllowedInputTypes interface{} `field:"optional" json:"allowedInputTypes" yaml:"allowedInputTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/lexv2models_slot#allow_interrupt Lexv2ModelsSlot#allow_interrupt}.
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
	// audio_and_dtmf_input_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/lexv2models_slot#audio_and_dtmf_input_specification Lexv2ModelsSlot#audio_and_dtmf_input_specification}
	AudioAndDtmfInputSpecification interface{} `field:"optional" json:"audioAndDtmfInputSpecification" yaml:"audioAndDtmfInputSpecification"`
	// text_input_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/lexv2models_slot#text_input_specification Lexv2ModelsSlot#text_input_specification}
	TextInputSpecification interface{} `field:"optional" json:"textInputSpecification" yaml:"textInputSpecification"`
}

