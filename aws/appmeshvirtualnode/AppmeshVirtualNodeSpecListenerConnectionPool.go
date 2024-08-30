// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualnode


type AppmeshVirtualNodeSpecListenerConnectionPool struct {
	// grpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/appmesh_virtual_node#grpc AppmeshVirtualNode#grpc}
	Grpc *AppmeshVirtualNodeSpecListenerConnectionPoolGrpc `field:"optional" json:"grpc" yaml:"grpc"`
	// http block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/appmesh_virtual_node#http AppmeshVirtualNode#http}
	Http interface{} `field:"optional" json:"http" yaml:"http"`
	// http2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/appmesh_virtual_node#http2 AppmeshVirtualNode#http2}
	Http2 interface{} `field:"optional" json:"http2" yaml:"http2"`
	// tcp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/appmesh_virtual_node#tcp AppmeshVirtualNode#tcp}
	Tcp interface{} `field:"optional" json:"tcp" yaml:"tcp"`
}

