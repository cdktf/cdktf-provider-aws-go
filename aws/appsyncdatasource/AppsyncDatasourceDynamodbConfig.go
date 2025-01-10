// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncdatasource


type AppsyncDatasourceDynamodbConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/appsync_datasource#table_name AppsyncDatasource#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// delta_sync_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/appsync_datasource#delta_sync_config AppsyncDatasource#delta_sync_config}
	DeltaSyncConfig *AppsyncDatasourceDynamodbConfigDeltaSyncConfig `field:"optional" json:"deltaSyncConfig" yaml:"deltaSyncConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/appsync_datasource#region AppsyncDatasource#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/appsync_datasource#use_caller_credentials AppsyncDatasource#use_caller_credentials}.
	UseCallerCredentials interface{} `field:"optional" json:"useCallerCredentials" yaml:"useCallerCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/appsync_datasource#versioned AppsyncDatasource#versioned}.
	Versioned interface{} `field:"optional" json:"versioned" yaml:"versioned"`
}

