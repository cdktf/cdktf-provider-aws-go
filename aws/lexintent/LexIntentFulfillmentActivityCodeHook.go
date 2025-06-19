// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexintent


type LexIntentFulfillmentActivityCodeHook struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/lex_intent#message_version LexIntent#message_version}.
	MessageVersion *string `field:"required" json:"messageVersion" yaml:"messageVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/resources/lex_intent#uri LexIntent#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

