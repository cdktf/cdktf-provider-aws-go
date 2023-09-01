// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewallfirewallpolicy


type NetworkfirewallFirewallPolicyFirewallPolicyStatelessRuleGroupReference struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/networkfirewall_firewall_policy#priority NetworkfirewallFirewallPolicy#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/networkfirewall_firewall_policy#resource_arn NetworkfirewallFirewallPolicy#resource_arn}.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

