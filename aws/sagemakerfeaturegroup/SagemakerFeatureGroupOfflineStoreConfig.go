// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerfeaturegroup


type SagemakerFeatureGroupOfflineStoreConfig struct {
	// s3_storage_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/sagemaker_feature_group#s3_storage_config SagemakerFeatureGroup#s3_storage_config}
	S3StorageConfig *SagemakerFeatureGroupOfflineStoreConfigS3StorageConfig `field:"required" json:"s3StorageConfig" yaml:"s3StorageConfig"`
	// data_catalog_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/sagemaker_feature_group#data_catalog_config SagemakerFeatureGroup#data_catalog_config}
	DataCatalogConfig *SagemakerFeatureGroupOfflineStoreConfigDataCatalogConfig `field:"optional" json:"dataCatalogConfig" yaml:"dataCatalogConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/sagemaker_feature_group#disable_glue_table_creation SagemakerFeatureGroup#disable_glue_table_creation}.
	DisableGlueTableCreation interface{} `field:"optional" json:"disableGlueTableCreation" yaml:"disableGlueTableCreation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/sagemaker_feature_group#table_format SagemakerFeatureGroup#table_format}.
	TableFormat *string `field:"optional" json:"tableFormat" yaml:"tableFormat"`
}

