// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupReferenceSets struct {
	// ip_set_references block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/networkfirewall_rule_group#ip_set_references NetworkfirewallRuleGroup#ip_set_references}
	IpSetReferences interface{} `field:"optional" json:"ipSetReferences" yaml:"ipSetReferences"`
}

