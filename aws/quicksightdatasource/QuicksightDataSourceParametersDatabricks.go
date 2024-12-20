// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceParametersDatabricks struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/quicksight_data_source#host QuicksightDataSource#host}.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/quicksight_data_source#port QuicksightDataSource#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/quicksight_data_source#sql_endpoint_path QuicksightDataSource#sql_endpoint_path}.
	SqlEndpointPath *string `field:"required" json:"sqlEndpointPath" yaml:"sqlEndpointPath"`
}

