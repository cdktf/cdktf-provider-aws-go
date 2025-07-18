// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinitionNodeConfigurationInlineCode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/bedrockagent_flow#code BedrockagentFlow#code}.
	Code *string `field:"required" json:"code" yaml:"code"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/bedrockagent_flow#language BedrockagentFlow#language}.
	Language *string `field:"required" json:"language" yaml:"language"`
}

