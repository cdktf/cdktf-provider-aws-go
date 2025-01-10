// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentknowledgebase


type BedrockagentKnowledgeBaseKnowledgeBaseConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/bedrockagent_knowledge_base#type BedrockagentKnowledgeBase#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// vector_knowledge_base_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/bedrockagent_knowledge_base#vector_knowledge_base_configuration BedrockagentKnowledgeBase#vector_knowledge_base_configuration}
	VectorKnowledgeBaseConfiguration interface{} `field:"optional" json:"vectorKnowledgeBaseConfiguration" yaml:"vectorKnowledgeBaseConfiguration"`
}

