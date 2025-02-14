// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentagentactiongroup


type BedrockagentAgentActionGroupActionGroupExecutor struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/bedrockagent_agent_action_group#custom_control BedrockagentAgentActionGroup#custom_control}.
	CustomControl *string `field:"optional" json:"customControl" yaml:"customControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/bedrockagent_agent_action_group#lambda BedrockagentAgentActionGroup#lambda}.
	Lambda *string `field:"optional" json:"lambda" yaml:"lambda"`
}

