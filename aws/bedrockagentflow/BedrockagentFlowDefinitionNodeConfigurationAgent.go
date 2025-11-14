// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinitionNodeConfigurationAgent struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/bedrockagent_flow#agent_alias_arn BedrockagentFlow#agent_alias_arn}.
	AgentAliasArn *string `field:"required" json:"agentAliasArn" yaml:"agentAliasArn"`
}

