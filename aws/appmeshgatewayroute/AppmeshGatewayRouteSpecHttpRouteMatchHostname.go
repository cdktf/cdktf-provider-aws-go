// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshgatewayroute


type AppmeshGatewayRouteSpecHttpRouteMatchHostname struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/appmesh_gateway_route#exact AppmeshGatewayRoute#exact}.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/appmesh_gateway_route#suffix AppmeshGatewayRoute#suffix}.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

