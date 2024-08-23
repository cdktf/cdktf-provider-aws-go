// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsslot


type Lexv2ModelsSlotValueElicitationSetting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/lexv2models_slot#slot_constraint Lexv2ModelsSlot#slot_constraint}.
	SlotConstraint *string `field:"required" json:"slotConstraint" yaml:"slotConstraint"`
	// default_value_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/lexv2models_slot#default_value_specification Lexv2ModelsSlot#default_value_specification}
	DefaultValueSpecification interface{} `field:"optional" json:"defaultValueSpecification" yaml:"defaultValueSpecification"`
	// prompt_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/lexv2models_slot#prompt_specification Lexv2ModelsSlot#prompt_specification}
	PromptSpecification interface{} `field:"optional" json:"promptSpecification" yaml:"promptSpecification"`
	// sample_utterance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/lexv2models_slot#sample_utterance Lexv2ModelsSlot#sample_utterance}
	SampleUtterance interface{} `field:"optional" json:"sampleUtterance" yaml:"sampleUtterance"`
	// slot_resolution_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/lexv2models_slot#slot_resolution_setting Lexv2ModelsSlot#slot_resolution_setting}
	SlotResolutionSetting interface{} `field:"optional" json:"slotResolutionSetting" yaml:"slotResolutionSetting"`
	// wait_and_continue_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/lexv2models_slot#wait_and_continue_specification Lexv2ModelsSlot#wait_and_continue_specification}
	WaitAndContinueSpecification interface{} `field:"optional" json:"waitAndContinueSpecification" yaml:"waitAndContinueSpecification"`
}

