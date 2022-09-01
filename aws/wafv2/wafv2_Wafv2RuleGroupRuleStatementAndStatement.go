package wafv2


type Wafv2RuleGroupRuleStatementAndStatement struct {
	// statement block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#statement Wafv2RuleGroup#statement}
	Statement *Wafv2RuleGroupRuleStatement `field:"required" json:"statement" yaml:"statement"`
}

