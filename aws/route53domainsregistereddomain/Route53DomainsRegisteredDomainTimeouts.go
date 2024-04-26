// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53domainsregistereddomain


type Route53DomainsRegisteredDomainTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/route53domains_registered_domain#create Route53DomainsRegisteredDomain#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/route53domains_registered_domain#update Route53DomainsRegisteredDomain#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

