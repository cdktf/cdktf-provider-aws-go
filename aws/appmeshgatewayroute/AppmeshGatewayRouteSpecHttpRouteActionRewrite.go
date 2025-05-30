// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshgatewayroute


type AppmeshGatewayRouteSpecHttpRouteActionRewrite struct {
	// hostname block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/appmesh_gateway_route#hostname AppmeshGatewayRoute#hostname}
	Hostname *AppmeshGatewayRouteSpecHttpRouteActionRewriteHostname `field:"optional" json:"hostname" yaml:"hostname"`
	// path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/appmesh_gateway_route#path AppmeshGatewayRoute#path}
	Path *AppmeshGatewayRouteSpecHttpRouteActionRewritePath `field:"optional" json:"path" yaml:"path"`
	// prefix block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/appmesh_gateway_route#prefix AppmeshGatewayRoute#prefix}
	Prefix *AppmeshGatewayRouteSpecHttpRouteActionRewritePrefix `field:"optional" json:"prefix" yaml:"prefix"`
}

