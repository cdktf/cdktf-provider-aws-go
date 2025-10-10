// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrulegroupassociation


type Wafv2WebAclRuleGroupAssociationManagedRuleGroup struct {
	// Name of the managed rule group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/wafv2_web_acl_rule_group_association#name Wafv2WebAclRuleGroupAssociation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of the managed rule group vendor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/wafv2_web_acl_rule_group_association#vendor_name Wafv2WebAclRuleGroupAssociation#vendor_name}
	VendorName *string `field:"required" json:"vendorName" yaml:"vendorName"`
	// rule_action_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/wafv2_web_acl_rule_group_association#rule_action_override Wafv2WebAclRuleGroupAssociation#rule_action_override}
	RuleActionOverride interface{} `field:"optional" json:"ruleActionOverride" yaml:"ruleActionOverride"`
	// Version of the managed rule group. Omit this to use the default version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/wafv2_web_acl_rule_group_association#version Wafv2WebAclRuleGroupAssociation#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

