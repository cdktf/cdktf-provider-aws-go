package cognito


type CognitoUserPoolAdminCreateUserConfigInviteMessageTemplate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool#email_message CognitoUserPool#email_message}.
	EmailMessage *string `field:"optional" json:"emailMessage" yaml:"emailMessage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool#email_subject CognitoUserPool#email_subject}.
	EmailSubject *string `field:"optional" json:"emailSubject" yaml:"emailSubject"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool#sms_message CognitoUserPool#sms_message}.
	SmsMessage *string `field:"optional" json:"smsMessage" yaml:"smsMessage"`
}

