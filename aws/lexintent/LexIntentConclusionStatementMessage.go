// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexintent


type LexIntentConclusionStatementMessage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/lex_intent#content LexIntent#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/lex_intent#content_type LexIntent#content_type}.
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/lex_intent#group_number LexIntent#group_number}.
	GroupNumber *float64 `field:"optional" json:"groupNumber" yaml:"groupNumber"`
}

