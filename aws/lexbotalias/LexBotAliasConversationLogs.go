// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexbotalias


type LexBotAliasConversationLogs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/lex_bot_alias#iam_role_arn LexBotAlias#iam_role_arn}.
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// log_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/lex_bot_alias#log_settings LexBotAlias#log_settings}
	LogSettings interface{} `field:"optional" json:"logSettings" yaml:"logSettings"`
}

