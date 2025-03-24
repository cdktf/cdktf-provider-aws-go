// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fmspolicy


type FmsPolicySecurityServicePolicyDataPolicyOptionNetworkAclCommonPolicyNetworkAclEntrySet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.92.0/docs/resources/fms_policy#force_remediate_for_first_entries FmsPolicy#force_remediate_for_first_entries}.
	ForceRemediateForFirstEntries interface{} `field:"required" json:"forceRemediateForFirstEntries" yaml:"forceRemediateForFirstEntries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.92.0/docs/resources/fms_policy#force_remediate_for_last_entries FmsPolicy#force_remediate_for_last_entries}.
	ForceRemediateForLastEntries interface{} `field:"required" json:"forceRemediateForLastEntries" yaml:"forceRemediateForLastEntries"`
	// first_entry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.92.0/docs/resources/fms_policy#first_entry FmsPolicy#first_entry}
	FirstEntry interface{} `field:"optional" json:"firstEntry" yaml:"firstEntry"`
	// last_entry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.92.0/docs/resources/fms_policy#last_entry FmsPolicy#last_entry}
	LastEntry interface{} `field:"optional" json:"lastEntry" yaml:"lastEntry"`
}

