package networkfirewall


type NetworkfirewallRuleGroupRuleGroupRuleVariablesIpSets struct {
	// ip_set block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_rule_group#ip_set NetworkfirewallRuleGroup#ip_set}
	IpSet *NetworkfirewallRuleGroupRuleGroupRuleVariablesIpSetsIpSet `field:"required" json:"ipSet" yaml:"ipSet"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_rule_group#key NetworkfirewallRuleGroup#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
}

