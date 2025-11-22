// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewallfirewallpolicy


type NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariables struct {
	// ip_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/networkfirewall_firewall_policy#ip_set NetworkfirewallFirewallPolicy#ip_set}
	IpSet *NetworkfirewallFirewallPolicyFirewallPolicyPolicyVariablesRuleVariablesIpSet `field:"required" json:"ipSet" yaml:"ipSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/networkfirewall_firewall_policy#key NetworkfirewallFirewallPolicy#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
}

