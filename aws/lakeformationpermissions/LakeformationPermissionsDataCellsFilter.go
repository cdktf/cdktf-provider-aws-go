// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lakeformationpermissions


type LakeformationPermissionsDataCellsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lakeformation_permissions#database_name LakeformationPermissions#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lakeformation_permissions#name LakeformationPermissions#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lakeformation_permissions#table_catalog_id LakeformationPermissions#table_catalog_id}.
	TableCatalogId *string `field:"required" json:"tableCatalogId" yaml:"tableCatalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/lakeformation_permissions#table_name LakeformationPermissions#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

