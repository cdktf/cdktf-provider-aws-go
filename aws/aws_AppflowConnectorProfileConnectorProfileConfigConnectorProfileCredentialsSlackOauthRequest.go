// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSlackOauthRequest struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#auth_code AppflowConnectorProfile#auth_code}.
	AuthCode *string `field:"optional" json:"authCode" yaml:"authCode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_connector_profile#redirect_uri AppflowConnectorProfile#redirect_uri}.
	RedirectUri *string `field:"optional" json:"redirectUri" yaml:"redirectUri"`
}

