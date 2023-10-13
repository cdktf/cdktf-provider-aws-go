// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetFieldFolders struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/quicksight_data_set#field_folders_id QuicksightDataSet#field_folders_id}.
	FieldFoldersId *string `field:"required" json:"fieldFoldersId" yaml:"fieldFoldersId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/quicksight_data_set#columns QuicksightDataSet#columns}.
	Columns *[]*string `field:"optional" json:"columns" yaml:"columns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/quicksight_data_set#description QuicksightDataSet#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

