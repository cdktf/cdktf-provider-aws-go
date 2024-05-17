// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceParametersAmazonElasticsearch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.50.0/docs/resources/quicksight_data_source#domain QuicksightDataSource#domain}.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
}

