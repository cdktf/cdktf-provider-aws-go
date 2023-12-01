// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecBackendDefaults struct {
	// client_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/appmesh_virtual_gateway#client_policy AppmeshVirtualGateway#client_policy}
	ClientPolicy *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicy `field:"optional" json:"clientPolicy" yaml:"clientPolicy"`
}

