// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rolesanywheretrustanchor


type RolesanywhereTrustAnchorSource struct {
	// source_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/rolesanywhere_trust_anchor#source_data RolesanywhereTrustAnchor#source_data}
	SourceData *RolesanywhereTrustAnchorSourceSourceData `field:"required" json:"sourceData" yaml:"sourceData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/rolesanywhere_trust_anchor#source_type RolesanywhereTrustAnchor#source_type}.
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
}

