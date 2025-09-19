// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitouserpool


type CognitoUserPoolPasswordPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/cognito_user_pool#minimum_length CognitoUserPool#minimum_length}.
	MinimumLength *float64 `field:"optional" json:"minimumLength" yaml:"minimumLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/cognito_user_pool#password_history_size CognitoUserPool#password_history_size}.
	PasswordHistorySize *float64 `field:"optional" json:"passwordHistorySize" yaml:"passwordHistorySize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/cognito_user_pool#require_lowercase CognitoUserPool#require_lowercase}.
	RequireLowercase interface{} `field:"optional" json:"requireLowercase" yaml:"requireLowercase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/cognito_user_pool#require_numbers CognitoUserPool#require_numbers}.
	RequireNumbers interface{} `field:"optional" json:"requireNumbers" yaml:"requireNumbers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/cognito_user_pool#require_symbols CognitoUserPool#require_symbols}.
	RequireSymbols interface{} `field:"optional" json:"requireSymbols" yaml:"requireSymbols"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/cognito_user_pool#require_uppercase CognitoUserPool#require_uppercase}.
	RequireUppercase interface{} `field:"optional" json:"requireUppercase" yaml:"requireUppercase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/cognito_user_pool#temporary_password_validity_days CognitoUserPool#temporary_password_validity_days}.
	TemporaryPasswordValidityDays *float64 `field:"optional" json:"temporaryPasswordValidityDays" yaml:"temporaryPasswordValidityDays"`
}

