// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitouserpool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CognitoUserPoolConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#name CognitoUserPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// account_recovery_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#account_recovery_setting CognitoUserPool#account_recovery_setting}
	AccountRecoverySetting *CognitoUserPoolAccountRecoverySetting `field:"optional" json:"accountRecoverySetting" yaml:"accountRecoverySetting"`
	// admin_create_user_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#admin_create_user_config CognitoUserPool#admin_create_user_config}
	AdminCreateUserConfig *CognitoUserPoolAdminCreateUserConfig `field:"optional" json:"adminCreateUserConfig" yaml:"adminCreateUserConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#alias_attributes CognitoUserPool#alias_attributes}.
	AliasAttributes *[]*string `field:"optional" json:"aliasAttributes" yaml:"aliasAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#auto_verified_attributes CognitoUserPool#auto_verified_attributes}.
	AutoVerifiedAttributes *[]*string `field:"optional" json:"autoVerifiedAttributes" yaml:"autoVerifiedAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#deletion_protection CognitoUserPool#deletion_protection}.
	DeletionProtection *string `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// device_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#device_configuration CognitoUserPool#device_configuration}
	DeviceConfiguration *CognitoUserPoolDeviceConfiguration `field:"optional" json:"deviceConfiguration" yaml:"deviceConfiguration"`
	// email_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#email_configuration CognitoUserPool#email_configuration}
	EmailConfiguration *CognitoUserPoolEmailConfiguration `field:"optional" json:"emailConfiguration" yaml:"emailConfiguration"`
	// email_mfa_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#email_mfa_configuration CognitoUserPool#email_mfa_configuration}
	EmailMfaConfiguration *CognitoUserPoolEmailMfaConfiguration `field:"optional" json:"emailMfaConfiguration" yaml:"emailMfaConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#email_verification_message CognitoUserPool#email_verification_message}.
	EmailVerificationMessage *string `field:"optional" json:"emailVerificationMessage" yaml:"emailVerificationMessage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#email_verification_subject CognitoUserPool#email_verification_subject}.
	EmailVerificationSubject *string `field:"optional" json:"emailVerificationSubject" yaml:"emailVerificationSubject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#id CognitoUserPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// lambda_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#lambda_config CognitoUserPool#lambda_config}
	LambdaConfig *CognitoUserPoolLambdaConfig `field:"optional" json:"lambdaConfig" yaml:"lambdaConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#mfa_configuration CognitoUserPool#mfa_configuration}.
	MfaConfiguration *string `field:"optional" json:"mfaConfiguration" yaml:"mfaConfiguration"`
	// password_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#password_policy CognitoUserPool#password_policy}
	PasswordPolicy *CognitoUserPoolPasswordPolicy `field:"optional" json:"passwordPolicy" yaml:"passwordPolicy"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#schema CognitoUserPool#schema}
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
	// sign_in_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#sign_in_policy CognitoUserPool#sign_in_policy}
	SignInPolicy *CognitoUserPoolSignInPolicy `field:"optional" json:"signInPolicy" yaml:"signInPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#sms_authentication_message CognitoUserPool#sms_authentication_message}.
	SmsAuthenticationMessage *string `field:"optional" json:"smsAuthenticationMessage" yaml:"smsAuthenticationMessage"`
	// sms_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#sms_configuration CognitoUserPool#sms_configuration}
	SmsConfiguration *CognitoUserPoolSmsConfiguration `field:"optional" json:"smsConfiguration" yaml:"smsConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#sms_verification_message CognitoUserPool#sms_verification_message}.
	SmsVerificationMessage *string `field:"optional" json:"smsVerificationMessage" yaml:"smsVerificationMessage"`
	// software_token_mfa_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#software_token_mfa_configuration CognitoUserPool#software_token_mfa_configuration}
	SoftwareTokenMfaConfiguration *CognitoUserPoolSoftwareTokenMfaConfiguration `field:"optional" json:"softwareTokenMfaConfiguration" yaml:"softwareTokenMfaConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#tags CognitoUserPool#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#tags_all CognitoUserPool#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// user_attribute_update_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#user_attribute_update_settings CognitoUserPool#user_attribute_update_settings}
	UserAttributeUpdateSettings *CognitoUserPoolUserAttributeUpdateSettings `field:"optional" json:"userAttributeUpdateSettings" yaml:"userAttributeUpdateSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#username_attributes CognitoUserPool#username_attributes}.
	UsernameAttributes *[]*string `field:"optional" json:"usernameAttributes" yaml:"usernameAttributes"`
	// username_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#username_configuration CognitoUserPool#username_configuration}
	UsernameConfiguration *CognitoUserPoolUsernameConfiguration `field:"optional" json:"usernameConfiguration" yaml:"usernameConfiguration"`
	// user_pool_add_ons block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#user_pool_add_ons CognitoUserPool#user_pool_add_ons}
	UserPoolAddOns *CognitoUserPoolUserPoolAddOns `field:"optional" json:"userPoolAddOns" yaml:"userPoolAddOns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#user_pool_tier CognitoUserPool#user_pool_tier}.
	UserPoolTier *string `field:"optional" json:"userPoolTier" yaml:"userPoolTier"`
	// verification_message_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#verification_message_template CognitoUserPool#verification_message_template}
	VerificationMessageTemplate *CognitoUserPoolVerificationMessageTemplate `field:"optional" json:"verificationMessageTemplate" yaml:"verificationMessageTemplate"`
	// web_authn_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/cognito_user_pool#web_authn_configuration CognitoUserPool#web_authn_configuration}
	WebAuthnConfiguration *CognitoUserPoolWebAuthnConfiguration `field:"optional" json:"webAuthnConfiguration" yaml:"webAuthnConfiguration"`
}

