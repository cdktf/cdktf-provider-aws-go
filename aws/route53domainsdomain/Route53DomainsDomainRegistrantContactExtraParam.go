// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53domainsdomain


type Route53DomainsDomainRegistrantContactExtraParam struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/route53domains_domain#name Route53DomainsDomain#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/route53domains_domain#value Route53DomainsDomain#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

