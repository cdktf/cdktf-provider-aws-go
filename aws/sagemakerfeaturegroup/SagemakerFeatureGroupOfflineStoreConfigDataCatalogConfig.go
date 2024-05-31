// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerfeaturegroup


type SagemakerFeatureGroupOfflineStoreConfigDataCatalogConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/sagemaker_feature_group#catalog SagemakerFeatureGroup#catalog}.
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/sagemaker_feature_group#database SagemakerFeatureGroup#database}.
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/sagemaker_feature_group#table_name SagemakerFeatureGroup#table_name}.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

