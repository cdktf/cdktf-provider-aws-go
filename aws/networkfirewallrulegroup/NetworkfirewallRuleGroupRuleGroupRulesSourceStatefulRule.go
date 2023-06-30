package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/networkfirewall_rule_group#action NetworkfirewallRuleGroup#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/networkfirewall_rule_group#header NetworkfirewallRuleGroup#header}
	Header *NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRuleHeader `field:"required" json:"header" yaml:"header"`
	// rule_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/networkfirewall_rule_group#rule_option NetworkfirewallRuleGroup#rule_option}
	RuleOption interface{} `field:"required" json:"ruleOption" yaml:"ruleOption"`
}

