// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webaclrulegroupassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2WebAclRuleGroupAssociationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Priority of the rule within the Web ACL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/wafv2_web_acl_rule_group_association#priority Wafv2WebAclRuleGroupAssociation#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Name of the rule to create in the Web ACL that references the rule group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/wafv2_web_acl_rule_group_association#rule_name Wafv2WebAclRuleGroupAssociation#rule_name}
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// ARN of the Web ACL to associate the Rule Group with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/wafv2_web_acl_rule_group_association#web_acl_arn Wafv2WebAclRuleGroupAssociation#web_acl_arn}
	WebAclArn *string `field:"required" json:"webAclArn" yaml:"webAclArn"`
	// managed_rule_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/wafv2_web_acl_rule_group_association#managed_rule_group Wafv2WebAclRuleGroupAssociation#managed_rule_group}
	ManagedRuleGroup interface{} `field:"optional" json:"managedRuleGroup" yaml:"managedRuleGroup"`
	// Override action for the rule group. Valid values are 'none' and 'count'. Defaults to 'none'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/wafv2_web_acl_rule_group_association#override_action Wafv2WebAclRuleGroupAssociation#override_action}
	OverrideAction *string `field:"optional" json:"overrideAction" yaml:"overrideAction"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/wafv2_web_acl_rule_group_association#region Wafv2WebAclRuleGroupAssociation#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// rule_group_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/wafv2_web_acl_rule_group_association#rule_group_reference Wafv2WebAclRuleGroupAssociation#rule_group_reference}
	RuleGroupReference interface{} `field:"optional" json:"ruleGroupReference" yaml:"ruleGroupReference"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/wafv2_web_acl_rule_group_association#timeouts Wafv2WebAclRuleGroupAssociation#timeouts}
	Timeouts *Wafv2WebAclRuleGroupAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

