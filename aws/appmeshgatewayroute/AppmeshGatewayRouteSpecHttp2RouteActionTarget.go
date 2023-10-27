// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshgatewayroute


type AppmeshGatewayRouteSpecHttp2RouteActionTarget struct {
	// virtual_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/appmesh_gateway_route#virtual_service AppmeshGatewayRoute#virtual_service}
	VirtualService *AppmeshGatewayRouteSpecHttp2RouteActionTargetVirtualService `field:"required" json:"virtualService" yaml:"virtualService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/appmesh_gateway_route#port AppmeshGatewayRoute#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

