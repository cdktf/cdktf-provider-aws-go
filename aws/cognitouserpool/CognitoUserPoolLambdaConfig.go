// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitouserpool


type CognitoUserPoolLambdaConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/cognito_user_pool#create_auth_challenge CognitoUserPool#create_auth_challenge}.
	CreateAuthChallenge *string `field:"optional" json:"createAuthChallenge" yaml:"createAuthChallenge"`
	// custom_email_sender block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/cognito_user_pool#custom_email_sender CognitoUserPool#custom_email_sender}
	CustomEmailSender *CognitoUserPoolLambdaConfigCustomEmailSender `field:"optional" json:"customEmailSender" yaml:"customEmailSender"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/cognito_user_pool#custom_message CognitoUserPool#custom_message}.
	CustomMessage *string `field:"optional" json:"customMessage" yaml:"customMessage"`
	// custom_sms_sender block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/cognito_user_pool#custom_sms_sender CognitoUserPool#custom_sms_sender}
	CustomSmsSender *CognitoUserPoolLambdaConfigCustomSmsSender `field:"optional" json:"customSmsSender" yaml:"customSmsSender"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/cognito_user_pool#define_auth_challenge CognitoUserPool#define_auth_challenge}.
	DefineAuthChallenge *string `field:"optional" json:"defineAuthChallenge" yaml:"defineAuthChallenge"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/cognito_user_pool#kms_key_id CognitoUserPool#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/cognito_user_pool#post_authentication CognitoUserPool#post_authentication}.
	PostAuthentication *string `field:"optional" json:"postAuthentication" yaml:"postAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/cognito_user_pool#post_confirmation CognitoUserPool#post_confirmation}.
	PostConfirmation *string `field:"optional" json:"postConfirmation" yaml:"postConfirmation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/cognito_user_pool#pre_authentication CognitoUserPool#pre_authentication}.
	PreAuthentication *string `field:"optional" json:"preAuthentication" yaml:"preAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/cognito_user_pool#pre_sign_up CognitoUserPool#pre_sign_up}.
	PreSignUp *string `field:"optional" json:"preSignUp" yaml:"preSignUp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/cognito_user_pool#pre_token_generation CognitoUserPool#pre_token_generation}.
	PreTokenGeneration *string `field:"optional" json:"preTokenGeneration" yaml:"preTokenGeneration"`
	// pre_token_generation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/cognito_user_pool#pre_token_generation_config CognitoUserPool#pre_token_generation_config}
	PreTokenGenerationConfig *CognitoUserPoolLambdaConfigPreTokenGenerationConfig `field:"optional" json:"preTokenGenerationConfig" yaml:"preTokenGenerationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/cognito_user_pool#user_migration CognitoUserPool#user_migration}.
	UserMigration *string `field:"optional" json:"userMigration" yaml:"userMigration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/cognito_user_pool#verify_auth_challenge_response CognitoUserPool#verify_auth_challenge_response}.
	VerifyAuthChallengeResponse *string `field:"optional" json:"verifyAuthChallengeResponse" yaml:"verifyAuthChallengeResponse"`
}

