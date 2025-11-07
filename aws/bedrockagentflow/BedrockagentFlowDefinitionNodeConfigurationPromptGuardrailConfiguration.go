// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinitionNodeConfigurationPromptGuardrailConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/bedrockagent_flow#guardrail_identifier BedrockagentFlow#guardrail_identifier}.
	GuardrailIdentifier *string `field:"required" json:"guardrailIdentifier" yaml:"guardrailIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/bedrockagent_flow#guardrail_version BedrockagentFlow#guardrail_version}.
	GuardrailVersion *string `field:"required" json:"guardrailVersion" yaml:"guardrailVersion"`
}

