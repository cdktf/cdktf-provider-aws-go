// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinitionNode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/bedrockagent_flow#name BedrockagentFlow#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/bedrockagent_flow#type BedrockagentFlow#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/bedrockagent_flow#configuration BedrockagentFlow#configuration}
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/bedrockagent_flow#input BedrockagentFlow#input}
	Input interface{} `field:"optional" json:"input" yaml:"input"`
	// output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/bedrockagent_flow#output BedrockagentFlow#output}
	Output interface{} `field:"optional" json:"output" yaml:"output"`
}

