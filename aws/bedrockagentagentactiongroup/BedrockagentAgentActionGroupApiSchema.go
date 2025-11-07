// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentagentactiongroup


type BedrockagentAgentActionGroupApiSchema struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/bedrockagent_agent_action_group#payload BedrockagentAgentActionGroup#payload}.
	Payload *string `field:"optional" json:"payload" yaml:"payload"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/bedrockagent_agent_action_group#s3 BedrockagentAgentActionGroup#s3}
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

