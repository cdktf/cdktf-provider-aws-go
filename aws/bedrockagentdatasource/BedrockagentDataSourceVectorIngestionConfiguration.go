// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentdatasource


type BedrockagentDataSourceVectorIngestionConfiguration struct {
	// chunking_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/bedrockagent_data_source#chunking_configuration BedrockagentDataSource#chunking_configuration}
	ChunkingConfiguration interface{} `field:"optional" json:"chunkingConfiguration" yaml:"chunkingConfiguration"`
	// custom_transformation_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/bedrockagent_data_source#custom_transformation_configuration BedrockagentDataSource#custom_transformation_configuration}
	CustomTransformationConfiguration interface{} `field:"optional" json:"customTransformationConfiguration" yaml:"customTransformationConfiguration"`
	// parsing_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/bedrockagent_data_source#parsing_configuration BedrockagentDataSource#parsing_configuration}
	ParsingConfiguration interface{} `field:"optional" json:"parsingConfiguration" yaml:"parsingConfiguration"`
}

