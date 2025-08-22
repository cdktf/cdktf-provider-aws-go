// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceParametersSnowflake struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/quicksight_data_source#database QuicksightDataSource#database}.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/quicksight_data_source#host QuicksightDataSource#host}.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/quicksight_data_source#warehouse QuicksightDataSource#warehouse}.
	Warehouse *string `field:"required" json:"warehouse" yaml:"warehouse"`
}

