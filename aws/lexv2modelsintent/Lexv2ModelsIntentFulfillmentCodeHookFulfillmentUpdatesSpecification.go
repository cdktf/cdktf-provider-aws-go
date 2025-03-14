// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent


type Lexv2ModelsIntentFulfillmentCodeHookFulfillmentUpdatesSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/lexv2models_intent#active Lexv2ModelsIntent#active}.
	Active interface{} `field:"required" json:"active" yaml:"active"`
	// start_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/lexv2models_intent#start_response Lexv2ModelsIntent#start_response}
	StartResponse interface{} `field:"optional" json:"startResponse" yaml:"startResponse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/lexv2models_intent#timeout_in_seconds Lexv2ModelsIntent#timeout_in_seconds}.
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
	// update_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/lexv2models_intent#update_response Lexv2ModelsIntent#update_response}
	UpdateResponse interface{} `field:"optional" json:"updateResponse" yaml:"updateResponse"`
}

