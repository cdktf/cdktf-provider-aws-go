// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockguardrail


type BedrockGuardrailSensitiveInformationPolicyConfigPiiEntitiesConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/bedrock_guardrail#action BedrockGuardrail#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.93.0/docs/resources/bedrock_guardrail#type BedrockGuardrail#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

