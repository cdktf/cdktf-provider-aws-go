// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentagentactiongroup


type BedrockagentAgentActionGroupFunctionSchemaMemberFunctionsFunctions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/bedrockagent_agent_action_group#name BedrockagentAgentActionGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/bedrockagent_agent_action_group#description BedrockagentAgentActionGroup#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/bedrockagent_agent_action_group#parameters BedrockagentAgentActionGroup#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

