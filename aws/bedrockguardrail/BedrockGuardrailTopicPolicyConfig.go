// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockguardrail


type BedrockGuardrailTopicPolicyConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/bedrock_guardrail#tier_config BedrockGuardrail#tier_config}.
	TierConfig interface{} `field:"optional" json:"tierConfig" yaml:"tierConfig"`
	// topics_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/bedrock_guardrail#topics_config BedrockGuardrail#topics_config}
	TopicsConfig interface{} `field:"optional" json:"topicsConfig" yaml:"topicsConfig"`
}

