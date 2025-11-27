// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceParametersSpark struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/quicksight_data_source#host QuicksightDataSource#host}.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/quicksight_data_source#port QuicksightDataSource#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

