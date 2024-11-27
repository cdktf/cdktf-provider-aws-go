// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentInitialResponseSetting struct {
	// code_hook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/lexv2models_intent#code_hook Lexv2ModelsIntent#code_hook}
	CodeHook interface{} `field:"optional" json:"codeHook" yaml:"codeHook"`
	// conditional block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/lexv2models_intent#conditional Lexv2ModelsIntent#conditional}
	Conditional interface{} `field:"optional" json:"conditional" yaml:"conditional"`
	// initial_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/lexv2models_intent#initial_response Lexv2ModelsIntent#initial_response}
	InitialResponse interface{} `field:"optional" json:"initialResponse" yaml:"initialResponse"`
	// next_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/lexv2models_intent#next_step Lexv2ModelsIntent#next_step}
	NextStep interface{} `field:"optional" json:"nextStep" yaml:"nextStep"`
}

