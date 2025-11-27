// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinitionNodeConfigurationPromptSourceConfigurationInlineTemplateConfiguration struct {
	// chat block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/bedrockagent_flow#chat BedrockagentFlow#chat}
	Chat interface{} `field:"optional" json:"chat" yaml:"chat"`
	// text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/bedrockagent_flow#text BedrockagentFlow#text}
	Text interface{} `field:"optional" json:"text" yaml:"text"`
}

