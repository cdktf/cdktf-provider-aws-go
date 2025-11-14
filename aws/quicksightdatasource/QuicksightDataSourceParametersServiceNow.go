// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceParametersServiceNow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/quicksight_data_source#site_base_url QuicksightDataSource#site_base_url}.
	SiteBaseUrl *string `field:"required" json:"siteBaseUrl" yaml:"siteBaseUrl"`
}

