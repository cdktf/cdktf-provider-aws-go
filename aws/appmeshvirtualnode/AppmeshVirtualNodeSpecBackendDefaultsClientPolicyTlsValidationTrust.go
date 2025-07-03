// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualnode


type AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrust struct {
	// acm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/appmesh_virtual_node#acm AppmeshVirtualNode#acm}
	Acm *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustAcm `field:"optional" json:"acm" yaml:"acm"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/appmesh_virtual_node#file AppmeshVirtualNode#file}
	File *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustFile `field:"optional" json:"file" yaml:"file"`
	// sds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/appmesh_virtual_node#sds AppmeshVirtualNode#sds}
	Sds *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustSds `field:"optional" json:"sds" yaml:"sds"`
}

