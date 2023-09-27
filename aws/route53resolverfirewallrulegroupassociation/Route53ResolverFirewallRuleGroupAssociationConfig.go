// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53resolverfirewallrulegroupassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53ResolverFirewallRuleGroupAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/route53_resolver_firewall_rule_group_association#firewall_rule_group_id Route53ResolverFirewallRuleGroupAssociation#firewall_rule_group_id}.
	FirewallRuleGroupId *string `field:"required" json:"firewallRuleGroupId" yaml:"firewallRuleGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/route53_resolver_firewall_rule_group_association#name Route53ResolverFirewallRuleGroupAssociation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/route53_resolver_firewall_rule_group_association#priority Route53ResolverFirewallRuleGroupAssociation#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/route53_resolver_firewall_rule_group_association#vpc_id Route53ResolverFirewallRuleGroupAssociation#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/route53_resolver_firewall_rule_group_association#id Route53ResolverFirewallRuleGroupAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/route53_resolver_firewall_rule_group_association#mutation_protection Route53ResolverFirewallRuleGroupAssociation#mutation_protection}.
	MutationProtection *string `field:"optional" json:"mutationProtection" yaml:"mutationProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/route53_resolver_firewall_rule_group_association#tags Route53ResolverFirewallRuleGroupAssociation#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/route53_resolver_firewall_rule_group_association#tags_all Route53ResolverFirewallRuleGroupAssociation#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

