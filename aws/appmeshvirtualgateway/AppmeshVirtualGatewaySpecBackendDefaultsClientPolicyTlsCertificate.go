// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificate struct {
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/appmesh_virtual_gateway#file AppmeshVirtualGateway#file}
	File *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificateFile `field:"optional" json:"file" yaml:"file"`
	// sds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/appmesh_virtual_gateway#sds AppmeshVirtualGateway#sds}
	Sds *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificateSds `field:"optional" json:"sds" yaml:"sds"`
}

