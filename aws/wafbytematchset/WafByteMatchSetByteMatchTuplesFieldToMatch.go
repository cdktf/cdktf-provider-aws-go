// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafbytematchset


type WafByteMatchSetByteMatchTuplesFieldToMatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/waf_byte_match_set#type WafByteMatchSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/waf_byte_match_set#data WafByteMatchSet#data}.
	Data *string `field:"optional" json:"data" yaml:"data"`
}

