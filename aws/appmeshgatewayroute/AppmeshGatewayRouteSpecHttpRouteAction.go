// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshgatewayroute


type AppmeshGatewayRouteSpecHttpRouteAction struct {
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/appmesh_gateway_route#target AppmeshGatewayRoute#target}
	Target *AppmeshGatewayRouteSpecHttpRouteActionTarget `field:"required" json:"target" yaml:"target"`
	// rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/appmesh_gateway_route#rewrite AppmeshGatewayRoute#rewrite}
	Rewrite *AppmeshGatewayRouteSpecHttpRouteActionRewrite `field:"optional" json:"rewrite" yaml:"rewrite"`
}

