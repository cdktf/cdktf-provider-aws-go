// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualnode


type AppmeshVirtualNodeSpecBackend struct {
	// virtual_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/appmesh_virtual_node#virtual_service AppmeshVirtualNode#virtual_service}
	VirtualService *AppmeshVirtualNodeSpecBackendVirtualService `field:"required" json:"virtualService" yaml:"virtualService"`
}

