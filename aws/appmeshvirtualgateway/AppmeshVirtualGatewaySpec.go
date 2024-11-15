// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualgateway


type AppmeshVirtualGatewaySpec struct {
	// listener block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/appmesh_virtual_gateway#listener AppmeshVirtualGateway#listener}
	Listener interface{} `field:"required" json:"listener" yaml:"listener"`
	// backend_defaults block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/appmesh_virtual_gateway#backend_defaults AppmeshVirtualGateway#backend_defaults}
	BackendDefaults *AppmeshVirtualGatewaySpecBackendDefaults `field:"optional" json:"backendDefaults" yaml:"backendDefaults"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/appmesh_virtual_gateway#logging AppmeshVirtualGateway#logging}
	Logging *AppmeshVirtualGatewaySpecLogging `field:"optional" json:"logging" yaml:"logging"`
}

