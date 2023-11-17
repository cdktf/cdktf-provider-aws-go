// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupReferenceSetsIpSetReferences struct {
	// ip_set_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/networkfirewall_rule_group#ip_set_reference NetworkfirewallRuleGroup#ip_set_reference}
	IpSetReference interface{} `field:"required" json:"ipSetReference" yaml:"ipSetReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/networkfirewall_rule_group#key NetworkfirewallRuleGroup#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
}

