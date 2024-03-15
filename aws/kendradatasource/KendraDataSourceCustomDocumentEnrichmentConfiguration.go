// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kendradatasource


type KendraDataSourceCustomDocumentEnrichmentConfiguration struct {
	// inline_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/kendra_data_source#inline_configurations KendraDataSource#inline_configurations}
	InlineConfigurations interface{} `field:"optional" json:"inlineConfigurations" yaml:"inlineConfigurations"`
	// post_extraction_hook_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/kendra_data_source#post_extraction_hook_configuration KendraDataSource#post_extraction_hook_configuration}
	PostExtractionHookConfiguration *KendraDataSourceCustomDocumentEnrichmentConfigurationPostExtractionHookConfiguration `field:"optional" json:"postExtractionHookConfiguration" yaml:"postExtractionHookConfiguration"`
	// pre_extraction_hook_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/kendra_data_source#pre_extraction_hook_configuration KendraDataSource#pre_extraction_hook_configuration}
	PreExtractionHookConfiguration *KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfiguration `field:"optional" json:"preExtractionHookConfiguration" yaml:"preExtractionHookConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/kendra_data_source#role_arn KendraDataSource#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

