package wafv2rulegroup


type Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatementStatementLabelMatchStatement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/wafv2_rule_group#key Wafv2RuleGroup#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/wafv2_rule_group#scope Wafv2RuleGroup#scope}.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

