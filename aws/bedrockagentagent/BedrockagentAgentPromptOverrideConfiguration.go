// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentagent


type BedrockagentAgentPromptOverrideConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/bedrockagent_agent#override_lambda BedrockagentAgent#override_lambda}.
	OverrideLambda *string `field:"optional" json:"overrideLambda" yaml:"overrideLambda"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/bedrockagent_agent#prompt_configurations BedrockagentAgent#prompt_configurations}.
	PromptConfigurations interface{} `field:"optional" json:"promptConfigurations" yaml:"promptConfigurations"`
}

