// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinitionNodeConfigurationPromptSourceConfigurationInlineTemplateConfigurationChatToolConfiguration struct {
	// tool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/bedrockagent_flow#tool BedrockagentFlow#tool}
	Tool interface{} `field:"optional" json:"tool" yaml:"tool"`
	// tool_choice block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/bedrockagent_flow#tool_choice BedrockagentFlow#tool_choice}
	ToolChoice interface{} `field:"optional" json:"toolChoice" yaml:"toolChoice"`
}

