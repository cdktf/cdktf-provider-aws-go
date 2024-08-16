// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualnode


type AppmeshVirtualNodeSpecListenerTimeoutGrpc struct {
	// idle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/appmesh_virtual_node#idle AppmeshVirtualNode#idle}
	Idle *AppmeshVirtualNodeSpecListenerTimeoutGrpcIdle `field:"optional" json:"idle" yaml:"idle"`
	// per_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/appmesh_virtual_node#per_request AppmeshVirtualNode#per_request}
	PerRequest *AppmeshVirtualNodeSpecListenerTimeoutGrpcPerRequest `field:"optional" json:"perRequest" yaml:"perRequest"`
}

