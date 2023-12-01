// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotClarificationPrompt struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/lex_bot#max_attempts LexBot#max_attempts}.
	MaxAttempts *float64 `field:"required" json:"maxAttempts" yaml:"maxAttempts"`
	// message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/lex_bot#message LexBot#message}
	Message interface{} `field:"required" json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/lex_bot#response_card LexBot#response_card}.
	ResponseCard *string `field:"optional" json:"responseCard" yaml:"responseCard"`
}

