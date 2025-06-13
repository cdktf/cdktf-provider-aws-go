// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentprompt


type BedrockagentPromptVariantTemplateConfigurationChatToolConfigurationTool struct {
	// cache_point block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/bedrockagent_prompt#cache_point BedrockagentPrompt#cache_point}
	CachePoint interface{} `field:"optional" json:"cachePoint" yaml:"cachePoint"`
	// tool_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/bedrockagent_prompt#tool_spec BedrockagentPrompt#tool_spec}
	ToolSpec interface{} `field:"optional" json:"toolSpec" yaml:"toolSpec"`
}

