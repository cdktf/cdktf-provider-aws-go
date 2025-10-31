// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow


type BedrockagentFlowDefinitionNodeConfigurationRetrieval struct {
	// service_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/bedrockagent_flow#service_configuration BedrockagentFlow#service_configuration}
	ServiceConfiguration interface{} `field:"optional" json:"serviceConfiguration" yaml:"serviceConfiguration"`
}

