// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncdatasource


type AppsyncDatasourceDynamodbConfigDeltaSyncConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/appsync_datasource#delta_sync_table_name AppsyncDatasource#delta_sync_table_name}.
	DeltaSyncTableName *string `field:"required" json:"deltaSyncTableName" yaml:"deltaSyncTableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/appsync_datasource#base_table_ttl AppsyncDatasource#base_table_ttl}.
	BaseTableTtl *float64 `field:"optional" json:"baseTableTtl" yaml:"baseTableTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/appsync_datasource#delta_sync_table_ttl AppsyncDatasource#delta_sync_table_ttl}.
	DeltaSyncTableTtl *float64 `field:"optional" json:"deltaSyncTableTtl" yaml:"deltaSyncTableTtl"`
}

