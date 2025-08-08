// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockguardrail


type BedrockGuardrailContentPolicyConfig struct {
	// filters_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/bedrock_guardrail#filters_config BedrockGuardrail#filters_config}
	FiltersConfig interface{} `field:"optional" json:"filtersConfig" yaml:"filtersConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/bedrock_guardrail#tier_config BedrockGuardrail#tier_config}.
	TierConfig interface{} `field:"optional" json:"tierConfig" yaml:"tierConfig"`
}

