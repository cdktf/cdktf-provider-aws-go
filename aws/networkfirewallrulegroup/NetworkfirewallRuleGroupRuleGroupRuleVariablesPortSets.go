package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupRuleVariablesPortSets struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_rule_group#key NetworkfirewallRuleGroup#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// port_set block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_rule_group#port_set NetworkfirewallRuleGroup#port_set}
	PortSet *NetworkfirewallRuleGroupRuleGroupRuleVariablesPortSetsPortSet `field:"required" json:"portSet" yaml:"portSet"`
}

