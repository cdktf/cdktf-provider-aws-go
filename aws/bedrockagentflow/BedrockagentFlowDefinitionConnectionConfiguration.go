// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinitionConnectionConfiguration struct {
	// conditional block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/bedrockagent_flow#conditional BedrockagentFlow#conditional}
	Conditional interface{} `field:"optional" json:"conditional" yaml:"conditional"`
	// data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/bedrockagent_flow#data BedrockagentFlow#data}
	Data interface{} `field:"optional" json:"data" yaml:"data"`
}

