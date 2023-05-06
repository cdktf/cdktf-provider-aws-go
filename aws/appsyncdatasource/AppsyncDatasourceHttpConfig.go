package appsyncdatasource


type AppsyncDatasourceHttpConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/appsync_datasource#endpoint AppsyncDatasource#endpoint}.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// authorization_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/appsync_datasource#authorization_config AppsyncDatasource#authorization_config}
	AuthorizationConfig *AppsyncDatasourceHttpConfigAuthorizationConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
}

