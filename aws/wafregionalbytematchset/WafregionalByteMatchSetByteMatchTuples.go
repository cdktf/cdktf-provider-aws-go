// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafregionalbytematchset


type WafregionalByteMatchSetByteMatchTuples struct {
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/wafregional_byte_match_set#field_to_match WafregionalByteMatchSet#field_to_match}
	FieldToMatch *WafregionalByteMatchSetByteMatchTuplesFieldToMatch `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/wafregional_byte_match_set#positional_constraint WafregionalByteMatchSet#positional_constraint}.
	PositionalConstraint *string `field:"required" json:"positionalConstraint" yaml:"positionalConstraint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/wafregional_byte_match_set#text_transformation WafregionalByteMatchSet#text_transformation}.
	TextTransformation *string `field:"required" json:"textTransformation" yaml:"textTransformation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/wafregional_byte_match_set#target_string WafregionalByteMatchSet#target_string}.
	TargetString *string `field:"optional" json:"targetString" yaml:"targetString"`
}

