// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsslottype


type Lexv2ModelsSlotTypeValueSelectionSetting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/lexv2models_slot_type#resolution_strategy Lexv2ModelsSlotType#resolution_strategy}.
	ResolutionStrategy *string `field:"required" json:"resolutionStrategy" yaml:"resolutionStrategy"`
	// advanced_recognition_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/lexv2models_slot_type#advanced_recognition_setting Lexv2ModelsSlotType#advanced_recognition_setting}
	AdvancedRecognitionSetting interface{} `field:"optional" json:"advancedRecognitionSetting" yaml:"advancedRecognitionSetting"`
	// regex_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/lexv2models_slot_type#regex_filter Lexv2ModelsSlotType#regex_filter}
	RegexFilter interface{} `field:"optional" json:"regexFilter" yaml:"regexFilter"`
}

