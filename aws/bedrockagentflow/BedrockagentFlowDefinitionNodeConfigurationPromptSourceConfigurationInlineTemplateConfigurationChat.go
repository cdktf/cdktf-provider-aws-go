// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinitionNodeConfigurationPromptSourceConfigurationInlineTemplateConfigurationChat struct {
	// input_variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/bedrockagent_flow#input_variable BedrockagentFlow#input_variable}
	InputVariable interface{} `field:"optional" json:"inputVariable" yaml:"inputVariable"`
	// message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/bedrockagent_flow#message BedrockagentFlow#message}
	Message interface{} `field:"optional" json:"message" yaml:"message"`
	// system block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/bedrockagent_flow#system BedrockagentFlow#system}
	SystemAttribute interface{} `field:"optional" json:"systemAttribute" yaml:"systemAttribute"`
	// tool_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/bedrockagent_flow#tool_configuration BedrockagentFlow#tool_configuration}
	ToolConfiguration interface{} `field:"optional" json:"toolConfiguration" yaml:"toolConfiguration"`
}

