// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceParametersRedshift struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/quicksight_data_source#database QuicksightDataSource#database}.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/quicksight_data_source#cluster_id QuicksightDataSource#cluster_id}.
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/quicksight_data_source#host QuicksightDataSource#host}.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/quicksight_data_source#port QuicksightDataSource#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

