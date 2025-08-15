// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentprompt


type BedrockagentPromptVariantTemplateConfigurationChat struct {
	// input_variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/bedrockagent_prompt#input_variable BedrockagentPrompt#input_variable}
	InputVariable interface{} `field:"optional" json:"inputVariable" yaml:"inputVariable"`
	// message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/bedrockagent_prompt#message BedrockagentPrompt#message}
	Message interface{} `field:"optional" json:"message" yaml:"message"`
	// system block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/bedrockagent_prompt#system BedrockagentPrompt#system}
	SystemAttribute interface{} `field:"optional" json:"systemAttribute" yaml:"systemAttribute"`
	// tool_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/bedrockagent_prompt#tool_configuration BedrockagentPrompt#tool_configuration}
	ToolConfiguration interface{} `field:"optional" json:"toolConfiguration" yaml:"toolConfiguration"`
}

