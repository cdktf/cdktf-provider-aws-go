// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kendradatasource


type KendraDataSourceConfigurationS3Configuration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/kendra_data_source#bucket_name KendraDataSource#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// access_control_list_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/kendra_data_source#access_control_list_configuration KendraDataSource#access_control_list_configuration}
	AccessControlListConfiguration *KendraDataSourceConfigurationS3ConfigurationAccessControlListConfiguration `field:"optional" json:"accessControlListConfiguration" yaml:"accessControlListConfiguration"`
	// documents_metadata_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/kendra_data_source#documents_metadata_configuration KendraDataSource#documents_metadata_configuration}
	DocumentsMetadataConfiguration *KendraDataSourceConfigurationS3ConfigurationDocumentsMetadataConfiguration `field:"optional" json:"documentsMetadataConfiguration" yaml:"documentsMetadataConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/kendra_data_source#exclusion_patterns KendraDataSource#exclusion_patterns}.
	ExclusionPatterns *[]*string `field:"optional" json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/kendra_data_source#inclusion_patterns KendraDataSource#inclusion_patterns}.
	InclusionPatterns *[]*string `field:"optional" json:"inclusionPatterns" yaml:"inclusionPatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/kendra_data_source#inclusion_prefixes KendraDataSource#inclusion_prefixes}.
	InclusionPrefixes *[]*string `field:"optional" json:"inclusionPrefixes" yaml:"inclusionPrefixes"`
}

