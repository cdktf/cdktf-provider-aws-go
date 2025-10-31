// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockguardrail


type BedrockGuardrailCrossRegionConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/bedrock_guardrail#guardrail_profile_identifier BedrockGuardrail#guardrail_profile_identifier}.
	GuardrailProfileIdentifier *string `field:"required" json:"guardrailProfileIdentifier" yaml:"guardrailProfileIdentifier"`
}

