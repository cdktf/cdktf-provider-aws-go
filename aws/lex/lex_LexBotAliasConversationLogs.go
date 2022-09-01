package lex


type LexBotAliasConversationLogs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot_alias#iam_role_arn LexBotAlias#iam_role_arn}.
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// log_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lex_bot_alias#log_settings LexBotAlias#log_settings}
	LogSettings interface{} `field:"optional" json:"logSettings" yaml:"logSettings"`
}

