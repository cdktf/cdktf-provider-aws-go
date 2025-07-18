// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualnode


type AppmeshVirtualNodeSpec struct {
	// backend block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/appmesh_virtual_node#backend AppmeshVirtualNode#backend}
	Backend interface{} `field:"optional" json:"backend" yaml:"backend"`
	// backend_defaults block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/appmesh_virtual_node#backend_defaults AppmeshVirtualNode#backend_defaults}
	BackendDefaults *AppmeshVirtualNodeSpecBackendDefaults `field:"optional" json:"backendDefaults" yaml:"backendDefaults"`
	// listener block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/appmesh_virtual_node#listener AppmeshVirtualNode#listener}
	Listener interface{} `field:"optional" json:"listener" yaml:"listener"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/appmesh_virtual_node#logging AppmeshVirtualNode#logging}
	Logging *AppmeshVirtualNodeSpecLogging `field:"optional" json:"logging" yaml:"logging"`
	// service_discovery block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/appmesh_virtual_node#service_discovery AppmeshVirtualNode#service_discovery}
	ServiceDiscovery *AppmeshVirtualNodeSpecServiceDiscovery `field:"optional" json:"serviceDiscovery" yaml:"serviceDiscovery"`
}

