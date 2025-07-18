// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerfeaturegroup


type SagemakerFeatureGroupFeatureDefinition struct {
	// collection_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sagemaker_feature_group#collection_config SagemakerFeatureGroup#collection_config}
	CollectionConfig *SagemakerFeatureGroupFeatureDefinitionCollectionConfig `field:"optional" json:"collectionConfig" yaml:"collectionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sagemaker_feature_group#collection_type SagemakerFeatureGroup#collection_type}.
	CollectionType *string `field:"optional" json:"collectionType" yaml:"collectionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sagemaker_feature_group#feature_name SagemakerFeatureGroup#feature_name}.
	FeatureName *string `field:"optional" json:"featureName" yaml:"featureName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sagemaker_feature_group#feature_type SagemakerFeatureGroup#feature_type}.
	FeatureType *string `field:"optional" json:"featureType" yaml:"featureType"`
}

