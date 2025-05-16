// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentagent


type BedrockagentAgentMemoryConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/bedrockagent_agent#enabled_memory_types BedrockagentAgent#enabled_memory_types}.
	EnabledMemoryTypes *[]*string `field:"optional" json:"enabledMemoryTypes" yaml:"enabledMemoryTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/bedrockagent_agent#storage_days BedrockagentAgent#storage_days}.
	StorageDays *float64 `field:"optional" json:"storageDays" yaml:"storageDays"`
}

