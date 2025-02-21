// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentagent


type BedrockagentAgentPromptOverrideConfigurationPromptConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/bedrockagent_agent#base_prompt_template BedrockagentAgent#base_prompt_template}.
	BasePromptTemplate *string `field:"optional" json:"basePromptTemplate" yaml:"basePromptTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/bedrockagent_agent#inference_configuration BedrockagentAgent#inference_configuration}.
	InferenceConfiguration interface{} `field:"optional" json:"inferenceConfiguration" yaml:"inferenceConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/bedrockagent_agent#parser_mode BedrockagentAgent#parser_mode}.
	ParserMode *string `field:"optional" json:"parserMode" yaml:"parserMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/bedrockagent_agent#prompt_creation_mode BedrockagentAgent#prompt_creation_mode}.
	PromptCreationMode *string `field:"optional" json:"promptCreationMode" yaml:"promptCreationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/bedrockagent_agent#prompt_state BedrockagentAgent#prompt_state}.
	PromptState *string `field:"optional" json:"promptState" yaml:"promptState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/bedrockagent_agent#prompt_type BedrockagentAgent#prompt_type}.
	PromptType *string `field:"optional" json:"promptType" yaml:"promptType"`
}

