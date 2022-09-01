// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type KendraDataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_data_source#seed_urls KendraDataSource#seed_urls}.
	SeedUrls *[]*string `field:"required" json:"seedUrls" yaml:"seedUrls"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_data_source#web_crawler_mode KendraDataSource#web_crawler_mode}.
	WebCrawlerMode *string `field:"optional" json:"webCrawlerMode" yaml:"webCrawlerMode"`
}

