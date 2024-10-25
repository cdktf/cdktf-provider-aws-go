// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualservice


type AppmeshVirtualServiceSpecProviderVirtualRouter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/appmesh_virtual_service#virtual_router_name AppmeshVirtualService#virtual_router_name}.
	VirtualRouterName *string `field:"required" json:"virtualRouterName" yaml:"virtualRouterName"`
}

