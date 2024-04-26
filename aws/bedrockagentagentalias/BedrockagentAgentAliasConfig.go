// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentagentalias

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BedrockagentAgentAliasConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/bedrockagent_agent_alias#agent_alias_name BedrockagentAgentAlias#agent_alias_name}.
	AgentAliasName *string `field:"required" json:"agentAliasName" yaml:"agentAliasName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/bedrockagent_agent_alias#agent_id BedrockagentAgentAlias#agent_id}.
	AgentId *string `field:"required" json:"agentId" yaml:"agentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/bedrockagent_agent_alias#description BedrockagentAgentAlias#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/bedrockagent_agent_alias#routing_configuration BedrockagentAgentAlias#routing_configuration}.
	RoutingConfiguration interface{} `field:"optional" json:"routingConfiguration" yaml:"routingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/bedrockagent_agent_alias#tags BedrockagentAgentAlias#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/bedrockagent_agent_alias#timeouts BedrockagentAgentAlias#timeouts}
	Timeouts *BedrockagentAgentAliasTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

