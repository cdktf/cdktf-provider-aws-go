// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitoidentitypool


type CognitoIdentityPoolCognitoIdentityProviders struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/cognito_identity_pool#client_id CognitoIdentityPool#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/cognito_identity_pool#provider_name CognitoIdentityPool#provider_name}.
	ProviderName *string `field:"optional" json:"providerName" yaml:"providerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/cognito_identity_pool#server_side_token_check CognitoIdentityPool#server_side_token_check}.
	ServerSideTokenCheck interface{} `field:"optional" json:"serverSideTokenCheck" yaml:"serverSideTokenCheck"`
}

