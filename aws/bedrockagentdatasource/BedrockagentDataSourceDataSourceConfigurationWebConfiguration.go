// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentdatasource


type BedrockagentDataSourceDataSourceConfigurationWebConfiguration struct {
	// crawler_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/bedrockagent_data_source#crawler_configuration BedrockagentDataSource#crawler_configuration}
	CrawlerConfiguration interface{} `field:"optional" json:"crawlerConfiguration" yaml:"crawlerConfiguration"`
	// source_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/bedrockagent_data_source#source_configuration BedrockagentDataSource#source_configuration}
	SourceConfiguration interface{} `field:"optional" json:"sourceConfiguration" yaml:"sourceConfiguration"`
}

