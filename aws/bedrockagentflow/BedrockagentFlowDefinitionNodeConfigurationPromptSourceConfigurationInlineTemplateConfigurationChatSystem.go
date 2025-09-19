// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinitionNodeConfigurationPromptSourceConfigurationInlineTemplateConfigurationChatSystem struct {
	// cache_point block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/bedrockagent_flow#cache_point BedrockagentFlow#cache_point}
	CachePoint interface{} `field:"optional" json:"cachePoint" yaml:"cachePoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/bedrockagent_flow#text BedrockagentFlow#text}.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

