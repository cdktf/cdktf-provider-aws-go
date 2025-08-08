// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentdatasource


type BedrockagentDataSourceDataSourceConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/bedrockagent_data_source#type BedrockagentDataSource#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// confluence_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/bedrockagent_data_source#confluence_configuration BedrockagentDataSource#confluence_configuration}
	ConfluenceConfiguration interface{} `field:"optional" json:"confluenceConfiguration" yaml:"confluenceConfiguration"`
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/bedrockagent_data_source#s3_configuration BedrockagentDataSource#s3_configuration}
	S3Configuration interface{} `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
	// salesforce_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/bedrockagent_data_source#salesforce_configuration BedrockagentDataSource#salesforce_configuration}
	SalesforceConfiguration interface{} `field:"optional" json:"salesforceConfiguration" yaml:"salesforceConfiguration"`
	// share_point_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/bedrockagent_data_source#share_point_configuration BedrockagentDataSource#share_point_configuration}
	SharePointConfiguration interface{} `field:"optional" json:"sharePointConfiguration" yaml:"sharePointConfiguration"`
	// web_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/bedrockagent_data_source#web_configuration BedrockagentDataSource#web_configuration}
	WebConfiguration interface{} `field:"optional" json:"webConfiguration" yaml:"webConfiguration"`
}

