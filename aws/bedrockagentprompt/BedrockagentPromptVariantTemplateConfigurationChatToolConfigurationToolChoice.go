// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentprompt


type BedrockagentPromptVariantTemplateConfigurationChatToolConfigurationToolChoice struct {
	// any block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/bedrockagent_prompt#any BedrockagentPrompt#any}
	Any interface{} `field:"optional" json:"any" yaml:"any"`
	// auto block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/bedrockagent_prompt#auto BedrockagentPrompt#auto}
	Auto interface{} `field:"optional" json:"auto" yaml:"auto"`
	// tool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/bedrockagent_prompt#tool BedrockagentPrompt#tool}
	Tool interface{} `field:"optional" json:"tool" yaml:"tool"`
}

