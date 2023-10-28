// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafsqlinjectionmatchset


type WafSqlInjectionMatchSetSqlInjectionMatchTuples struct {
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/waf_sql_injection_match_set#field_to_match WafSqlInjectionMatchSet#field_to_match}
	FieldToMatch *WafSqlInjectionMatchSetSqlInjectionMatchTuplesFieldToMatch `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/waf_sql_injection_match_set#text_transformation WafSqlInjectionMatchSet#text_transformation}.
	TextTransformation *string `field:"required" json:"textTransformation" yaml:"textTransformation"`
}

