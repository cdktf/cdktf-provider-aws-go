// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshgatewayroute


type AppmeshGatewayRouteSpec struct {
	// grpc_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/appmesh_gateway_route#grpc_route AppmeshGatewayRoute#grpc_route}
	GrpcRoute *AppmeshGatewayRouteSpecGrpcRoute `field:"optional" json:"grpcRoute" yaml:"grpcRoute"`
	// http2_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/appmesh_gateway_route#http2_route AppmeshGatewayRoute#http2_route}
	Http2Route *AppmeshGatewayRouteSpecHttp2Route `field:"optional" json:"http2Route" yaml:"http2Route"`
	// http_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/appmesh_gateway_route#http_route AppmeshGatewayRoute#http_route}
	HttpRoute *AppmeshGatewayRouteSpecHttpRoute `field:"optional" json:"httpRoute" yaml:"httpRoute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/appmesh_gateway_route#priority AppmeshGatewayRoute#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

