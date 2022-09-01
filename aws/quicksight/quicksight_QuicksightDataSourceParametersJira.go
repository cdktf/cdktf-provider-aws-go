package quicksight


type QuicksightDataSourceParametersJira struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source#site_base_url QuicksightDataSource#site_base_url}.
	SiteBaseUrl *string `field:"required" json:"siteBaseUrl" yaml:"siteBaseUrl"`
}

