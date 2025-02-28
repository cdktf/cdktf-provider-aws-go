// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentdatasource


type BedrockagentDataSourceVectorIngestionConfigurationParsingConfigurationBedrockFoundationModelConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/bedrockagent_data_source#model_arn BedrockagentDataSource#model_arn}.
	ModelArn *string `field:"required" json:"modelArn" yaml:"modelArn"`
	// parsing_prompt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/bedrockagent_data_source#parsing_prompt BedrockagentDataSource#parsing_prompt}
	ParsingPrompt interface{} `field:"optional" json:"parsingPrompt" yaml:"parsingPrompt"`
}

