// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetLogicalTableMapDataTransformsCastColumnTypeOperation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/quicksight_data_set#column_name QuicksightDataSet#column_name}.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/quicksight_data_set#new_column_type QuicksightDataSet#new_column_type}.
	NewColumnType *string `field:"required" json:"newColumnType" yaml:"newColumnType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/quicksight_data_set#format QuicksightDataSet#format}.
	Format *string `field:"optional" json:"format" yaml:"format"`
}

