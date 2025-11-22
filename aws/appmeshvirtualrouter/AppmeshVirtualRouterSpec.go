// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualrouter


type AppmeshVirtualRouterSpec struct {
	// listener block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/appmesh_virtual_router#listener AppmeshVirtualRouter#listener}
	Listener interface{} `field:"optional" json:"listener" yaml:"listener"`
}

