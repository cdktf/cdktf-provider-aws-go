// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSalesforce struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/appflow_connector_profile#access_token AppflowConnectorProfile#access_token}.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/appflow_connector_profile#client_credentials_arn AppflowConnectorProfile#client_credentials_arn}.
	ClientCredentialsArn *string `field:"optional" json:"clientCredentialsArn" yaml:"clientCredentialsArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/appflow_connector_profile#jwt_token AppflowConnectorProfile#jwt_token}.
	JwtToken *string `field:"optional" json:"jwtToken" yaml:"jwtToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/appflow_connector_profile#oauth2_grant_type AppflowConnectorProfile#oauth2_grant_type}.
	Oauth2GrantType *string `field:"optional" json:"oauth2GrantType" yaml:"oauth2GrantType"`
	// oauth_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/appflow_connector_profile#oauth_request AppflowConnectorProfile#oauth_request}
	OauthRequest *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSalesforceOauthRequest `field:"optional" json:"oauthRequest" yaml:"oauthRequest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/appflow_connector_profile#refresh_token AppflowConnectorProfile#refresh_token}.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

