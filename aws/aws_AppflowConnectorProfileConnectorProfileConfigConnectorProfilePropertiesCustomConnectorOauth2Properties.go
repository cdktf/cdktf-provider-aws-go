// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type AppflowConnectorProfileConnectorProfileConfigConnectorProfilePropertiesCustomConnectorOauth2Properties struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#oauth2_grant_type AppflowConnectorProfile#oauth2_grant_type}.
	Oauth2GrantType *string `field:"required" json:"oauth2GrantType" yaml:"oauth2GrantType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#token_url AppflowConnectorProfile#token_url}.
	TokenUrl *string `field:"required" json:"tokenUrl" yaml:"tokenUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#token_url_custom_properties AppflowConnectorProfile#token_url_custom_properties}.
	TokenUrlCustomProperties *map[string]*string `field:"optional" json:"tokenUrlCustomProperties" yaml:"tokenUrlCustomProperties"`
}

