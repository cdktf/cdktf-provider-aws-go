package wafv2rulegroup


type Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementIpSetReferenceStatement struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#arn Wafv2RuleGroup#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// ip_set_forwarded_ip_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#ip_set_forwarded_ip_config Wafv2RuleGroup#ip_set_forwarded_ip_config}
	IpSetForwardedIpConfig *Wafv2RuleGroupRuleStatementNotStatementStatementOrStatementStatementIpSetReferenceStatementIpSetForwardedIpConfig `field:"optional" json:"ipSetForwardedIpConfig" yaml:"ipSetForwardedIpConfig"`
}

