// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewaydomainname


type ApiGatewayDomainNameEndpointConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/api_gateway_domain_name#types ApiGatewayDomainName#types}.
	Types *[]*string `field:"required" json:"types" yaml:"types"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/api_gateway_domain_name#ip_address_type ApiGatewayDomainName#ip_address_type}.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
}

