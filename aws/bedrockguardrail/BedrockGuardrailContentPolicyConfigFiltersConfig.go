// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockguardrail


type BedrockGuardrailContentPolicyConfigFiltersConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/bedrock_guardrail#input_strength BedrockGuardrail#input_strength}.
	InputStrength *string `field:"required" json:"inputStrength" yaml:"inputStrength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/bedrock_guardrail#output_strength BedrockGuardrail#output_strength}.
	OutputStrength *string `field:"required" json:"outputStrength" yaml:"outputStrength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/bedrock_guardrail#type BedrockGuardrail#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

