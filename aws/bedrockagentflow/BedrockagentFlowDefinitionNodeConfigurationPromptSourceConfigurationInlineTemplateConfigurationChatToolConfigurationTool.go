// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinitionNodeConfigurationPromptSourceConfigurationInlineTemplateConfigurationChatToolConfigurationTool struct {
	// cache_point block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/bedrockagent_flow#cache_point BedrockagentFlow#cache_point}
	CachePoint interface{} `field:"optional" json:"cachePoint" yaml:"cachePoint"`
	// tool_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/bedrockagent_flow#tool_spec BedrockagentFlow#tool_spec}
	ToolSpec interface{} `field:"optional" json:"toolSpec" yaml:"toolSpec"`
}

