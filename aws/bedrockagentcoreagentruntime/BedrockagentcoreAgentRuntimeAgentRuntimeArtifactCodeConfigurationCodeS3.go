// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreagentruntime


type BedrockagentcoreAgentRuntimeAgentRuntimeArtifactCodeConfigurationCodeS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/bedrockagentcore_agent_runtime#bucket BedrockagentcoreAgentRuntime#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/bedrockagentcore_agent_runtime#prefix BedrockagentcoreAgentRuntime#prefix}.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/bedrockagentcore_agent_runtime#version_id BedrockagentcoreAgentRuntime#version_id}.
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

