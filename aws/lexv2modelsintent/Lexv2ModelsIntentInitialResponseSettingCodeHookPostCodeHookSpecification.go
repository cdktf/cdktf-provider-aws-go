// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentInitialResponseSettingCodeHookPostCodeHookSpecification struct {
	// failure_conditional block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/lexv2models_intent#failure_conditional Lexv2ModelsIntent#failure_conditional}
	FailureConditional interface{} `field:"optional" json:"failureConditional" yaml:"failureConditional"`
	// failure_next_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/lexv2models_intent#failure_next_step Lexv2ModelsIntent#failure_next_step}
	FailureNextStep interface{} `field:"optional" json:"failureNextStep" yaml:"failureNextStep"`
	// failure_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/lexv2models_intent#failure_response Lexv2ModelsIntent#failure_response}
	FailureResponse interface{} `field:"optional" json:"failureResponse" yaml:"failureResponse"`
	// success_conditional block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/lexv2models_intent#success_conditional Lexv2ModelsIntent#success_conditional}
	SuccessConditional interface{} `field:"optional" json:"successConditional" yaml:"successConditional"`
	// success_next_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/lexv2models_intent#success_next_step Lexv2ModelsIntent#success_next_step}
	SuccessNextStep interface{} `field:"optional" json:"successNextStep" yaml:"successNextStep"`
	// success_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/lexv2models_intent#success_response Lexv2ModelsIntent#success_response}
	SuccessResponse interface{} `field:"optional" json:"successResponse" yaml:"successResponse"`
	// timeout_conditional block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/lexv2models_intent#timeout_conditional Lexv2ModelsIntent#timeout_conditional}
	TimeoutConditional interface{} `field:"optional" json:"timeoutConditional" yaml:"timeoutConditional"`
	// timeout_next_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/lexv2models_intent#timeout_next_step Lexv2ModelsIntent#timeout_next_step}
	TimeoutNextStep interface{} `field:"optional" json:"timeoutNextStep" yaml:"timeoutNextStep"`
	// timeout_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/lexv2models_intent#timeout_response Lexv2ModelsIntent#timeout_response}
	TimeoutResponse interface{} `field:"optional" json:"timeoutResponse" yaml:"timeoutResponse"`
}

