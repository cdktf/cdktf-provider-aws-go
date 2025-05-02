// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockguardrail


type BedrockGuardrailWordPolicyConfig struct {
	// managed_word_lists_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/bedrock_guardrail#managed_word_lists_config BedrockGuardrail#managed_word_lists_config}
	ManagedWordListsConfig interface{} `field:"optional" json:"managedWordListsConfig" yaml:"managedWordListsConfig"`
	// words_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.97.0/docs/resources/bedrock_guardrail#words_config BedrockGuardrail#words_config}
	WordsConfig interface{} `field:"optional" json:"wordsConfig" yaml:"wordsConfig"`
}

