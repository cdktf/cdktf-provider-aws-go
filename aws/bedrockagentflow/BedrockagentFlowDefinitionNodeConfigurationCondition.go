// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinitionNodeConfigurationCondition struct {
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/bedrockagent_flow#condition BedrockagentFlow#condition}
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
}

