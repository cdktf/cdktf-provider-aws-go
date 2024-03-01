// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecBackendDefaultsClientPolicy struct {
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.0/docs/resources/appmesh_virtual_gateway#tls AppmeshVirtualGateway#tls}
	Tls *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTls `field:"optional" json:"tls" yaml:"tls"`
}

