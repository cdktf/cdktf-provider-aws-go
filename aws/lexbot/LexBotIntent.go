// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexbot


type LexBotIntent struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/lex_bot#intent_name LexBot#intent_name}.
	IntentName *string `field:"required" json:"intentName" yaml:"intentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/lex_bot#intent_version LexBot#intent_version}.
	IntentVersion *string `field:"required" json:"intentVersion" yaml:"intentVersion"`
}

