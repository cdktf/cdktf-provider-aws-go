// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetLogicalTableMapDataTransformsCreateColumnsOperation struct {
	// columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.17.0/docs/resources/quicksight_data_set#columns QuicksightDataSet#columns}
	Columns interface{} `field:"required" json:"columns" yaml:"columns"`
}

