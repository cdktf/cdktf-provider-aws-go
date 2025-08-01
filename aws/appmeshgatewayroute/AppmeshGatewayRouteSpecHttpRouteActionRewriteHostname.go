// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshgatewayroute


type AppmeshGatewayRouteSpecHttpRouteActionRewriteHostname struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/appmesh_gateway_route#default_target_hostname AppmeshGatewayRoute#default_target_hostname}.
	DefaultTargetHostname *string `field:"required" json:"defaultTargetHostname" yaml:"defaultTargetHostname"`
}

