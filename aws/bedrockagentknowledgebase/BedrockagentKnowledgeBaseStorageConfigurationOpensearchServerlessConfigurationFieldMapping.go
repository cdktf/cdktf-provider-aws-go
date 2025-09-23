// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentknowledgebase


type BedrockagentKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMapping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/bedrockagent_knowledge_base#metadata_field BedrockagentKnowledgeBase#metadata_field}.
	MetadataField *string `field:"optional" json:"metadataField" yaml:"metadataField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/bedrockagent_knowledge_base#text_field BedrockagentKnowledgeBase#text_field}.
	TextField *string `field:"optional" json:"textField" yaml:"textField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/bedrockagent_knowledge_base#vector_field BedrockagentKnowledgeBase#vector_field}.
	VectorField *string `field:"optional" json:"vectorField" yaml:"vectorField"`
}

