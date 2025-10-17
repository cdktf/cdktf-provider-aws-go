// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentprompt


type BedrockagentPromptVariantTemplateConfigurationChatSystem struct {
	// cache_point block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/bedrockagent_prompt#cache_point BedrockagentPrompt#cache_point}
	CachePoint interface{} `field:"optional" json:"cachePoint" yaml:"cachePoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/bedrockagent_prompt#text BedrockagentPrompt#text}.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

