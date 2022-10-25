package wafv2rulegroup


type Wafv2RuleGroupRuleStatementNotStatementStatementNotStatementStatementByteMatchStatementFieldToMatchCookiesMatchPattern struct {
	// all block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#all Wafv2RuleGroup#all}
	All *Wafv2RuleGroupRuleStatementNotStatementStatementNotStatementStatementByteMatchStatementFieldToMatchCookiesMatchPatternAll `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#excluded_cookies Wafv2RuleGroup#excluded_cookies}.
	ExcludedCookies *[]*string `field:"optional" json:"excludedCookies" yaml:"excludedCookies"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#included_cookies Wafv2RuleGroup#included_cookies}.
	IncludedCookies *[]*string `field:"optional" json:"includedCookies" yaml:"includedCookies"`
}

