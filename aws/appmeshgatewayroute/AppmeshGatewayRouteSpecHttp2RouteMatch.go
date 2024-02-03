// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshgatewayroute


type AppmeshGatewayRouteSpecHttp2RouteMatch struct {
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/appmesh_gateway_route#header AppmeshGatewayRoute#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// hostname block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/appmesh_gateway_route#hostname AppmeshGatewayRoute#hostname}
	Hostname *AppmeshGatewayRouteSpecHttp2RouteMatchHostname `field:"optional" json:"hostname" yaml:"hostname"`
	// path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/appmesh_gateway_route#path AppmeshGatewayRoute#path}
	Path *AppmeshGatewayRouteSpecHttp2RouteMatchPath `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/appmesh_gateway_route#port AppmeshGatewayRoute#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/appmesh_gateway_route#prefix AppmeshGatewayRoute#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// query_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/appmesh_gateway_route#query_parameter AppmeshGatewayRoute#query_parameter}
	QueryParameter interface{} `field:"optional" json:"queryParameter" yaml:"queryParameter"`
}

