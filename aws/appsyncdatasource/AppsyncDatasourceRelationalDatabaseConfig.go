// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncdatasource


type AppsyncDatasourceRelationalDatabaseConfig struct {
	// http_endpoint_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/appsync_datasource#http_endpoint_config AppsyncDatasource#http_endpoint_config}
	HttpEndpointConfig *AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfig `field:"optional" json:"httpEndpointConfig" yaml:"httpEndpointConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/appsync_datasource#source_type AppsyncDatasource#source_type}.
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
}

