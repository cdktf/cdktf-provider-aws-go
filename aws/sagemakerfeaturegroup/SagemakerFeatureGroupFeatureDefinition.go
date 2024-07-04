// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerfeaturegroup


type SagemakerFeatureGroupFeatureDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/sagemaker_feature_group#feature_name SagemakerFeatureGroup#feature_name}.
	FeatureName *string `field:"optional" json:"featureName" yaml:"featureName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/sagemaker_feature_group#feature_type SagemakerFeatureGroup#feature_type}.
	FeatureType *string `field:"optional" json:"featureType" yaml:"featureType"`
}

