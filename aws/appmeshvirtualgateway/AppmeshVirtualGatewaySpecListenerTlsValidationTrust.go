// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecListenerTlsValidationTrust struct {
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/appmesh_virtual_gateway#file AppmeshVirtualGateway#file}
	File *AppmeshVirtualGatewaySpecListenerTlsValidationTrustFile `field:"optional" json:"file" yaml:"file"`
	// sds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/appmesh_virtual_gateway#sds AppmeshVirtualGateway#sds}
	Sds *AppmeshVirtualGatewaySpecListenerTlsValidationTrustSds `field:"optional" json:"sds" yaml:"sds"`
}

