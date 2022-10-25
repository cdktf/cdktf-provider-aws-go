package networkfirewallfirewallpolicy


type NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReference struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_firewall_policy#resource_arn NetworkfirewallFirewallPolicy#resource_arn}.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// override block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_firewall_policy#override NetworkfirewallFirewallPolicy#override}
	Override *NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceOverride `field:"optional" json:"override" yaml:"override"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/networkfirewall_firewall_policy#priority NetworkfirewallFirewallPolicy#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

