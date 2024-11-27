// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualrouter


type AppmeshVirtualRouterSpecListenerPortMapping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/appmesh_virtual_router#port AppmeshVirtualRouter#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/appmesh_virtual_router#protocol AppmeshVirtualRouter#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
}

