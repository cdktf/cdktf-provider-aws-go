// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentprompt


type BedrockagentPromptVariantTemplateConfigurationChatMessage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/bedrockagent_prompt#role BedrockagentPrompt#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// content block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/bedrockagent_prompt#content BedrockagentPrompt#content}
	Content interface{} `field:"optional" json:"content" yaml:"content"`
}

