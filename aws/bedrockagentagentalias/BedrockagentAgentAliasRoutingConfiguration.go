// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentagentalias


type BedrockagentAgentAliasRoutingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/bedrockagent_agent_alias#agent_version BedrockagentAgentAlias#agent_version}.
	AgentVersion *string `field:"optional" json:"agentVersion" yaml:"agentVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/bedrockagent_agent_alias#provisioned_throughput BedrockagentAgentAlias#provisioned_throughput}.
	ProvisionedThroughput *string `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
}

