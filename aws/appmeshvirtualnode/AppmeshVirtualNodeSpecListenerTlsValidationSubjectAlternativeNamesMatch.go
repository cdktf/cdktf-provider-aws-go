// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualnode


type AppmeshVirtualNodeSpecListenerTlsValidationSubjectAlternativeNamesMatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/appmesh_virtual_node#exact AppmeshVirtualNode#exact}.
	Exact *[]*string `field:"required" json:"exact" yaml:"exact"`
}

