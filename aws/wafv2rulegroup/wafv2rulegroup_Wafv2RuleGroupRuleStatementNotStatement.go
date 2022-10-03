package wafv2rulegroup


type Wafv2RuleGroupRuleStatementNotStatement struct {
	// statement block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#statement Wafv2RuleGroup#statement}
	Statement *Wafv2RuleGroupRuleStatement `field:"required" json:"statement" yaml:"statement"`
}

