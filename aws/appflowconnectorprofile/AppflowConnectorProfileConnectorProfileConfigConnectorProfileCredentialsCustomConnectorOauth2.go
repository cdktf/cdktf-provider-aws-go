// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/appflow_connector_profile#access_token AppflowConnectorProfile#access_token}.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/appflow_connector_profile#client_id AppflowConnectorProfile#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/appflow_connector_profile#client_secret AppflowConnectorProfile#client_secret}.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// oauth_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/appflow_connector_profile#oauth_request AppflowConnectorProfile#oauth_request}
	OauthRequest *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorOauth2OauthRequest `field:"optional" json:"oauthRequest" yaml:"oauthRequest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/appflow_connector_profile#refresh_token AppflowConnectorProfile#refresh_token}.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

