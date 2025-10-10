// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncdatasource


type AppsyncDatasourceElasticsearchConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appsync_datasource#endpoint AppsyncDatasource#endpoint}.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appsync_datasource#region AppsyncDatasource#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

