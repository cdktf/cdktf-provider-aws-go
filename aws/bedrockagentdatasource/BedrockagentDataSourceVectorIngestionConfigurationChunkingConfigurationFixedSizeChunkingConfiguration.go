// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentdatasource


type BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationFixedSizeChunkingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/bedrockagent_data_source#max_tokens BedrockagentDataSource#max_tokens}.
	MaxTokens *float64 `field:"required" json:"maxTokens" yaml:"maxTokens"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/bedrockagent_data_source#overlap_percentage BedrockagentDataSource#overlap_percentage}.
	OverlapPercentage *float64 `field:"required" json:"overlapPercentage" yaml:"overlapPercentage"`
}

