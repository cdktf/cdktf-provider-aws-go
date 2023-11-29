// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2rulegroup


type Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatement struct {
	// byte_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/wafv2_rule_group#byte_match_statement Wafv2RuleGroup#byte_match_statement}
	ByteMatchStatement *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementByteMatchStatement `field:"optional" json:"byteMatchStatement" yaml:"byteMatchStatement"`
	// geo_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/wafv2_rule_group#geo_match_statement Wafv2RuleGroup#geo_match_statement}
	GeoMatchStatement *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementGeoMatchStatement `field:"optional" json:"geoMatchStatement" yaml:"geoMatchStatement"`
	// ip_set_reference_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/wafv2_rule_group#ip_set_reference_statement Wafv2RuleGroup#ip_set_reference_statement}
	IpSetReferenceStatement *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementIpSetReferenceStatement `field:"optional" json:"ipSetReferenceStatement" yaml:"ipSetReferenceStatement"`
	// label_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/wafv2_rule_group#label_match_statement Wafv2RuleGroup#label_match_statement}
	LabelMatchStatement *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementLabelMatchStatement `field:"optional" json:"labelMatchStatement" yaml:"labelMatchStatement"`
	// regex_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/wafv2_rule_group#regex_match_statement Wafv2RuleGroup#regex_match_statement}
	RegexMatchStatement *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexMatchStatement `field:"optional" json:"regexMatchStatement" yaml:"regexMatchStatement"`
	// regex_pattern_set_reference_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/wafv2_rule_group#regex_pattern_set_reference_statement Wafv2RuleGroup#regex_pattern_set_reference_statement}
	RegexPatternSetReferenceStatement *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementRegexPatternSetReferenceStatement `field:"optional" json:"regexPatternSetReferenceStatement" yaml:"regexPatternSetReferenceStatement"`
	// size_constraint_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/wafv2_rule_group#size_constraint_statement Wafv2RuleGroup#size_constraint_statement}
	SizeConstraintStatement *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSizeConstraintStatement `field:"optional" json:"sizeConstraintStatement" yaml:"sizeConstraintStatement"`
	// sqli_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/wafv2_rule_group#sqli_match_statement Wafv2RuleGroup#sqli_match_statement}
	SqliMatchStatement *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementSqliMatchStatement `field:"optional" json:"sqliMatchStatement" yaml:"sqliMatchStatement"`
	// xss_match_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/wafv2_rule_group#xss_match_statement Wafv2RuleGroup#xss_match_statement}
	XssMatchStatement *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementNotStatementStatementXssMatchStatement `field:"optional" json:"xssMatchStatement" yaml:"xssMatchStatement"`
}

