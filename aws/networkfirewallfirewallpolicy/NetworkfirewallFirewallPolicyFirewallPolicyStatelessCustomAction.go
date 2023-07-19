package networkfirewallfirewallpolicy


type NetworkfirewallFirewallPolicyFirewallPolicyStatelessCustomAction struct {
	// action_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/networkfirewall_firewall_policy#action_definition NetworkfirewallFirewallPolicy#action_definition}
	ActionDefinition *NetworkfirewallFirewallPolicyFirewallPolicyStatelessCustomActionActionDefinition `field:"required" json:"actionDefinition" yaml:"actionDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/networkfirewall_firewall_policy#action_name NetworkfirewallFirewallPolicy#action_name}.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
}

