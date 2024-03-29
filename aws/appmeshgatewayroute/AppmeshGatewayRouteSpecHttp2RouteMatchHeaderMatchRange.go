// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshgatewayroute


type AppmeshGatewayRouteSpecHttp2RouteMatchHeaderMatchRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/appmesh_gateway_route#end AppmeshGatewayRoute#end}.
	End *float64 `field:"required" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/appmesh_gateway_route#start AppmeshGatewayRoute#start}.
	Start *float64 `field:"required" json:"start" yaml:"start"`
}

