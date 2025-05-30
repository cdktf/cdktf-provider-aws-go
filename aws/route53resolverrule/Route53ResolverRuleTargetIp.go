// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53resolverrule


type Route53ResolverRuleTargetIp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/route53_resolver_rule#ip Route53ResolverRule#ip}.
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/route53_resolver_rule#ipv6 Route53ResolverRule#ipv6}.
	Ipv6 *string `field:"optional" json:"ipv6" yaml:"ipv6"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/route53_resolver_rule#port Route53ResolverRule#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/route53_resolver_rule#protocol Route53ResolverRule#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

