// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetLogicalTableMapDataTransforms struct {
	// cast_column_type_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/quicksight_data_set#cast_column_type_operation QuicksightDataSet#cast_column_type_operation}
	CastColumnTypeOperation *QuicksightDataSetLogicalTableMapDataTransformsCastColumnTypeOperation `field:"optional" json:"castColumnTypeOperation" yaml:"castColumnTypeOperation"`
	// create_columns_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/quicksight_data_set#create_columns_operation QuicksightDataSet#create_columns_operation}
	CreateColumnsOperation *QuicksightDataSetLogicalTableMapDataTransformsCreateColumnsOperation `field:"optional" json:"createColumnsOperation" yaml:"createColumnsOperation"`
	// filter_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/quicksight_data_set#filter_operation QuicksightDataSet#filter_operation}
	FilterOperation *QuicksightDataSetLogicalTableMapDataTransformsFilterOperation `field:"optional" json:"filterOperation" yaml:"filterOperation"`
	// project_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/quicksight_data_set#project_operation QuicksightDataSet#project_operation}
	ProjectOperation *QuicksightDataSetLogicalTableMapDataTransformsProjectOperation `field:"optional" json:"projectOperation" yaml:"projectOperation"`
	// rename_column_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/quicksight_data_set#rename_column_operation QuicksightDataSet#rename_column_operation}
	RenameColumnOperation *QuicksightDataSetLogicalTableMapDataTransformsRenameColumnOperation `field:"optional" json:"renameColumnOperation" yaml:"renameColumnOperation"`
	// tag_column_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/quicksight_data_set#tag_column_operation QuicksightDataSet#tag_column_operation}
	TagColumnOperation *QuicksightDataSetLogicalTableMapDataTransformsTagColumnOperation `field:"optional" json:"tagColumnOperation" yaml:"tagColumnOperation"`
	// untag_column_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/quicksight_data_set#untag_column_operation QuicksightDataSet#untag_column_operation}
	UntagColumnOperation *QuicksightDataSetLogicalTableMapDataTransformsUntagColumnOperation `field:"optional" json:"untagColumnOperation" yaml:"untagColumnOperation"`
}

