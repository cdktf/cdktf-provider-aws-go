// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualnode


type AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTls struct {
	// validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/appmesh_virtual_node#validation AppmeshVirtualNode#validation}
	Validation *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidation `field:"required" json:"validation" yaml:"validation"`
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/appmesh_virtual_node#certificate AppmeshVirtualNode#certificate}
	Certificate *AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/appmesh_virtual_node#enforce AppmeshVirtualNode#enforce}.
	Enforce interface{} `field:"optional" json:"enforce" yaml:"enforce"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/appmesh_virtual_node#ports AppmeshVirtualNode#ports}.
	Ports *[]*float64 `field:"optional" json:"ports" yaml:"ports"`
}

