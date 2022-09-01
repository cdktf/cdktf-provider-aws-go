package cognito


type CognitoUserPoolUsernameConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool#case_sensitive CognitoUserPool#case_sensitive}.
	CaseSensitive interface{} `field:"required" json:"caseSensitive" yaml:"caseSensitive"`
}

