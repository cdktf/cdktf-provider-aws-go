// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockguardrail


type BedrockGuardrailTopicPolicyConfig struct {
	// topics_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/bedrock_guardrail#topics_config BedrockGuardrail#topics_config}
	TopicsConfig interface{} `field:"optional" json:"topicsConfig" yaml:"topicsConfig"`
}

