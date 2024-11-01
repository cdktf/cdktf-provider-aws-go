// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cleanroomsconfiguredtable


type CleanroomsConfiguredTableTableReference struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/cleanrooms_configured_table#database_name CleanroomsConfiguredTable#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/cleanrooms_configured_table#table_name CleanroomsConfiguredTable#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

