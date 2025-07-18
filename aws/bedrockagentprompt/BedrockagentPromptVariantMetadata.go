// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentprompt


type BedrockagentPromptVariantMetadata struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/bedrockagent_prompt#key BedrockagentPrompt#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/bedrockagent_prompt#value BedrockagentPrompt#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

