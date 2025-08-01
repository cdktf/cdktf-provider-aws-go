// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshgatewayroute


type AppmeshGatewayRouteSpecHttpRouteActionRewritePrefix struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/appmesh_gateway_route#default_prefix AppmeshGatewayRoute#default_prefix}.
	DefaultPrefix *string `field:"optional" json:"defaultPrefix" yaml:"defaultPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/appmesh_gateway_route#value AppmeshGatewayRoute#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

