// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/appflow_connector_profile#client_id AppflowConnectorProfile#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/appflow_connector_profile#client_secret AppflowConnectorProfile#client_secret}.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/appflow_connector_profile#access_token AppflowConnectorProfile#access_token}.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// oauth_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/appflow_connector_profile#oauth_request AppflowConnectorProfile#oauth_request}
	OauthRequest *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsMarketoOauthRequest `field:"optional" json:"oauthRequest" yaml:"oauthRequest"`
}

