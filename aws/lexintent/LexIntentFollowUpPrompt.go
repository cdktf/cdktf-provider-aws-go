// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexintent


type LexIntentFollowUpPrompt struct {
	// prompt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/lex_intent#prompt LexIntent#prompt}
	Prompt *LexIntentFollowUpPromptPrompt `field:"required" json:"prompt" yaml:"prompt"`
	// rejection_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/lex_intent#rejection_statement LexIntent#rejection_statement}
	RejectionStatement *LexIntentFollowUpPromptRejectionStatement `field:"required" json:"rejectionStatement" yaml:"rejectionStatement"`
}

