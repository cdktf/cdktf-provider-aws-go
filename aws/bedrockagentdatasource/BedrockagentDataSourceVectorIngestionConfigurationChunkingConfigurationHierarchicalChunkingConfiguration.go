// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentdatasource


type BedrockagentDataSourceVectorIngestionConfigurationChunkingConfigurationHierarchicalChunkingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/bedrockagent_data_source#overlap_tokens BedrockagentDataSource#overlap_tokens}.
	OverlapTokens *float64 `field:"required" json:"overlapTokens" yaml:"overlapTokens"`
	// level_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/bedrockagent_data_source#level_configuration BedrockagentDataSource#level_configuration}
	LevelConfiguration interface{} `field:"optional" json:"levelConfiguration" yaml:"levelConfiguration"`
}

