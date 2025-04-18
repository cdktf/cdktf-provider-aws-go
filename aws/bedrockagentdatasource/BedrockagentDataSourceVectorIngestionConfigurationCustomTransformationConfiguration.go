// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentdatasource


type BedrockagentDataSourceVectorIngestionConfigurationCustomTransformationConfiguration struct {
	// intermediate_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/bedrockagent_data_source#intermediate_storage BedrockagentDataSource#intermediate_storage}
	IntermediateStorage interface{} `field:"optional" json:"intermediateStorage" yaml:"intermediateStorage"`
	// transformation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/bedrockagent_data_source#transformation BedrockagentDataSource#transformation}
	Transformation interface{} `field:"optional" json:"transformation" yaml:"transformation"`
}

