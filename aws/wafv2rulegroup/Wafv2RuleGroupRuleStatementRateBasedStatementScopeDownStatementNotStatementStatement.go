// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2rulegroup


type Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatement struct {
	// and_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/wafv2_rule_group#and_statement Wafv2RuleGroup#and_statement}
	AndStatement *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementAndStatement `field:"optional" json:"andStatement" yaml:"andStatement"`
	// byte_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/wafv2_rule_group#byte_match_statement Wafv2RuleGroup#byte_match_statement}
	ByteMatchStatement *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementByteMatchStatement `field:"optional" json:"byteMatchStatement" yaml:"byteMatchStatement"`
	// geo_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/wafv2_rule_group#geo_match_statement Wafv2RuleGroup#geo_match_statement}
	GeoMatchStatement *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementGeoMatchStatement `field:"optional" json:"geoMatchStatement" yaml:"geoMatchStatement"`
	// ip_set_reference_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/wafv2_rule_group#ip_set_reference_statement Wafv2RuleGroup#ip_set_reference_statement}
	IpSetReferenceStatement *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementIpSetReferenceStatement `field:"optional" json:"ipSetReferenceStatement" yaml:"ipSetReferenceStatement"`
	// label_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/wafv2_rule_group#label_match_statement Wafv2RuleGroup#label_match_statement}
	LabelMatchStatement *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementLabelMatchStatement `field:"optional" json:"labelMatchStatement" yaml:"labelMatchStatement"`
	// not_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/wafv2_rule_group#not_statement Wafv2RuleGroup#not_statement}
	NotStatement *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatement `field:"optional" json:"notStatement" yaml:"notStatement"`
	// or_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/wafv2_rule_group#or_statement Wafv2RuleGroup#or_statement}
	OrStatement *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementOrStatement `field:"optional" json:"orStatement" yaml:"orStatement"`
	// regex_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/wafv2_rule_group#regex_match_statement Wafv2RuleGroup#regex_match_statement}
	RegexMatchStatement *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementRegexMatchStatement `field:"optional" json:"regexMatchStatement" yaml:"regexMatchStatement"`
	// regex_pattern_set_reference_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/wafv2_rule_group#regex_pattern_set_reference_statement Wafv2RuleGroup#regex_pattern_set_reference_statement}
	RegexPatternSetReferenceStatement *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementRegexPatternSetReferenceStatement `field:"optional" json:"regexPatternSetReferenceStatement" yaml:"regexPatternSetReferenceStatement"`
	// size_constraint_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/wafv2_rule_group#size_constraint_statement Wafv2RuleGroup#size_constraint_statement}
	SizeConstraintStatement *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSizeConstraintStatement `field:"optional" json:"sizeConstraintStatement" yaml:"sizeConstraintStatement"`
	// sqli_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/wafv2_rule_group#sqli_match_statement Wafv2RuleGroup#sqli_match_statement}
	SqliMatchStatement *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementSqliMatchStatement `field:"optional" json:"sqliMatchStatement" yaml:"sqliMatchStatement"`
	// xss_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/wafv2_rule_group#xss_match_statement Wafv2RuleGroup#xss_match_statement}
	XssMatchStatement *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementXssMatchStatement `field:"optional" json:"xssMatchStatement" yaml:"xssMatchStatement"`
}

