// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsroute53resolverendpoint


type DataAwsRoute53ResolverEndpointFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/data-sources/route53_resolver_endpoint#name DataAwsRoute53ResolverEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/data-sources/route53_resolver_endpoint#values DataAwsRoute53ResolverEndpoint#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

