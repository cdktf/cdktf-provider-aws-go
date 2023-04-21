package kendradatasource


type KendraDataSourceConfigurationWebCrawlerConfiguration struct {
	// urls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/kendra_data_source#urls KendraDataSource#urls}
	Urls *KendraDataSourceConfigurationWebCrawlerConfigurationUrls `field:"required" json:"urls" yaml:"urls"`
	// authentication_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/kendra_data_source#authentication_configuration KendraDataSource#authentication_configuration}
	AuthenticationConfiguration *KendraDataSourceConfigurationWebCrawlerConfigurationAuthenticationConfiguration `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/kendra_data_source#crawl_depth KendraDataSource#crawl_depth}.
	CrawlDepth *float64 `field:"optional" json:"crawlDepth" yaml:"crawlDepth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/kendra_data_source#max_content_size_per_page_in_mega_bytes KendraDataSource#max_content_size_per_page_in_mega_bytes}.
	MaxContentSizePerPageInMegaBytes *float64 `field:"optional" json:"maxContentSizePerPageInMegaBytes" yaml:"maxContentSizePerPageInMegaBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/kendra_data_source#max_links_per_page KendraDataSource#max_links_per_page}.
	MaxLinksPerPage *float64 `field:"optional" json:"maxLinksPerPage" yaml:"maxLinksPerPage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/kendra_data_source#max_urls_per_minute_crawl_rate KendraDataSource#max_urls_per_minute_crawl_rate}.
	MaxUrlsPerMinuteCrawlRate *float64 `field:"optional" json:"maxUrlsPerMinuteCrawlRate" yaml:"maxUrlsPerMinuteCrawlRate"`
	// proxy_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/kendra_data_source#proxy_configuration KendraDataSource#proxy_configuration}
	ProxyConfiguration *KendraDataSourceConfigurationWebCrawlerConfigurationProxyConfiguration `field:"optional" json:"proxyConfiguration" yaml:"proxyConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/kendra_data_source#url_exclusion_patterns KendraDataSource#url_exclusion_patterns}.
	UrlExclusionPatterns *[]*string `field:"optional" json:"urlExclusionPatterns" yaml:"urlExclusionPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/kendra_data_source#url_inclusion_patterns KendraDataSource#url_inclusion_patterns}.
	UrlInclusionPatterns *[]*string `field:"optional" json:"urlInclusionPatterns" yaml:"urlInclusionPatterns"`
}

