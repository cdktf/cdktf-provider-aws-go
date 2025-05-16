// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentdatasource


type BedrockagentDataSourceDataSourceConfigurationSalesforceConfigurationCrawlerConfigurationFilterConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/bedrockagent_data_source#type BedrockagentDataSource#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// pattern_object_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/bedrockagent_data_source#pattern_object_filter BedrockagentDataSource#pattern_object_filter}
	PatternObjectFilter interface{} `field:"optional" json:"patternObjectFilter" yaml:"patternObjectFilter"`
}

