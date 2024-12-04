// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2domainname


type Apigatewayv2DomainNameTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/apigatewayv2_domain_name#create Apigatewayv2DomainName#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/apigatewayv2_domain_name#update Apigatewayv2DomainName#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

