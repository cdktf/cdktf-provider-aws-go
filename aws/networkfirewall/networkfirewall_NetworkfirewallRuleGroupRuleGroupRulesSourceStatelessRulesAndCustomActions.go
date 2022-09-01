package networkfirewall


type NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActions struct {
	// stateless_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_rule_group#stateless_rule NetworkfirewallRuleGroup#stateless_rule}
	StatelessRule interface{} `field:"required" json:"statelessRule" yaml:"statelessRule"`
	// custom_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_rule_group#custom_action NetworkfirewallRuleGroup#custom_action}
	CustomAction interface{} `field:"optional" json:"customAction" yaml:"customAction"`
}

