// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fmspolicy


type FmsPolicySecurityServicePolicyDataPolicyOptionNetworkAclCommonPolicy struct {
	// network_acl_entry_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/fms_policy#network_acl_entry_set FmsPolicy#network_acl_entry_set}
	NetworkAclEntrySet *FmsPolicySecurityServicePolicyDataPolicyOptionNetworkAclCommonPolicyNetworkAclEntrySet `field:"optional" json:"networkAclEntrySet" yaml:"networkAclEntrySet"`
}

