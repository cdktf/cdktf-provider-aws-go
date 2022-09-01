package cognito


type CognitoUserPoolPasswordPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool#minimum_length CognitoUserPool#minimum_length}.
	MinimumLength *float64 `field:"optional" json:"minimumLength" yaml:"minimumLength"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool#require_lowercase CognitoUserPool#require_lowercase}.
	RequireLowercase interface{} `field:"optional" json:"requireLowercase" yaml:"requireLowercase"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool#require_numbers CognitoUserPool#require_numbers}.
	RequireNumbers interface{} `field:"optional" json:"requireNumbers" yaml:"requireNumbers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool#require_symbols CognitoUserPool#require_symbols}.
	RequireSymbols interface{} `field:"optional" json:"requireSymbols" yaml:"requireSymbols"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool#require_uppercase CognitoUserPool#require_uppercase}.
	RequireUppercase interface{} `field:"optional" json:"requireUppercase" yaml:"requireUppercase"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool#temporary_password_validity_days CognitoUserPool#temporary_password_validity_days}.
	TemporaryPasswordValidityDays *float64 `field:"optional" json:"temporaryPasswordValidityDays" yaml:"temporaryPasswordValidityDays"`
}

