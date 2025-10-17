// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerfeaturegroup


type SagemakerFeatureGroupFeatureDefinitionCollectionConfig struct {
	// vector_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/sagemaker_feature_group#vector_config SagemakerFeatureGroup#vector_config}
	VectorConfig *SagemakerFeatureGroupFeatureDefinitionCollectionConfigVectorConfig `field:"optional" json:"vectorConfig" yaml:"vectorConfig"`
}

