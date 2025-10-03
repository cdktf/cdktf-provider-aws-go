// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinition struct {
	// connection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/bedrockagent_flow#connection BedrockagentFlow#connection}
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/bedrockagent_flow#node BedrockagentFlow#node}
	NodeAttribute interface{} `field:"optional" json:"nodeAttribute" yaml:"nodeAttribute"`
}

