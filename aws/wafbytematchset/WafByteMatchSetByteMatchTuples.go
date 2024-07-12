// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafbytematchset


type WafByteMatchSetByteMatchTuples struct {
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/waf_byte_match_set#field_to_match WafByteMatchSet#field_to_match}
	FieldToMatch *WafByteMatchSetByteMatchTuplesFieldToMatch `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/waf_byte_match_set#positional_constraint WafByteMatchSet#positional_constraint}.
	PositionalConstraint *string `field:"required" json:"positionalConstraint" yaml:"positionalConstraint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/waf_byte_match_set#text_transformation WafByteMatchSet#text_transformation}.
	TextTransformation *string `field:"required" json:"textTransformation" yaml:"textTransformation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/waf_byte_match_set#target_string WafByteMatchSet#target_string}.
	TargetString *string `field:"optional" json:"targetString" yaml:"targetString"`
}

