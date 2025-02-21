// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetLogicalTableMap struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/quicksight_data_set#alias QuicksightDataSet#alias}.
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/quicksight_data_set#logical_table_map_id QuicksightDataSet#logical_table_map_id}.
	LogicalTableMapId *string `field:"required" json:"logicalTableMapId" yaml:"logicalTableMapId"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/quicksight_data_set#source QuicksightDataSet#source}
	Source *QuicksightDataSetLogicalTableMapSource `field:"required" json:"source" yaml:"source"`
	// data_transforms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/quicksight_data_set#data_transforms QuicksightDataSet#data_transforms}
	DataTransforms interface{} `field:"optional" json:"dataTransforms" yaml:"dataTransforms"`
}

