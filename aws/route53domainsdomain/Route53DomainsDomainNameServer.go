// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53domainsdomain


type Route53DomainsDomainNameServer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/route53domains_domain#glue_ips Route53DomainsDomain#glue_ips}.
	GlueIps *[]*string `field:"optional" json:"glueIps" yaml:"glueIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/route53domains_domain#name Route53DomainsDomain#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

