// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafregexmatchset


type WafRegexMatchSetRegexMatchTuple struct {
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.19.0/docs/resources/waf_regex_match_set#field_to_match WafRegexMatchSet#field_to_match}
	FieldToMatch *WafRegexMatchSetRegexMatchTupleFieldToMatch `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.19.0/docs/resources/waf_regex_match_set#regex_pattern_set_id WafRegexMatchSet#regex_pattern_set_id}.
	RegexPatternSetId *string `field:"required" json:"regexPatternSetId" yaml:"regexPatternSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.19.0/docs/resources/waf_regex_match_set#text_transformation WafRegexMatchSet#text_transformation}.
	TextTransformation *string `field:"required" json:"textTransformation" yaml:"textTransformation"`
}

