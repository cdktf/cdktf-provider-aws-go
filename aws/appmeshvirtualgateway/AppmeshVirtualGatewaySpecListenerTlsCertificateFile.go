// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecListenerTlsCertificateFile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/appmesh_virtual_gateway#certificate_chain AppmeshVirtualGateway#certificate_chain}.
	CertificateChain *string `field:"required" json:"certificateChain" yaml:"certificateChain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/appmesh_virtual_gateway#private_key AppmeshVirtualGateway#private_key}.
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
}

