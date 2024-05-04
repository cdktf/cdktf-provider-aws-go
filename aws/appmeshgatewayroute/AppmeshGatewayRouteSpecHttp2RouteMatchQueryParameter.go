// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshgatewayroute


type AppmeshGatewayRouteSpecHttp2RouteMatchQueryParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/appmesh_gateway_route#name AppmeshGatewayRoute#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/resources/appmesh_gateway_route#match AppmeshGatewayRoute#match}
	Match *AppmeshGatewayRouteSpecHttp2RouteMatchQueryParameterMatch `field:"optional" json:"match" yaml:"match"`
}

