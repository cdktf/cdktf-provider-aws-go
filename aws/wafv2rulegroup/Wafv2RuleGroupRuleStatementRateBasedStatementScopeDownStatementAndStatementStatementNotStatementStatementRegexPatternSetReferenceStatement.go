// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2rulegroup


type Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementNotStatementStatementRegexPatternSetReferenceStatement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/wafv2_rule_group#arn Wafv2RuleGroup#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// text_transformation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/wafv2_rule_group#text_transformation Wafv2RuleGroup#text_transformation}
	TextTransformation interface{} `field:"required" json:"textTransformation" yaml:"textTransformation"`
	// field_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/wafv2_rule_group#field_to_match Wafv2RuleGroup#field_to_match}
	FieldToMatch *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatch `field:"optional" json:"fieldToMatch" yaml:"fieldToMatch"`
}

