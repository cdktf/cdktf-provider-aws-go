// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceParametersTwitter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/quicksight_data_source#max_rows QuicksightDataSource#max_rows}.
	MaxRows *float64 `field:"required" json:"maxRows" yaml:"maxRows"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/quicksight_data_source#query QuicksightDataSource#query}.
	Query *string `field:"required" json:"query" yaml:"query"`
}

