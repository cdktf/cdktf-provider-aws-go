package cognitouserpool


type CognitoUserPoolAdminCreateUserConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool#allow_admin_create_user_only CognitoUserPool#allow_admin_create_user_only}.
	AllowAdminCreateUserOnly interface{} `field:"optional" json:"allowAdminCreateUserOnly" yaml:"allowAdminCreateUserOnly"`
	// invite_message_template block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool#invite_message_template CognitoUserPool#invite_message_template}
	InviteMessageTemplate *CognitoUserPoolAdminCreateUserConfigInviteMessageTemplate `field:"optional" json:"inviteMessageTemplate" yaml:"inviteMessageTemplate"`
}

