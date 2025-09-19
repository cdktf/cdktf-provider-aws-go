// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinitionConnectionConfigurationData struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/bedrockagent_flow#source_output BedrockagentFlow#source_output}.
	SourceOutput *string `field:"required" json:"sourceOutput" yaml:"sourceOutput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/bedrockagent_flow#target_input BedrockagentFlow#target_input}.
	TargetInput *string `field:"required" json:"targetInput" yaml:"targetInput"`
}

