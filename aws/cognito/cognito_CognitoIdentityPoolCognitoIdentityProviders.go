package cognito


type CognitoIdentityPoolCognitoIdentityProviders struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool#client_id CognitoIdentityPool#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool#provider_name CognitoIdentityPool#provider_name}.
	ProviderName *string `field:"optional" json:"providerName" yaml:"providerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_identity_pool#server_side_token_check CognitoIdentityPool#server_side_token_check}.
	ServerSideTokenCheck interface{} `field:"optional" json:"serverSideTokenCheck" yaml:"serverSideTokenCheck"`
}

