package kendradatasource


type KendraDataSourceConfigurationWebCrawlerConfigurationUrls struct {
	// seed_url_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_data_source#seed_url_configuration KendraDataSource#seed_url_configuration}
	SeedUrlConfiguration *KendraDataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfiguration `field:"optional" json:"seedUrlConfiguration" yaml:"seedUrlConfiguration"`
	// site_maps_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_data_source#site_maps_configuration KendraDataSource#site_maps_configuration}
	SiteMapsConfiguration *KendraDataSourceConfigurationWebCrawlerConfigurationUrlsSiteMapsConfiguration `field:"optional" json:"siteMapsConfiguration" yaml:"siteMapsConfiguration"`
}

