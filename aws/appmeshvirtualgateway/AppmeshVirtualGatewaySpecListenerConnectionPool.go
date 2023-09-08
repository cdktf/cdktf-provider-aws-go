// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecListenerConnectionPool struct {
	// grpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.0/docs/resources/appmesh_virtual_gateway#grpc AppmeshVirtualGateway#grpc}
	Grpc *AppmeshVirtualGatewaySpecListenerConnectionPoolGrpc `field:"optional" json:"grpc" yaml:"grpc"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.0/docs/resources/appmesh_virtual_gateway#http AppmeshVirtualGateway#http}
	Http *AppmeshVirtualGatewaySpecListenerConnectionPoolHttp `field:"optional" json:"http" yaml:"http"`
	// http2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.0/docs/resources/appmesh_virtual_gateway#http2 AppmeshVirtualGateway#http2}
	Http2 *AppmeshVirtualGatewaySpecListenerConnectionPoolHttp2 `field:"optional" json:"http2" yaml:"http2"`
}

