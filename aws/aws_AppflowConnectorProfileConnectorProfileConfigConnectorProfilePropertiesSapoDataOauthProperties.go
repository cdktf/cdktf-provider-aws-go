// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSapoDataOauthProperties struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#auth_code_url AppflowConnectorProfile#auth_code_url}.
	AuthCodeUrl *string `field:"required" json:"authCodeUrl" yaml:"authCodeUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#oauth_scopes AppflowConnectorProfile#oauth_scopes}.
	OauthScopes *[]*string `field:"required" json:"oauthScopes" yaml:"oauthScopes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#token_url AppflowConnectorProfile#token_url}.
	TokenUrl *string `field:"required" json:"tokenUrl" yaml:"tokenUrl"`
}

