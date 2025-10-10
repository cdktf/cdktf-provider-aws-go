// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoData struct {
	// basic_auth_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appflow_connector_profile#basic_auth_credentials AppflowConnectorProfile#basic_auth_credentials}
	BasicAuthCredentials *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoDataBasicAuthCredentials `field:"optional" json:"basicAuthCredentials" yaml:"basicAuthCredentials"`
	// oauth_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/appflow_connector_profile#oauth_credentials AppflowConnectorProfile#oauth_credentials}
	OauthCredentials *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoDataOauthCredentials `field:"optional" json:"oauthCredentials" yaml:"oauthCredentials"`
}

