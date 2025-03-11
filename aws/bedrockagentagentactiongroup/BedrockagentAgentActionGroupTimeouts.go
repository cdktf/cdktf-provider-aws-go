// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentagentactiongroup


type BedrockagentAgentActionGroupTimeouts struct {
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/bedrockagent_agent_action_group#create BedrockagentAgentActionGroup#create}
	Create *string `field:"optional" json:"create" yaml:"create"`
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/bedrockagent_agent_action_group#update BedrockagentAgentActionGroup#update}
	Update *string `field:"optional" json:"update" yaml:"update"`
}

