// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type KendraDataSourceConfigurationWebCrawlerConfigurationAuthenticationConfigurationBasicAuthentication struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_data_source#credentials KendraDataSource#credentials}.
	Credentials *string `field:"required" json:"credentials" yaml:"credentials"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_data_source#host KendraDataSource#host}.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/kendra_data_source#port KendraDataSource#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

