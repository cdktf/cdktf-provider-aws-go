package wafv2rulegroup


type Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatement struct {
	// statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/wafv2_rule_group#statement Wafv2RuleGroup#statement}
	Statement interface{} `field:"required" json:"statement" yaml:"statement"`
}

