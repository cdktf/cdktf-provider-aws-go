// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureConditionalDefaultBranch struct {
	// next_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/lexv2models_intent#next_step Lexv2ModelsIntent#next_step}
	NextStep interface{} `field:"optional" json:"nextStep" yaml:"nextStep"`
	// response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/lexv2models_intent#response Lexv2ModelsIntent#response}
	Response interface{} `field:"optional" json:"response" yaml:"response"`
}

