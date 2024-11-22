// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualnode


type AppmeshVirtualNodeSpecListenerTlsCertificateSds struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.77.0/docs/resources/appmesh_virtual_node#secret_name AppmeshVirtualNode#secret_name}.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
}

