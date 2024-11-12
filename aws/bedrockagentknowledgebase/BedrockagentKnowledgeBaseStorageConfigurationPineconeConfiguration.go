// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentknowledgebase


type BedrockagentKnowledgeBaseStorageConfigurationPineconeConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/bedrockagent_knowledge_base#connection_string BedrockagentKnowledgeBase#connection_string}.
	ConnectionString *string `field:"required" json:"connectionString" yaml:"connectionString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/bedrockagent_knowledge_base#credentials_secret_arn BedrockagentKnowledgeBase#credentials_secret_arn}.
	CredentialsSecretArn *string `field:"required" json:"credentialsSecretArn" yaml:"credentialsSecretArn"`
	// field_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/bedrockagent_knowledge_base#field_mapping BedrockagentKnowledgeBase#field_mapping}
	FieldMapping interface{} `field:"optional" json:"fieldMapping" yaml:"fieldMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/bedrockagent_knowledge_base#namespace BedrockagentKnowledgeBase#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

