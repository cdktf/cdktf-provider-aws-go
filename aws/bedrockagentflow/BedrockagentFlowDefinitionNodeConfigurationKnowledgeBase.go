// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinitionNodeConfigurationKnowledgeBase struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/bedrockagent_flow#knowledge_base_id BedrockagentFlow#knowledge_base_id}.
	KnowledgeBaseId *string `field:"required" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/bedrockagent_flow#model_id BedrockagentFlow#model_id}.
	ModelId *string `field:"required" json:"modelId" yaml:"modelId"`
	// guardrail_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/bedrockagent_flow#guardrail_configuration BedrockagentFlow#guardrail_configuration}
	GuardrailConfiguration interface{} `field:"optional" json:"guardrailConfiguration" yaml:"guardrailConfiguration"`
	// inference_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/bedrockagent_flow#inference_configuration BedrockagentFlow#inference_configuration}
	InferenceConfiguration interface{} `field:"optional" json:"inferenceConfiguration" yaml:"inferenceConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/bedrockagent_flow#number_of_results BedrockagentFlow#number_of_results}.
	NumberOfResults *float64 `field:"optional" json:"numberOfResults" yaml:"numberOfResults"`
}

