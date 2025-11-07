// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinitionNodeConfigurationPromptSourceConfigurationInlineTemplateConfigurationText struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/bedrockagent_flow#text BedrockagentFlow#text}.
	Text *string `field:"required" json:"text" yaml:"text"`
	// cache_point block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/bedrockagent_flow#cache_point BedrockagentFlow#cache_point}
	CachePoint interface{} `field:"optional" json:"cachePoint" yaml:"cachePoint"`
	// input_variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/bedrockagent_flow#input_variable BedrockagentFlow#input_variable}
	InputVariable interface{} `field:"optional" json:"inputVariable" yaml:"inputVariable"`
}

