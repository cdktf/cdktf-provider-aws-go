package wafv2rulegroup


type Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatementStatementNotStatementStatementByteMatchStatementFieldToMatchCookies struct {
	// match_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/wafv2_rule_group#match_pattern Wafv2RuleGroup#match_pattern}
	MatchPattern interface{} `field:"required" json:"matchPattern" yaml:"matchPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/wafv2_rule_group#match_scope Wafv2RuleGroup#match_scope}.
	MatchScope *string `field:"required" json:"matchScope" yaml:"matchScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/wafv2_rule_group#oversize_handling Wafv2RuleGroup#oversize_handling}.
	OversizeHandling *string `field:"required" json:"oversizeHandling" yaml:"oversizeHandling"`
}

