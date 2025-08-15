// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fmspolicy


type FmsPolicySecurityServicePolicyDataPolicyOptionNetworkAclCommonPolicyNetworkAclEntrySetLastEntryPortRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/fms_policy#from FmsPolicy#from}.
	From *float64 `field:"optional" json:"from" yaml:"from"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/fms_policy#to FmsPolicy#to}.
	To *float64 `field:"optional" json:"to" yaml:"to"`
}

