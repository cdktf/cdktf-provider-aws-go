// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualnode


type AppmeshVirtualNodeSpecListenerTlsCertificate struct {
	// acm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/appmesh_virtual_node#acm AppmeshVirtualNode#acm}
	Acm *AppmeshVirtualNodeSpecListenerTlsCertificateAcm `field:"optional" json:"acm" yaml:"acm"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/appmesh_virtual_node#file AppmeshVirtualNode#file}
	File *AppmeshVirtualNodeSpecListenerTlsCertificateFile `field:"optional" json:"file" yaml:"file"`
	// sds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/appmesh_virtual_node#sds AppmeshVirtualNode#sds}
	Sds *AppmeshVirtualNodeSpecListenerTlsCertificateSds `field:"optional" json:"sds" yaml:"sds"`
}

