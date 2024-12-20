// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentConfirmationSettingPromptSpecificationPromptAttemptsSpecificationAudioAndDtmfInputSpecificationDtmfSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/lexv2models_intent#deletion_character Lexv2ModelsIntent#deletion_character}.
	DeletionCharacter *string `field:"required" json:"deletionCharacter" yaml:"deletionCharacter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/lexv2models_intent#end_character Lexv2ModelsIntent#end_character}.
	EndCharacter *string `field:"required" json:"endCharacter" yaml:"endCharacter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/lexv2models_intent#end_timeout_ms Lexv2ModelsIntent#end_timeout_ms}.
	EndTimeoutMs *float64 `field:"required" json:"endTimeoutMs" yaml:"endTimeoutMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/lexv2models_intent#max_length Lexv2ModelsIntent#max_length}.
	MaxLength *float64 `field:"required" json:"maxLength" yaml:"maxLength"`
}

