// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fmspolicy


type FmsPolicySecurityServicePolicyDataPolicyOption struct {
	// network_acl_common_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/fms_policy#network_acl_common_policy FmsPolicy#network_acl_common_policy}
	NetworkAclCommonPolicy *FmsPolicySecurityServicePolicyDataPolicyOptionNetworkAclCommonPolicy `field:"optional" json:"networkAclCommonPolicy" yaml:"networkAclCommonPolicy"`
	// network_firewall_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/fms_policy#network_firewall_policy FmsPolicy#network_firewall_policy}
	NetworkFirewallPolicy *FmsPolicySecurityServicePolicyDataPolicyOptionNetworkFirewallPolicy `field:"optional" json:"networkFirewallPolicy" yaml:"networkFirewallPolicy"`
	// third_party_firewall_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/fms_policy#third_party_firewall_policy FmsPolicy#third_party_firewall_policy}
	ThirdPartyFirewallPolicy *FmsPolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicy `field:"optional" json:"thirdPartyFirewallPolicy" yaml:"thirdPartyFirewallPolicy"`
}

