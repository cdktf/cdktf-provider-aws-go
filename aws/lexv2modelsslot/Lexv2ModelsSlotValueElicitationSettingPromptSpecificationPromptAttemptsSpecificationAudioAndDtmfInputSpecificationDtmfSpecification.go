// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsslot


type Lexv2ModelsSlotValueElicitationSettingPromptSpecificationPromptAttemptsSpecificationAudioAndDtmfInputSpecificationDtmfSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/lexv2models_slot#deletion_character Lexv2ModelsSlot#deletion_character}.
	DeletionCharacter *string `field:"required" json:"deletionCharacter" yaml:"deletionCharacter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/lexv2models_slot#end_character Lexv2ModelsSlot#end_character}.
	EndCharacter *string `field:"required" json:"endCharacter" yaml:"endCharacter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/lexv2models_slot#end_timeout_ms Lexv2ModelsSlot#end_timeout_ms}.
	EndTimeoutMs *float64 `field:"required" json:"endTimeoutMs" yaml:"endTimeoutMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/lexv2models_slot#max_length Lexv2ModelsSlot#max_length}.
	MaxLength *float64 `field:"required" json:"maxLength" yaml:"maxLength"`
}

