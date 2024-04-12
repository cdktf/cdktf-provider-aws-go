// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lakeformationpermissions


type LakeformationPermissionsDatabase struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/lakeformation_permissions#name LakeformationPermissions#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/lakeformation_permissions#catalog_id LakeformationPermissions#catalog_id}.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
}

