// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/lexv2models_intent#name Lexv2ModelsIntent#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/lexv2models_intent#condition Lexv2ModelsIntent#condition}
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// next_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/lexv2models_intent#next_step Lexv2ModelsIntent#next_step}
	NextStep interface{} `field:"optional" json:"nextStep" yaml:"nextStep"`
	// response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/lexv2models_intent#response Lexv2ModelsIntent#response}
	Response interface{} `field:"optional" json:"response" yaml:"response"`
}

