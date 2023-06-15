package cognitouserpool


type CognitoUserPoolEmailConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/cognito_user_pool#configuration_set CognitoUserPool#configuration_set}.
	ConfigurationSet *string `field:"optional" json:"configurationSet" yaml:"configurationSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/cognito_user_pool#email_sending_account CognitoUserPool#email_sending_account}.
	EmailSendingAccount *string `field:"optional" json:"emailSendingAccount" yaml:"emailSendingAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/cognito_user_pool#from_email_address CognitoUserPool#from_email_address}.
	FromEmailAddress *string `field:"optional" json:"fromEmailAddress" yaml:"fromEmailAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/cognito_user_pool#reply_to_email_address CognitoUserPool#reply_to_email_address}.
	ReplyToEmailAddress *string `field:"optional" json:"replyToEmailAddress" yaml:"replyToEmailAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/resources/cognito_user_pool#source_arn CognitoUserPool#source_arn}.
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
}

