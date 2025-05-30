// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kendradatasource


type KendraDataSourceConfiguration struct {
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/kendra_data_source#s3_configuration KendraDataSource#s3_configuration}
	S3Configuration *KendraDataSourceConfigurationS3Configuration `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
	// template_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/kendra_data_source#template_configuration KendraDataSource#template_configuration}
	TemplateConfiguration *KendraDataSourceConfigurationTemplateConfiguration `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
	// web_crawler_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/kendra_data_source#web_crawler_configuration KendraDataSource#web_crawler_configuration}
	WebCrawlerConfiguration *KendraDataSourceConfigurationWebCrawlerConfiguration `field:"optional" json:"webCrawlerConfiguration" yaml:"webCrawlerConfiguration"`
}

