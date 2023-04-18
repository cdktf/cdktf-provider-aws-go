package cognitouserpool


type CognitoUserPoolAdminCreateUserConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/cognito_user_pool#allow_admin_create_user_only CognitoUserPool#allow_admin_create_user_only}.
	AllowAdminCreateUserOnly interface{} `field:"optional" json:"allowAdminCreateUserOnly" yaml:"allowAdminCreateUserOnly"`
	// invite_message_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/cognito_user_pool#invite_message_template CognitoUserPool#invite_message_template}
	InviteMessageTemplate *CognitoUserPoolAdminCreateUserConfigInviteMessageTemplate `field:"optional" json:"inviteMessageTemplate" yaml:"inviteMessageTemplate"`
}

