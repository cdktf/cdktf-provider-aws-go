// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualservice


type AppmeshVirtualServiceSpecProviderVirtualNode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/appmesh_virtual_service#virtual_node_name AppmeshVirtualService#virtual_node_name}.
	VirtualNodeName *string `field:"required" json:"virtualNodeName" yaml:"virtualNodeName"`
}

