// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrulegroupassociation


type Wafv2WebAclRuleGroupAssociationRuleGroupReferenceRuleActionOverrideActionToUse struct {
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/wafv2_web_acl_rule_group_association#allow Wafv2WebAclRuleGroupAssociation#allow}
	Allow interface{} `field:"optional" json:"allow" yaml:"allow"`
	// block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/wafv2_web_acl_rule_group_association#block Wafv2WebAclRuleGroupAssociation#block}
	Block interface{} `field:"optional" json:"block" yaml:"block"`
	// captcha block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/wafv2_web_acl_rule_group_association#captcha Wafv2WebAclRuleGroupAssociation#captcha}
	Captcha interface{} `field:"optional" json:"captcha" yaml:"captcha"`
	// challenge block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/wafv2_web_acl_rule_group_association#challenge Wafv2WebAclRuleGroupAssociation#challenge}
	Challenge interface{} `field:"optional" json:"challenge" yaml:"challenge"`
	// count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/wafv2_web_acl_rule_group_association#count Wafv2WebAclRuleGroupAssociation#count}
	Count interface{} `field:"optional" json:"count" yaml:"count"`
}

