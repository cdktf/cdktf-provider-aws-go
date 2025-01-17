// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsroute53resolverquerylogconfig


type DataAwsRoute53ResolverQueryLogConfigFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/data-sources/route53_resolver_query_log_config#name DataAwsRoute53ResolverQueryLogConfig#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.84.0/docs/data-sources/route53_resolver_query_log_config#values DataAwsRoute53ResolverQueryLogConfig#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

