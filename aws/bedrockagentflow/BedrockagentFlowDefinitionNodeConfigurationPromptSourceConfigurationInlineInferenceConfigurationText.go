// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinitionNodeConfigurationPromptSourceConfigurationInlineInferenceConfigurationText struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/bedrockagent_flow#max_tokens BedrockagentFlow#max_tokens}.
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/bedrockagent_flow#stop_sequences BedrockagentFlow#stop_sequences}.
	StopSequences *[]*string `field:"optional" json:"stopSequences" yaml:"stopSequences"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/bedrockagent_flow#temperature BedrockagentFlow#temperature}.
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/bedrockagent_flow#top_p BedrockagentFlow#top_p}.
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}

