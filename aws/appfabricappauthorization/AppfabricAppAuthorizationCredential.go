// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appfabricappauthorization


type AppfabricAppAuthorizationCredential struct {
	// api_key_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/appfabric_app_authorization#api_key_credential AppfabricAppAuthorization#api_key_credential}
	ApiKeyCredential interface{} `field:"optional" json:"apiKeyCredential" yaml:"apiKeyCredential"`
	// oauth2_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/appfabric_app_authorization#oauth2_credential AppfabricAppAuthorization#oauth2_credential}
	Oauth2Credential interface{} `field:"optional" json:"oauth2Credential" yaml:"oauth2Credential"`
}

