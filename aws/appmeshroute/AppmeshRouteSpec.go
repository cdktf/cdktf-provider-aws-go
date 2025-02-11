// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshroute


type AppmeshRouteSpec struct {
	// grpc_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/appmesh_route#grpc_route AppmeshRoute#grpc_route}
	GrpcRoute *AppmeshRouteSpecGrpcRoute `field:"optional" json:"grpcRoute" yaml:"grpcRoute"`
	// http2_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/appmesh_route#http2_route AppmeshRoute#http2_route}
	Http2Route *AppmeshRouteSpecHttp2Route `field:"optional" json:"http2Route" yaml:"http2Route"`
	// http_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/appmesh_route#http_route AppmeshRoute#http_route}
	HttpRoute *AppmeshRouteSpecHttpRoute `field:"optional" json:"httpRoute" yaml:"httpRoute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/appmesh_route#priority AppmeshRoute#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// tcp_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.1/docs/resources/appmesh_route#tcp_route AppmeshRoute#tcp_route}
	TcpRoute *AppmeshRouteSpecTcpRoute `field:"optional" json:"tcpRoute" yaml:"tcpRoute"`
}

