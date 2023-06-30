package wafv2rulegroup


type Wafv2RuleGroupRuleStatementNotStatementStatementOrStatement struct {
	// statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/wafv2_rule_group#statement Wafv2RuleGroup#statement}
	Statement interface{} `field:"required" json:"statement" yaml:"statement"`
}

