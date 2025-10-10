// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualnode


type AppmeshVirtualNodeSpecBackendVirtualService struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appmesh_virtual_node#virtual_service_name AppmeshVirtualNode#virtual_service_name}.
	VirtualServiceName *string `field:"required" json:"virtualServiceName" yaml:"virtualServiceName"`
	// client_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appmesh_virtual_node#client_policy AppmeshVirtualNode#client_policy}
	ClientPolicy *AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicy `field:"optional" json:"clientPolicy" yaml:"clientPolicy"`
}

