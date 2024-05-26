// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentConfirmationSettingPromptSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/lexv2models_intent#max_retries Lexv2ModelsIntent#max_retries}.
	MaxRetries *float64 `field:"required" json:"maxRetries" yaml:"maxRetries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/lexv2models_intent#allow_interrupt Lexv2ModelsIntent#allow_interrupt}.
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
	// message_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/lexv2models_intent#message_group Lexv2ModelsIntent#message_group}
	MessageGroup interface{} `field:"optional" json:"messageGroup" yaml:"messageGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/lexv2models_intent#message_selection_strategy Lexv2ModelsIntent#message_selection_strategy}.
	MessageSelectionStrategy *string `field:"optional" json:"messageSelectionStrategy" yaml:"messageSelectionStrategy"`
	// prompt_attempts_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/lexv2models_intent#prompt_attempts_specification Lexv2ModelsIntent#prompt_attempts_specification}
	PromptAttemptsSpecification interface{} `field:"optional" json:"promptAttemptsSpecification" yaml:"promptAttemptsSpecification"`
}

