// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentprompt


type BedrockagentPromptVariantTemplateConfigurationChatToolConfiguration struct {
	// tool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/bedrockagent_prompt#tool BedrockagentPrompt#tool}
	Tool interface{} `field:"optional" json:"tool" yaml:"tool"`
	// tool_choice block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/bedrockagent_prompt#tool_choice BedrockagentPrompt#tool_choice}
	ToolChoice interface{} `field:"optional" json:"toolChoice" yaml:"toolChoice"`
}

