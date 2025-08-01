// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentagentknowledgebaseassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BedrockagentAgentKnowledgeBaseAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/bedrockagent_agent_knowledge_base_association#agent_id BedrockagentAgentKnowledgeBaseAssociation#agent_id}.
	AgentId *string `field:"required" json:"agentId" yaml:"agentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/bedrockagent_agent_knowledge_base_association#description BedrockagentAgentKnowledgeBaseAssociation#description}.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/bedrockagent_agent_knowledge_base_association#knowledge_base_id BedrockagentAgentKnowledgeBaseAssociation#knowledge_base_id}.
	KnowledgeBaseId *string `field:"required" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/bedrockagent_agent_knowledge_base_association#knowledge_base_state BedrockagentAgentKnowledgeBaseAssociation#knowledge_base_state}.
	KnowledgeBaseState *string `field:"required" json:"knowledgeBaseState" yaml:"knowledgeBaseState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/bedrockagent_agent_knowledge_base_association#agent_version BedrockagentAgentKnowledgeBaseAssociation#agent_version}.
	AgentVersion *string `field:"optional" json:"agentVersion" yaml:"agentVersion"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/bedrockagent_agent_knowledge_base_association#region BedrockagentAgentKnowledgeBaseAssociation#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/bedrockagent_agent_knowledge_base_association#timeouts BedrockagentAgentKnowledgeBaseAssociation#timeouts}
	Timeouts *BedrockagentAgentKnowledgeBaseAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

