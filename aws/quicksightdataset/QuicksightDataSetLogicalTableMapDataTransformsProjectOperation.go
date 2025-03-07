// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetLogicalTableMapDataTransformsProjectOperation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/quicksight_data_set#projected_columns QuicksightDataSet#projected_columns}.
	ProjectedColumns *[]*string `field:"required" json:"projectedColumns" yaml:"projectedColumns"`
}

