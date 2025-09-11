// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinitionNodeConfigurationLex struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/bedrockagent_flow#bot_alias_arn BedrockagentFlow#bot_alias_arn}.
	BotAliasArn *string `field:"required" json:"botAliasArn" yaml:"botAliasArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/bedrockagent_flow#locale_id BedrockagentFlow#locale_id}.
	LocaleId *string `field:"required" json:"localeId" yaml:"localeId"`
}

