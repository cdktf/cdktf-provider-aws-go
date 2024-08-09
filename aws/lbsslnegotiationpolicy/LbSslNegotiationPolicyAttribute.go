// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbsslnegotiationpolicy


type LbSslNegotiationPolicyAttribute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/lb_ssl_negotiation_policy#name LbSslNegotiationPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/lb_ssl_negotiation_policy#value LbSslNegotiationPolicy#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

