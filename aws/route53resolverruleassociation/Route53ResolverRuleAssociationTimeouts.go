// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53resolverruleassociation


type Route53ResolverRuleAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/route53_resolver_rule_association#create Route53ResolverRuleAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/route53_resolver_rule_association#delete Route53ResolverRuleAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

