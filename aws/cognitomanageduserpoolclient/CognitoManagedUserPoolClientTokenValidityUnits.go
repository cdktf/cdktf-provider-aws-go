package cognitomanageduserpoolclient


type CognitoManagedUserPoolClientTokenValidityUnits struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/cognito_managed_user_pool_client#access_token CognitoManagedUserPoolClient#access_token}.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/cognito_managed_user_pool_client#id_token CognitoManagedUserPoolClient#id_token}.
	IdToken *string `field:"optional" json:"idToken" yaml:"idToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/cognito_managed_user_pool_client#refresh_token CognitoManagedUserPoolClient#refresh_token}.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

