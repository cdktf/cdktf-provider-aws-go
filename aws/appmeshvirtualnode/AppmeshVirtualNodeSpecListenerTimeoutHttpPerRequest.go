// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualnode


type AppmeshVirtualNodeSpecListenerTimeoutHttpPerRequest struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/appmesh_virtual_node#unit AppmeshVirtualNode#unit}.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.2/docs/resources/appmesh_virtual_node#value AppmeshVirtualNode#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

