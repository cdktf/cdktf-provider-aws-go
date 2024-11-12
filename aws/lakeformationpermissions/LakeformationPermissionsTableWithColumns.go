// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lakeformationpermissions


type LakeformationPermissionsTableWithColumns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/lakeformation_permissions#database_name LakeformationPermissions#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/lakeformation_permissions#name LakeformationPermissions#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/lakeformation_permissions#catalog_id LakeformationPermissions#catalog_id}.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/lakeformation_permissions#column_names LakeformationPermissions#column_names}.
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/lakeformation_permissions#excluded_column_names LakeformationPermissions#excluded_column_names}.
	ExcludedColumnNames *[]*string `field:"optional" json:"excludedColumnNames" yaml:"excludedColumnNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/lakeformation_permissions#wildcard LakeformationPermissions#wildcard}.
	Wildcard interface{} `field:"optional" json:"wildcard" yaml:"wildcard"`
}

