// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualnode


type AppmeshVirtualNodeSpecServiceDiscoveryAwsCloudMap struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/appmesh_virtual_node#namespace_name AppmeshVirtualNode#namespace_name}.
	NamespaceName *string `field:"required" json:"namespaceName" yaml:"namespaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/appmesh_virtual_node#service_name AppmeshVirtualNode#service_name}.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/appmesh_virtual_node#attributes AppmeshVirtualNode#attributes}.
	Attributes *map[string]*string `field:"optional" json:"attributes" yaml:"attributes"`
}

