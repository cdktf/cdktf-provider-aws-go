// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawslakeformationpermissions


type DataAwsLakeformationPermissionsDataCellsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/data-sources/lakeformation_permissions#database_name DataAwsLakeformationPermissions#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/data-sources/lakeformation_permissions#name DataAwsLakeformationPermissions#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/data-sources/lakeformation_permissions#table_catalog_id DataAwsLakeformationPermissions#table_catalog_id}.
	TableCatalogId *string `field:"required" json:"tableCatalogId" yaml:"tableCatalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/data-sources/lakeformation_permissions#table_name DataAwsLakeformationPermissions#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

