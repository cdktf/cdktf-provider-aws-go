// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshgatewayroute


type AppmeshGatewayRouteSpecHttpRouteMatchHeaderMatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/appmesh_gateway_route#exact AppmeshGatewayRoute#exact}.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/appmesh_gateway_route#prefix AppmeshGatewayRoute#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/appmesh_gateway_route#range AppmeshGatewayRoute#range}
	Range *AppmeshGatewayRouteSpecHttpRouteMatchHeaderMatchRange `field:"optional" json:"range" yaml:"range"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/appmesh_gateway_route#regex AppmeshGatewayRoute#regex}.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/appmesh_gateway_route#suffix AppmeshGatewayRoute#suffix}.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

