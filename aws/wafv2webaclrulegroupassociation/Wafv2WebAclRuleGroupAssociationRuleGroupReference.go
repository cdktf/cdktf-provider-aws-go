// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrulegroupassociation


type Wafv2WebAclRuleGroupAssociationRuleGroupReference struct {
	// ARN of the Rule Group to associate with the Web ACL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/wafv2_web_acl_rule_group_association#arn Wafv2WebAclRuleGroupAssociation#arn}
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// rule_action_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/wafv2_web_acl_rule_group_association#rule_action_override Wafv2WebAclRuleGroupAssociation#rule_action_override}
	RuleActionOverride interface{} `field:"optional" json:"ruleActionOverride" yaml:"ruleActionOverride"`
}

