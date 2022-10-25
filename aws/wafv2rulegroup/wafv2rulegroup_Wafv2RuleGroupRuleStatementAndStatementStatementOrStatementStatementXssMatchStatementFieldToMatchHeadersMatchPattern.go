package wafv2rulegroup


type Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementXssMatchStatementFieldToMatchHeadersMatchPattern struct {
	// all block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#all Wafv2RuleGroup#all}
	All *Wafv2RuleGroupRuleStatementAndStatementStatementOrStatementStatementXssMatchStatementFieldToMatchHeadersMatchPatternAll `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#excluded_headers Wafv2RuleGroup#excluded_headers}.
	ExcludedHeaders *[]*string `field:"optional" json:"excludedHeaders" yaml:"excludedHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#included_headers Wafv2RuleGroup#included_headers}.
	IncludedHeaders *[]*string `field:"optional" json:"includedHeaders" yaml:"includedHeaders"`
}

