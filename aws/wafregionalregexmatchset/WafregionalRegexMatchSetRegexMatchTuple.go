// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafregionalregexmatchset


type WafregionalRegexMatchSetRegexMatchTuple struct {
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/wafregional_regex_match_set#field_to_match WafregionalRegexMatchSet#field_to_match}
	FieldToMatch *WafregionalRegexMatchSetRegexMatchTupleFieldToMatch `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/wafregional_regex_match_set#regex_pattern_set_id WafregionalRegexMatchSet#regex_pattern_set_id}.
	RegexPatternSetId *string `field:"required" json:"regexPatternSetId" yaml:"regexPatternSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/wafregional_regex_match_set#text_transformation WafregionalRegexMatchSet#text_transformation}.
	TextTransformation *string `field:"required" json:"textTransformation" yaml:"textTransformation"`
}

