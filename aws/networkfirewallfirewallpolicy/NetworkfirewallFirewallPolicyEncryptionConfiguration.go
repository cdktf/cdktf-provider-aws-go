// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewallfirewallpolicy


type NetworkfirewallFirewallPolicyEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/networkfirewall_firewall_policy#type NetworkfirewallFirewallPolicy#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/networkfirewall_firewall_policy#key_id NetworkfirewallFirewallPolicy#key_id}.
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
}

