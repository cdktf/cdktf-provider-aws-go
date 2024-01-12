// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualnode


type AppmeshVirtualNodeSpecBackendDefaults struct {
	// client_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/appmesh_virtual_node#client_policy AppmeshVirtualNode#client_policy}
	ClientPolicy *AppmeshVirtualNodeSpecBackendDefaultsClientPolicy `field:"optional" json:"clientPolicy" yaml:"clientPolicy"`
}

