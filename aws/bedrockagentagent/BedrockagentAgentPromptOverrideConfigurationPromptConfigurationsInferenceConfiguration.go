// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentagent


type BedrockagentAgentPromptOverrideConfigurationPromptConfigurationsInferenceConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/bedrockagent_agent#max_length BedrockagentAgent#max_length}.
	MaxLength *float64 `field:"optional" json:"maxLength" yaml:"maxLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/bedrockagent_agent#stop_sequences BedrockagentAgent#stop_sequences}.
	StopSequences *[]*string `field:"optional" json:"stopSequences" yaml:"stopSequences"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/bedrockagent_agent#temperature BedrockagentAgent#temperature}.
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/bedrockagent_agent#top_k BedrockagentAgent#top_k}.
	TopK *float64 `field:"optional" json:"topK" yaml:"topK"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/bedrockagent_agent#top_p BedrockagentAgent#top_p}.
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}

