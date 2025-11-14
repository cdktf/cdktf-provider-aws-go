// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kendradatasource


type KendraDataSourceConfigurationWebCrawlerConfigurationUrls struct {
	// seed_url_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/kendra_data_source#seed_url_configuration KendraDataSource#seed_url_configuration}
	SeedUrlConfiguration *KendraDataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfiguration `field:"optional" json:"seedUrlConfiguration" yaml:"seedUrlConfiguration"`
	// site_maps_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/kendra_data_source#site_maps_configuration KendraDataSource#site_maps_configuration}
	SiteMapsConfiguration *KendraDataSourceConfigurationWebCrawlerConfigurationUrlsSiteMapsConfiguration `field:"optional" json:"siteMapsConfiguration" yaml:"siteMapsConfiguration"`
}

