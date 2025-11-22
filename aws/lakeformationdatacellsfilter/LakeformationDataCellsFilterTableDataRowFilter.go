// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lakeformationdatacellsfilter


type LakeformationDataCellsFilterTableDataRowFilter struct {
	// all_rows_wildcard block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/lakeformation_data_cells_filter#all_rows_wildcard LakeformationDataCellsFilter#all_rows_wildcard}
	AllRowsWildcard interface{} `field:"optional" json:"allRowsWildcard" yaml:"allRowsWildcard"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/lakeformation_data_cells_filter#filter_expression LakeformationDataCellsFilter#filter_expression}.
	FilterExpression *string `field:"optional" json:"filterExpression" yaml:"filterExpression"`
}

