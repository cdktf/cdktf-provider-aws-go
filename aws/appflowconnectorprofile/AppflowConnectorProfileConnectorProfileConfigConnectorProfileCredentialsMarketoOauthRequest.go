package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoOauthRequest struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/appflow_connector_profile#auth_code AppflowConnectorProfile#auth_code}.
	AuthCode *string `field:"optional" json:"authCode" yaml:"authCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/appflow_connector_profile#redirect_uri AppflowConnectorProfile#redirect_uri}.
	RedirectUri *string `field:"optional" json:"redirectUri" yaml:"redirectUri"`
}

