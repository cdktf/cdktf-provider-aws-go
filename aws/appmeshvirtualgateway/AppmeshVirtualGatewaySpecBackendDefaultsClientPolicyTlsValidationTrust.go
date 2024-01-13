// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrust struct {
	// acm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/appmesh_virtual_gateway#acm AppmeshVirtualGateway#acm}
	Acm *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustAcm `field:"optional" json:"acm" yaml:"acm"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/appmesh_virtual_gateway#file AppmeshVirtualGateway#file}
	File *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustFile `field:"optional" json:"file" yaml:"file"`
	// sds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/appmesh_virtual_gateway#sds AppmeshVirtualGateway#sds}
	Sds *AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustSds `field:"optional" json:"sds" yaml:"sds"`
}

