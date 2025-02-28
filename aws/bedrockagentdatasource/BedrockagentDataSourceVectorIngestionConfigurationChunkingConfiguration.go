// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentdatasource


type BedrockagentDataSourceVectorIngestionConfigurationChunkingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/bedrockagent_data_source#chunking_strategy BedrockagentDataSource#chunking_strategy}.
	ChunkingStrategy *string `field:"required" json:"chunkingStrategy" yaml:"chunkingStrategy"`
	// fixed_size_chunking_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/bedrockagent_data_source#fixed_size_chunking_configuration BedrockagentDataSource#fixed_size_chunking_configuration}
	FixedSizeChunkingConfiguration interface{} `field:"optional" json:"fixedSizeChunkingConfiguration" yaml:"fixedSizeChunkingConfiguration"`
	// hierarchical_chunking_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/bedrockagent_data_source#hierarchical_chunking_configuration BedrockagentDataSource#hierarchical_chunking_configuration}
	HierarchicalChunkingConfiguration interface{} `field:"optional" json:"hierarchicalChunkingConfiguration" yaml:"hierarchicalChunkingConfiguration"`
	// semantic_chunking_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/bedrockagent_data_source#semantic_chunking_configuration BedrockagentDataSource#semantic_chunking_configuration}
	SemanticChunkingConfiguration interface{} `field:"optional" json:"semanticChunkingConfiguration" yaml:"semanticChunkingConfiguration"`
}

