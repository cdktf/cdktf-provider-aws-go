// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentClosingSetting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/lexv2models_intent#active Lexv2ModelsIntent#active}.
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// closing_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/lexv2models_intent#closing_response Lexv2ModelsIntent#closing_response}
	ClosingResponse interface{} `field:"optional" json:"closingResponse" yaml:"closingResponse"`
	// conditional block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/lexv2models_intent#conditional Lexv2ModelsIntent#conditional}
	Conditional interface{} `field:"optional" json:"conditional" yaml:"conditional"`
	// next_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/lexv2models_intent#next_step Lexv2ModelsIntent#next_step}
	NextStep interface{} `field:"optional" json:"nextStep" yaml:"nextStep"`
}

