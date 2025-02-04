// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafregexmatchset


type WafRegexMatchSetRegexMatchTupleFieldToMatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/waf_regex_match_set#type WafRegexMatchSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/waf_regex_match_set#data WafRegexMatchSet#data}.
	Data *string `field:"optional" json:"data" yaml:"data"`
}

