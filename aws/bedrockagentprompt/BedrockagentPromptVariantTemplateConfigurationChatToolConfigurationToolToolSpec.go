// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentprompt


type BedrockagentPromptVariantTemplateConfigurationChatToolConfigurationToolToolSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/bedrockagent_prompt#name BedrockagentPrompt#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/bedrockagent_prompt#description BedrockagentPrompt#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// input_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/bedrockagent_prompt#input_schema BedrockagentPrompt#input_schema}
	InputSchema interface{} `field:"optional" json:"inputSchema" yaml:"inputSchema"`
}

