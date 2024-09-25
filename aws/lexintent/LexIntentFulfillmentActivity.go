// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexintent


type LexIntentFulfillmentActivity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/lex_intent#type LexIntent#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// code_hook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/lex_intent#code_hook LexIntent#code_hook}
	CodeHook *LexIntentFulfillmentActivityCodeHook `field:"optional" json:"codeHook" yaml:"codeHook"`
}

