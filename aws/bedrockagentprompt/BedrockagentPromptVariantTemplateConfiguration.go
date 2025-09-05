// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentprompt


type BedrockagentPromptVariantTemplateConfiguration struct {
	// chat block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/bedrockagent_prompt#chat BedrockagentPrompt#chat}
	Chat interface{} `field:"optional" json:"chat" yaml:"chat"`
	// text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/bedrockagent_prompt#text BedrockagentPrompt#text}
	Text interface{} `field:"optional" json:"text" yaml:"text"`
}

