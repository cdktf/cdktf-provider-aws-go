// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinitionNodeConfigurationPromptSourceConfigurationInline struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/bedrockagent_flow#model_id BedrockagentFlow#model_id}.
	ModelId *string `field:"required" json:"modelId" yaml:"modelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/bedrockagent_flow#template_type BedrockagentFlow#template_type}.
	TemplateType *string `field:"required" json:"templateType" yaml:"templateType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/bedrockagent_flow#additional_model_request_fields BedrockagentFlow#additional_model_request_fields}.
	AdditionalModelRequestFields *string `field:"optional" json:"additionalModelRequestFields" yaml:"additionalModelRequestFields"`
	// inference_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/bedrockagent_flow#inference_configuration BedrockagentFlow#inference_configuration}
	InferenceConfiguration interface{} `field:"optional" json:"inferenceConfiguration" yaml:"inferenceConfiguration"`
	// template_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/bedrockagent_flow#template_configuration BedrockagentFlow#template_configuration}
	TemplateConfiguration interface{} `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
}

