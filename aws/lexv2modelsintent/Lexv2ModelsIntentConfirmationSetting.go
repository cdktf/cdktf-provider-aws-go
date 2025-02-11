// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentConfirmationSetting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/lexv2models_intent#active Lexv2ModelsIntent#active}.
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// code_hook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/lexv2models_intent#code_hook Lexv2ModelsIntent#code_hook}
	CodeHook interface{} `field:"optional" json:"codeHook" yaml:"codeHook"`
	// confirmation_conditional block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/lexv2models_intent#confirmation_conditional Lexv2ModelsIntent#confirmation_conditional}
	ConfirmationConditional interface{} `field:"optional" json:"confirmationConditional" yaml:"confirmationConditional"`
	// confirmation_next_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/lexv2models_intent#confirmation_next_step Lexv2ModelsIntent#confirmation_next_step}
	ConfirmationNextStep interface{} `field:"optional" json:"confirmationNextStep" yaml:"confirmationNextStep"`
	// confirmation_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/lexv2models_intent#confirmation_response Lexv2ModelsIntent#confirmation_response}
	ConfirmationResponse interface{} `field:"optional" json:"confirmationResponse" yaml:"confirmationResponse"`
	// declination_conditional block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/lexv2models_intent#declination_conditional Lexv2ModelsIntent#declination_conditional}
	DeclinationConditional interface{} `field:"optional" json:"declinationConditional" yaml:"declinationConditional"`
	// declination_next_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/lexv2models_intent#declination_next_step Lexv2ModelsIntent#declination_next_step}
	DeclinationNextStep interface{} `field:"optional" json:"declinationNextStep" yaml:"declinationNextStep"`
	// declination_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/lexv2models_intent#declination_response Lexv2ModelsIntent#declination_response}
	DeclinationResponse interface{} `field:"optional" json:"declinationResponse" yaml:"declinationResponse"`
	// elicitation_code_hook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/lexv2models_intent#elicitation_code_hook Lexv2ModelsIntent#elicitation_code_hook}
	ElicitationCodeHook interface{} `field:"optional" json:"elicitationCodeHook" yaml:"elicitationCodeHook"`
	// failure_conditional block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/lexv2models_intent#failure_conditional Lexv2ModelsIntent#failure_conditional}
	FailureConditional interface{} `field:"optional" json:"failureConditional" yaml:"failureConditional"`
	// failure_next_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/lexv2models_intent#failure_next_step Lexv2ModelsIntent#failure_next_step}
	FailureNextStep interface{} `field:"optional" json:"failureNextStep" yaml:"failureNextStep"`
	// failure_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/lexv2models_intent#failure_response Lexv2ModelsIntent#failure_response}
	FailureResponse interface{} `field:"optional" json:"failureResponse" yaml:"failureResponse"`
	// prompt_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/lexv2models_intent#prompt_specification Lexv2ModelsIntent#prompt_specification}
	PromptSpecification interface{} `field:"optional" json:"promptSpecification" yaml:"promptSpecification"`
}

