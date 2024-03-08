// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafregionalsqlinjectionmatchset


type WafregionalSqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/wafregional_sql_injection_match_set#type WafregionalSqlInjectionMatchSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/wafregional_sql_injection_match_set#data WafregionalSqlInjectionMatchSet#data}.
	Data *string `field:"optional" json:"data" yaml:"data"`
}

