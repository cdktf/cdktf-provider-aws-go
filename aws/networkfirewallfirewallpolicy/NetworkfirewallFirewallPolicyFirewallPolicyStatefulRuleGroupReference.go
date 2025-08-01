// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewallfirewallpolicy


type NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReference struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/networkfirewall_firewall_policy#resource_arn NetworkfirewallFirewallPolicy#resource_arn}.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/networkfirewall_firewall_policy#deep_threat_inspection NetworkfirewallFirewallPolicy#deep_threat_inspection}.
	DeepThreatInspection *string `field:"optional" json:"deepThreatInspection" yaml:"deepThreatInspection"`
	// override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/networkfirewall_firewall_policy#override NetworkfirewallFirewallPolicy#override}
	Override *NetworkfirewallFirewallPolicyFirewallPolicyStatefulRuleGroupReferenceOverride `field:"optional" json:"override" yaml:"override"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/networkfirewall_firewall_policy#priority NetworkfirewallFirewallPolicy#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

